package swap

import (
	"encoding/json"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/anoideaopen/foundation/core/types"
	pbfound "github.com/anoideaopen/foundation/proto"
	"github.com/anoideaopen/foundation/test/integration/cmn"
	"github.com/anoideaopen/foundation/test/integration/cmn/client"
	"github.com/anoideaopen/foundation/test/integration/cmn/fabricnetwork"
	"github.com/anoideaopen/foundation/test/integration/cmn/runner"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/integration/nwo"
	"github.com/hyperledger/fabric/integration/nwo/fabricconfig"
	runnerFbk "github.com/hyperledger/fabric/integration/nwo/runner"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/tedsuo/ifrit"
	ginkgomon "github.com/tedsuo/ifrit/ginkgomon_v2"
)

var _ = Describe("Swap Tests", func() {
	var (
		testDir          string
		cli              *docker.Client
		network          *nwo.Network
		networkProcess   ifrit.Process
		ordererProcesses []ifrit.Process
		peerProcesses    ifrit.Process
		channels         = []string{cmn.ChannelAcl, cmn.ChannelCC, cmn.ChannelFiat, cmn.ChannelIndustrial}
		ordererRunners   []*ginkgomon.Runner
		redisProcess     ifrit.Process
		redisDB          *runner.RedisDB
		networkFound     *cmn.NetworkFoundation
		robotProc        ifrit.Process
		skiBackend       string
		skiRobot         string
		peer             *nwo.Peer
		admin            *client.UserFoundation
		feeSetter        *client.UserFoundation
		feeAddressSetter *client.UserFoundation
	)

	BeforeEach(func() {
		networkProcess = nil
		ordererProcesses = nil
		peerProcesses = nil
		var err error
		testDir, err = os.MkdirTemp("", "foundation")
		Expect(err).NotTo(HaveOccurred())

		cli, err = docker.NewClientFromEnv()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if networkProcess != nil {
			networkProcess.Signal(syscall.SIGTERM)
			Eventually(networkProcess.Wait(), network.EventuallyTimeout).Should(Receive())
		}
		if peerProcesses != nil {
			peerProcesses.Signal(syscall.SIGTERM)
			Eventually(peerProcesses.Wait(), network.EventuallyTimeout).Should(Receive())
		}
		if network != nil {
			network.Cleanup()
		}
		for _, ordererInstance := range ordererProcesses {
			ordererInstance.Signal(syscall.SIGTERM)
			Eventually(ordererInstance.Wait(), network.EventuallyTimeout).Should(Receive())
		}
		err := os.RemoveAll(testDir)
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {
		By("start redis")
		redisDB = &runner.RedisDB{}
		redisProcess = ifrit.Invoke(redisDB)
		Eventually(redisProcess.Ready(), runnerFbk.DefaultStartTimeout).Should(BeClosed())
		Consistently(redisProcess.Wait()).ShouldNot(Receive())
	})

	BeforeEach(func() {
		networkConfig := nwo.MultiNodeSmartBFT()
		networkConfig.Channels = nil

		pchs := make([]*nwo.PeerChannel, 0, cap(channels))
		for _, ch := range channels {
			pchs = append(pchs, &nwo.PeerChannel{
				Name:   ch,
				Anchor: true,
			})
		}
		for _, peer := range networkConfig.Peers {
			peer.Channels = pchs
		}

		network = nwo.New(networkConfig, testDir, cli, StartPort(), components)
		cwd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		network.ExternalBuilders = append(network.ExternalBuilders,
			fabricconfig.ExternalBuilder{
				Path:                 filepath.Join(cwd, ".", "externalbuilders", "binary"),
				Name:                 "binary",
				PropagateEnvironment: []string{"GOPROXY"},
			},
		)

		networkFound = cmn.New(network, channels)
		networkFound.Robot.RedisAddresses = []string{redisDB.Address()}

		networkFound.GenerateConfigTree()
		networkFound.Bootstrap()

		for _, orderer := range network.Orderers {
			runner := network.OrdererRunner(orderer)
			runner.Command.Env = append(runner.Command.Env, "FABRIC_LOGGING_SPEC=orderer.consensus.smartbft=debug:grpc=debug")
			ordererRunners = append(ordererRunners, runner)
			proc := ifrit.Invoke(runner)
			ordererProcesses = append(ordererProcesses, proc)
			Eventually(proc.Ready(), network.EventuallyTimeout).Should(BeClosed())
		}

		peerGroupRunner, _ := fabricnetwork.PeerGroupRunners(network)
		peerProcesses = ifrit.Invoke(peerGroupRunner)
		Eventually(peerProcesses.Ready(), network.EventuallyTimeout).Should(BeClosed())

		By("Joining orderers to channels")
		for _, channel := range channels {
			fabricnetwork.JoinChannel(network, channel)
		}

		By("Waiting for followers to see the leader")
		Eventually(ordererRunners[1].Err(), network.EventuallyTimeout, time.Second).Should(gbytes.Say("Message from 1"))
		Eventually(ordererRunners[2].Err(), network.EventuallyTimeout, time.Second).Should(gbytes.Say("Message from 1"))
		Eventually(ordererRunners[3].Err(), network.EventuallyTimeout, time.Second).Should(gbytes.Say("Message from 1"))

		By("Joining peers to channels")
		for _, channel := range channels {
			network.JoinChannel(channel, network.Orderers[0], network.PeersWithChannel(channel)...)
		}

		peer = network.Peer("Org1", "peer0")

		pathToPrivateKeyBackend := network.PeerUserKey(peer, "User1")
		skiBackend, err = cmn.ReadSKI(pathToPrivateKeyBackend)
		Expect(err).NotTo(HaveOccurred())

		pathToPrivateKeyRobot := network.PeerUserKey(peer, "User2")
		skiRobot, err = cmn.ReadSKI(pathToPrivateKeyRobot)
		Expect(err).NotTo(HaveOccurred())

		admin, err = client.NewUserFoundation(pbfound.KeyType_ed25519)
		Expect(err).NotTo(HaveOccurred())
		Expect(admin.PrivateKeyBytes).NotTo(Equal(nil))

		feeSetter, err = client.NewUserFoundation(pbfound.KeyType_ed25519)
		Expect(err).NotTo(HaveOccurred())
		Expect(feeSetter.PrivateKeyBytes).NotTo(Equal(nil))

		feeAddressSetter, err = client.NewUserFoundation(pbfound.KeyType_ed25519)
		Expect(err).NotTo(HaveOccurred())
		Expect(feeAddressSetter.PrivateKeyBytes).NotTo(Equal(nil))

		cmn.DeployACL(network, components, peer, testDir, skiBackend, admin.PublicKeyBase58, admin.KeyType)
		cmn.DeployCC(network, components, peer, testDir, skiRobot, admin.AddressBase58Check)
		cmn.DeployFiat(network, components, peer, testDir, skiRobot,
			admin.AddressBase58Check, feeSetter.AddressBase58Check, feeAddressSetter.AddressBase58Check)
		cmn.DeployIndustrial(network, components, peer, testDir, skiRobot,
			admin.AddressBase58Check, feeSetter.AddressBase58Check, feeAddressSetter.AddressBase58Check)
	})

	BeforeEach(func() {
		By("start robot")
		robotRunner := networkFound.RobotRunner()
		robotProc = ifrit.Invoke(robotRunner)
		Eventually(robotProc.Ready(), network.EventuallyTimeout).Should(BeClosed())
	})

	AfterEach(func() {
		By("stop robot")
		if robotProc != nil {
			robotProc.Signal(syscall.SIGTERM)
			Eventually(robotProc.Wait(), network.EventuallyTimeout).Should(Receive())
		}
	})

	AfterEach(func() {
		By("stop redis " + redisDB.Address())
		if redisProcess != nil {
			redisProcess.Signal(syscall.SIGTERM)
			Eventually(redisProcess.Wait(), time.Minute).Should(Receive())
		}
	})

	Describe("swap tests", func() {
		var (
			user1           *client.UserFoundation
			swapAmount      = "1"
			zeroAmount      = "0"
			defaultSwapHash = "7d4e3eec80026719639ed4dba68916eb94c7a49a053e05c8f9578fe4e5a3d7ea"
			defaultSwapKey  = "12345"
		)

		BeforeEach(func() {
			By("add admin to acl")
			client.AddUser(network, peer, network.Orderers[0], admin)

			By("add user to acl")
			var err error
			user1, err = client.NewUserFoundation(pbfound.KeyType_ed25519)
			Expect(err).NotTo(HaveOccurred())

			client.AddUser(network, peer, network.Orderers[0], user1)

			By("emit tokens 1000")
			client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelFiat, cmn.ChannelFiat, admin,
				"emit", "", client.NewNonceByTime().Get(), nil, user1.AddressBase58Check, swapAmount)

			By("emit check")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"balanceOf", user1.AddressBase58Check)
		})

		It("swap token from fiat to cc then swap from cc to fiat", func() {
			By("swap from fiat to cc")
			By("swap begin")
			swapBeginTxID := client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelFiat, cmn.ChannelFiat, user1,
				"swapBegin", "", client.NewNonceByTime().Get(), nil,
				"FIAT", "CC", swapAmount, defaultSwapHash)
			Expect(swapBeginTxID).ToNot(BeEmpty())

			By("swap get")
			fGet := func(out []byte) string {
				if len(out) == 0 {
					return "out is empty"
				}

				return ""
			}
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(fGet, nil),
				"swapGet", swapBeginTxID)

			By("check balance 1")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 1")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("swap done")
			client.NBTxInvoke(network, peer, network.Orderers[0], nil,
				cmn.ChannelCC, cmn.ChannelCC,
				"swapDone", swapBeginTxID, defaultSwapKey)

			By("check balance 2")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 2")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("swap from cc to fiat")
			By("swap begin")
			swapBeginTxID = client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelCC, cmn.ChannelCC, user1,
				"swapBegin", "", client.NewNonceByTime().Get(), nil,
				"FIAT", "FIAT", swapAmount, defaultSwapHash)
			Expect(swapBeginTxID).ToNot(BeEmpty())

			By("swap get")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(fGet, nil),
				"swapGet", swapBeginTxID)

			By("check balance 1")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 1")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("swap done")
			client.NBTxInvoke(network, peer, network.Orderers[0], nil,
				cmn.ChannelFiat, cmn.ChannelFiat,
				"swapDone", swapBeginTxID, defaultSwapKey)

			By("check balance 2")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 2")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")
		})

		It("swap token from fiat to cc and swap cancel", func() {
			By("swap from fiat to cc")
			By("swap begin")
			swapBeginTxID := client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelFiat, cmn.ChannelFiat, user1,
				"swapBegin", "", client.NewNonceByTime().Get(), nil,
				"FIAT", "CC", swapAmount, defaultSwapHash)
			Expect(swapBeginTxID).ToNot(BeEmpty())

			By("swap get")
			fGet := func(out []byte) string {
				if len(out) == 0 {
					return "out is empty"
				}

				return ""
			}
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(fGet, nil),
				"swapGet", swapBeginTxID)

			By("check balance 1")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 1")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("swap cancel on channel cc")
			client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelCC, cmn.ChannelCC, user1,
				"swapCancel", "", client.NewNonceByTime().Get(), nil,
				swapBeginTxID)

			By("swap cancel on channel fiat")
			client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelFiat, cmn.ChannelFiat, user1,
				"swapCancel", "", client.NewNonceByTime().Get(), nil,
				swapBeginTxID)

			By("check balance 2")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 2")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")
		})
	})

	Describe("multiswap tests", func() {
		var (
			user1           *client.UserFoundation
			swapAmount      = "1"
			zeroAmount      = "0"
			defaultSwapHash = "7d4e3eec80026719639ed4dba68916eb94c7a49a053e05c8f9578fe4e5a3d7ea"
			defaultSwapKey  = "12345"
		)

		BeforeEach(func() {
			By("add admin to acl")
			client.AddUser(network, peer, network.Orderers[0], admin)

			By("add user to acl")
			var err error
			user1, err = client.NewUserFoundation(pbfound.KeyType_ed25519)
			Expect(err).NotTo(HaveOccurred())

			client.AddUser(network, peer, network.Orderers[0], user1)

			By("emit tokens 1000")
			client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelFiat, cmn.ChannelFiat, admin,
				"emit", "", client.NewNonceByTime().Get(), nil, user1.AddressBase58Check, swapAmount)

			By("emit check")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"balanceOf", user1.AddressBase58Check)
		})

		It("multiswap token from fiat to cc then multiswap from cc to fiat", func() {
			By("multiswap from fiat to cc")
			By("multiswap begin")
			assets, err := json.Marshal(types.MultiSwapAssets{
				Assets: []*types.MultiSwapAsset{
					{
						Group:  "FIAT",
						Amount: swapAmount,
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			swapBeginTxID := client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelFiat, cmn.ChannelFiat, user1,
				"multiSwapBegin", "", client.NewNonceByTime().Get(), nil,
				"FIAT", string(assets), "CC", defaultSwapHash)
			Expect(swapBeginTxID).ToNot(BeEmpty())

			By("multiswap get 1")
			fGet := func(out []byte) string {
				if len(out) == 0 {
					return "out is empty"
				}

				return ""
			}
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(fGet, nil),
				"multiSwapGet", swapBeginTxID)

			By("check balance 1")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 1")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("multiswap done")
			client.NBTxInvoke(network, peer, network.Orderers[0], nil,
				cmn.ChannelCC, cmn.ChannelCC,
				"multiSwapDone", swapBeginTxID, defaultSwapKey)

			By("check balance 2")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 2")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("multiswap from cc to fiat")
			By("multiswap begin")
			swapBeginTxID = client.TxInvokeWithSign(network, peer, network.Orderers[0],
				cmn.ChannelCC, cmn.ChannelCC, user1,
				"multiSwapBegin", "", client.NewNonceByTime().Get(), nil,
				"FIAT", string(assets), "FIAT", defaultSwapHash)
			Expect(swapBeginTxID).ToNot(BeEmpty())

			By("multiswap get 2")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(fGet, nil),
				"multiSwapGet", swapBeginTxID)

			By("check balance 3")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 3")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")

			By("multiswap done")
			client.NBTxInvoke(network, peer, network.Orderers[0], nil,
				cmn.ChannelFiat, cmn.ChannelFiat,
				"multiSwapDone", swapBeginTxID, defaultSwapKey)

			By("check balance 4")
			client.Query(network, peer, cmn.ChannelFiat, cmn.ChannelFiat,
				client.CheckResult(client.CheckBalance(swapAmount), nil),
				"balanceOf", user1.AddressBase58Check)

			By("check allowed balance 4")
			client.Query(network, peer, cmn.ChannelCC, cmn.ChannelCC,
				client.CheckResult(client.CheckBalance(zeroAmount), nil),
				"allowedBalanceOf", user1.AddressBase58Check, "FIAT")
		})
	})
})

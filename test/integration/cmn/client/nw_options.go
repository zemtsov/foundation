package client

import (
	"errors"
	"fmt"

	"github.com/anoideaopen/foundation/test/integration/cmn"
	"github.com/hyperledger/fabric/integration"
	"github.com/hyperledger/fabric/integration/nwo"
)

// networkOptions is a struct for network mandatory and parameters that could be specified while network initializes
type networkOptions struct {
	Channels           []*cmn.Channel
	TestPort           integration.TestPortRange
	RobotCfg           *cmn.Robot
	ChannelTransferCfg *cmn.ChannelTransfer
	Templates          *cmn.TemplatesFound
}

// NetworkOption specifies some networkOption parameter
type NetworkOption func(opts *networkOptions) error

// Network mandatory parameters

// WithChannels specifies network channels
func WithChannels(channels []string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.Channels = cmn.ChannelsFromNames(channels)
		return nil
	}
}

// WithTestPort specifies network test port
func WithTestPort(port integration.TestPortRange) NetworkOption {
	return func(opt *networkOptions) error {
		opt.TestPort = port
		return nil
	}
}

// Robot options

// WithRobotPorts specifies robot ports
func WithRobotPorts(ports nwo.Ports) NetworkOption {
	return func(opt *networkOptions) error {
		opt.RobotCfg.Ports = ports
		return nil
	}
}

// WithRobotRedisAddresses specifies robot redis addresses
func WithRobotRedisAddresses(addresses []string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.RobotCfg.RedisAddresses = addresses
		return nil
	}
}

// WithRobotRedisAddress adds specified redis address to robot redis addresses
func WithRobotRedisAddress(address string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.RobotCfg.RedisAddresses = append(opt.RobotCfg.RedisAddresses, address)
		return nil
	}
}

// Channel transfer options

// WithChannelTransferHost specifies channel transfer host address
func WithChannelTransferHost(host string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.ChannelTransferCfg.HostAddress = host
		return nil
	}
}

// WithChannelTransferPorts specifies channel transfer ports
func WithChannelTransferPorts(ports nwo.Ports) NetworkOption {
	return func(opt *networkOptions) error {
		opt.ChannelTransferCfg.Ports = ports
		return nil
	}
}

// WithChannelTransferRedisAddresses specifies channel transfer redis addresses
func WithChannelTransferRedisAddresses(addresses []string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.ChannelTransferCfg.RedisAddresses = addresses
		return nil
	}
}

// WithChannelTransferRedisAddress adds specified redis address to channel transfer redis addresses
func WithChannelTransferRedisAddress(address string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.ChannelTransferCfg.RedisAddresses = append(opt.ChannelTransferCfg.RedisAddresses, address)
		return nil
	}
}

// WithChannelTransferAccessToken specifies channel transfer access token
func WithChannelTransferAccessToken(token string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.ChannelTransferCfg.AccessToken = token
		return nil
	}
}

// WithChannelTransferTTL specifies channel transfer TTL
func WithChannelTransferTTL(ttl string) NetworkOption {
	return func(opt *networkOptions) error {
		opt.ChannelTransferCfg.TTL = ttl
		return nil
	}
}

// TaskExecutor options

func WithTaskExecutorForChannels(host string, ports nwo.Ports, forChannels ...string) NetworkOption {
	return func(opt *networkOptions) error {
		if len(forChannels) == 0 {
			return errors.New("at least one channel is required")
		}

		allChannels := make(map[string]int)
		for i, channel := range opt.Channels {
			allChannels[channel.Name] = i
		}

		for _, forChannel := range forChannels {
			i, channelExists := allChannels[forChannel]
			if !channelExists {
				return fmt.Errorf("channel %s not found", forChannel)
			}
			opt.Channels[i].TaskExecutor = &cmn.TaskExecutor{
				HostAddress: host,
				Ports:       ports,
			}
		}
		return nil
	}
}

// Templates

// WithRobotTemplate specifies robot template
func WithRobotTemplate(robotTemplate string) NetworkOption {
	return func(opt *networkOptions) error {
		if robotTemplate != "" {
			opt.Templates.Robot = robotTemplate
		}
		return nil
	}
}

// WithChannelTransferTemplate specifies channel transfer template
func WithChannelTransferTemplate(ctTemplate string) NetworkOption {
	return func(opt *networkOptions) error {
		if ctTemplate != "" {
			opt.Templates.ChannelTransfer = ctTemplate
		}
		return nil
	}
}

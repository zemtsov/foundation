package client

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

	pbfound "github.com/anoideaopen/foundation/proto"
	"github.com/btcsuite/btcutil/base58"
	eth "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type UserFoundation struct {
	PrivateKeyBytes    []byte
	PublicKeyBytes     []byte
	PublicKeyType      string
	PublicKeyBase58    string
	AddressBase58Check string
	UserID             string
}

func NewUserFoundation(keyType string) *UserFoundation {
	var (
		privateKeyBytes []byte
		publicKeyBytes  []byte
		err             error
	)

	switch keyType {
	case pbfound.KeyType_ed25519.String():
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return &UserFoundation{}
		}

	case pbfound.KeyType_secp256k1.String():
		var privateKey *ecdsa.PrivateKey
		privateKey, err = ecdsa.GenerateKey(eth.S256(), rand.Reader)
		if err != nil {
			return &UserFoundation{}
		}
		privateKeyBytes = privateKey.D.Bytes()
		publicKeyBytes = append(privateKey.X.Bytes(), privateKey.Y.Bytes()...)

	default:
		return &UserFoundation{}
	}

	publicKeyBase58 := base58.Encode(publicKeyBytes)
	hash := sha3.Sum256(publicKeyBytes)
	addressBase58Check := base58.CheckEncode(hash[1:], hash[0])

	return &UserFoundation{
		PrivateKeyBytes:    privateKeyBytes,
		PublicKeyBytes:     publicKeyBytes,
		PublicKeyType:      keyType,
		PublicKeyBase58:    publicKeyBase58,
		AddressBase58Check: addressBase58Check,
		UserID:             "testuser",
	}
}

func (u *UserFoundation) Sign(args ...string) (publicKeyBase58 string, signMsg []byte, err error) {
	publicKeyBase58 = u.PublicKeyBase58
	msg := make([]string, 0, len(args)+1)
	msg = append(msg, args...)
	msg = append(msg, publicKeyBase58)

	bytesToSign := sha3.Sum256([]byte(strings.Join(msg, "")))

	switch u.PublicKeyType {
	case pbfound.KeyType_ed25519.String():
		signMsg = signMessageEd25519(u.PrivateKeyBytes, bytesToSign[:])
		err = verifyEd25519(u.PublicKeyBytes, bytesToSign[:], signMsg)
		if err != nil {
			return "", nil, err
		}

	case pbfound.KeyType_secp256k1.String():
		signMsg = signMessageSecp256k1(u.PrivateKeyBytes, bytesToSign[:])
		err = verifySecp256k1(u.PublicKeyBytes, bytesToSign[:], signMsg)
		if err != nil {
			return "", nil, err
		}

	default:
		return "", nil, errors.New("unknown key type")
	}

	return
}

func (u *UserFoundation) SetUserID(id string) {
	if len(id) != 0 {
		u.UserID = id
	}
}

// signMessageEd25519 - sign arguments with private key in ed25519
func signMessageEd25519(privateKeyBytes []byte, msgToSign []byte) []byte {
	sig := ed25519.Sign(privateKeyBytes, msgToSign)
	return sig
}

// verifyEd25519 - verify publicKey in ed25519 with message and signed message
func verifyEd25519(publicKey []byte, bytesToSign []byte, sMsg []byte) error {
	if !ed25519.Verify(publicKey, bytesToSign, sMsg) {
		err := errors.New("valid signature rejected")
		return err
	}
	return nil
}

// signMessageSecp256k1 - signs a message with private key in secp256k1
func signMessageSecp256k1(privateKeyBytes []byte, msgToSign []byte) []byte {
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = eth.S256()
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, msgToSign)
	if err != nil {
		return nil
	}
	return sig
}

// verifySecp256k1 - verify publicKey in secp256k1 with message and signed message
func verifySecp256k1(publicKeyBytes []byte, message []byte, sig []byte) error {
	const lenSecp256k1Key = 64

	if publicKeyBytes[0] == 0x04 {
		publicKeyBytes = publicKeyBytes[1:]
	}

	if len(publicKeyBytes) != lenSecp256k1Key {
		return errors.New("invalid length of secp256k1 key")
	}

	publicKey := &ecdsa.PublicKey{
		Curve: eth.S256(),
		X:     new(big.Int).SetBytes(publicKeyBytes[:lenSecp256k1Key/2]),
		Y:     new(big.Int).SetBytes(publicKeyBytes[lenSecp256k1Key/2:]),
	}

	if !ecdsa.VerifyASN1(publicKey, message, sig) {
		return errors.New("secp256k1 signature rejected")
	}

	return nil
}

package blsmultisig

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/incognitochain/incognito-chain/common"
)

// KeyGen take an input seed and return BLS Key
func KeyGen(seed []byte) (*big.Int, *bn256.G1) {
	sk := SKGen(seed)
	return sk, PKGen(sk)
}

func SKGen(seed []byte) *big.Int {
	sk := big.NewInt(0)
	sk.SetBytes(common.HashB(seed))
	for {
		if sk.Cmp(bn256.Order) == -1 {
			break
		}
		sk.SetBytes(common.HashB(sk.Bytes()))
	}
	return sk
}

func PKGen(sk *big.Int) *bn256.G1 {
	pk := new(bn256.G1)
	pk = pk.ScalarBaseMult(sk)
	return pk
}

// SKBytes take input secretkey integer and return secretkey bytes
func SKBytes(sk *big.Int) SecretKey {
	return I2Bytes(sk, CSKSz)
}

// PKBytes take input publickey point and return publickey bytes
func PKBytes(pk *bn256.G1) PublicKey {
	return CmprG1(pk)
}

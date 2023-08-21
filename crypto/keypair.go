package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/sharansh123/MyBlockChain/types"
)

type PrivateKey struct{
	Key *ecdsa.PrivateKey
}

type PublicKey struct{
	Key *ecdsa.PublicKey
}


func (k PrivateKey) Sign(data []byte) (*Signature, error){
	r,s, err := ecdsa.Sign(rand.Reader,k.Key, data)
	if err != nil{
		return nil, err
	}

	return &Signature{
		R:r,
		S:s,
	}, nil

}


func GeneratePrivateKey() PrivateKey{
	key, err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	if err != nil{
		panic(err)

	}
	return PrivateKey{
		Key: key,
	}
}

func (k PrivateKey) PublicKey() PublicKey{
	return PublicKey{
		Key: &k.Key.PublicKey,
	}
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.Key, k.Key.X, k.Key.Y)
}

func (k PublicKey) Address() types.Address{
	h := sha256.Sum256(k.ToSlice())
	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct{
	S *big.Int
	R *big.Int
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.Key, data, sig.R, sig.S)
}


package proxy

import (
	"errors"
	"fmt"

	"github.com/tripleblind/random"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	ErrRandom = "failed to retrieve randomness from source: %s"
	ErrUnseal = "broken seal"
)

type NaClProxy struct {
	key    *[32]byte
	source random.Source
}

func NewNaClProxy(source random.Source, key *[32]byte) *NaClProxy {
	return &NaClProxy{
		key:    key,
		source: source,
	}
}

func (n *NaClProxy) Generate(data []byte) ([]byte, error) {

	var nonce [24]byte

	if err := n.source.Fill(nonce[:]); err != nil {
		return nil, fmt.Errorf(ErrRandom, err)
	}

	out := secretbox.Seal(nil, data, &nonce, n.key)

	return append(nonce[:], out[:]...), nil

}

func (n *NaClProxy) Revert(data []byte) ([]byte, error) {

	var nonce [24]byte

	for i, e := range data[:24] {
		nonce[i] = e
	}

	out, ok := secretbox.Open(nil, data[24:], &nonce, n.key)

	if !ok {
		return nil, errors.New(ErrUnseal)
	}

	return out, nil

}

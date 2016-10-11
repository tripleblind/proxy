package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tripleblind/random"
)

func TestNaClProxy(t *testing.T) {

	const reps = 65536
	var key = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	assert := assert.New(t)

	var n Proxy = NewNaClProxy(random.Crypto, &key)

	o1, err := n.Generate([]byte("4711"))
	assert.NotNil(o1)
	assert.NoError(err)

	o2, err := n.Generate([]byte("4711"))
	assert.NotNil(o2)
	assert.NoError(err)

	assert.NotEqual(o1, o2)

	i1, err := n.Revert(o1)
	assert.NotNil(i1)
	assert.NoError(err)

	i2, err := n.Revert(o2)
	assert.NotNil(i2)
	assert.NoError(err)

	assert.Equal(i1, i2)
	assert.Equal([]byte("4711"), i1)

}

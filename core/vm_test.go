package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestVM(t *testing.T){
	data := []byte{0x1, 0xa, 0x2, 0xa, 0xb}
	vm := NewVM(data)
	assert.Nil(t,vm.Run())
	assert.Equal(t, byte(3), vm.stack[2])
}
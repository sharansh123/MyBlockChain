package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestVM(t *testing.T){
	data := []byte{0x2, 0xa, 0x61, 0xc, 0x61, 0xc, 0xd}
	vm := NewVM(data)
	assert.Nil(t,vm.Run())
	//fmt.Println(vm.stack.data)
	result := vm.stack.Pop().([]byte)
	assert.Equal(t, "aa", string(result))
}

func TestStack(t *testing.T){
	st := NewStack(125)
	st.Push(12)
	st.Push(23)
	st.Pop()
	fmt.Printf("%+v", st.data)
	st.Pop()
	fmt.Printf("%+v", st.data)
}
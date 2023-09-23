package core

import "fmt"


type Instruction byte

const (
	InstrPush Instruction = 0x0a
	InstrAdd Instruction = 0x0b
)

type VM struct {
	data []byte
	ip int
	stack []byte
	sp int //stack pointer
}

func NewVM(data []byte) *VM {
	return &VM{
		data: data,
		ip: 0,
		stack: make([]byte, 1024),
		sp: 0,
	}
}

func (vm *VM) Run() error{
	for {
		instr := vm.data[vm.ip]
		fmt.Println(instr)

		if err := vm.Exec(Instruction(instr)); err!=nil{
			fmt.Errorf("Instruction Not Valid")
		}

		vm.ip++
		if vm.ip > len(vm.data)-1{
			break
		}
	}
	return nil
}

func (vm *VM) Exec(instr Instruction) error {
	switch(instr){
	case InstrPush:
		vm.pushStack(vm.data[vm.ip - 1])
	case InstrAdd:
		a := vm.stack[0]
		b := vm.stack[1]
		c := a + b
		vm.pushStack(c)
	}
	return nil
}


func (vm *VM) pushStack(b byte){
	vm.stack[vm.sp] = b
	vm.sp++
}

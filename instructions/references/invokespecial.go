package references

import (
	"github.com/go-toy/jvm/instructions/base"
	"github.com/go-toy/jvm/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}

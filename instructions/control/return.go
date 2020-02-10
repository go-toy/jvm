package control

import (
	"github.com/go-toy/jvm/instructions/base"
	"github.com/go-toy/jvm/rtda"
)

type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	//frame.Thread().PopFrame()
}

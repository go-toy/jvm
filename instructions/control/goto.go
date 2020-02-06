package control

import "github.com/go-toy/jvm/instructions/base"
import "github.com/go-toy/jvm/rtda"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

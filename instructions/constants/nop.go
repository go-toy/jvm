package constants

import "github.com/go-toy/jvm/instructions/base"
import "github.com/go-toy/jvm/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}

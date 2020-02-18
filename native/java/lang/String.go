package lang

import (
	"github.com/go-toy/jvm/native"
	"github.com/go-toy/jvm/rtda"
	"github.com/go-toy/jvm/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned  := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}

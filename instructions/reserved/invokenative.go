package reserved

import (
	"github.com/go-toy/jvm/instructions/base"
	"github.com/go-toy/jvm/native"
	_ "github.com/go-toy/jvm/native/java/lang"
	_ "github.com/go-toy/jvm/native/sun/misc"
	"github.com/go-toy/jvm/rtda"
)

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
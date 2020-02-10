package heap

import "github.com/go-toy/jvm/classfile"

type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	code      []byte
}

func (self *Method) copyAttributes(method *classfile.MemberInfo) {
	if codeAttr := method.CodeAttribute(); codeAttr != nil {
		self.maxLocals = codeAttr.MaxLocals()
		self.maxStack = codeAttr.MaxStack()
		self.code = codeAttr.Code()
	}
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Name() string {
	return self.name
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) Code() []byte {
	return self.code
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

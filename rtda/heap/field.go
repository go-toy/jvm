package heap

import "github.com/go-toy/jvm/classfile"

type Field struct {
	ClassMember
	constantValueIndex uint
	slotId             uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constantValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstantValueIndex() uint {
	return self.constantValueIndex
}

func (self *Field) Class() *Class {
	return self.class
}

func (self *Field) Descriptor() string {
	return self.descriptor
}

func (self *Field) SlotId() uint {
	return self.slotId
}

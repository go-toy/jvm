package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"char":    "C",
	"short":   "S",
	"int":     "I",
	"long":    "L",
	"float":   "F",
	"double":  "D",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}

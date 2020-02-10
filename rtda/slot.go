package rtda

import "github.com/go-toy/jvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}

package Astruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BasicA struct {
	A11 int64
	A2  int32
	A3  string
}

type BasicB struct {
	A1 int64
	A2 int32
	A3 string
}

func TestBasicStruct(t *testing.T) {
	a := &BasicA{
		A11: int64(1),
		A2:  int32(2),
		A3:  "hello",
	}
	b := &BasicB{}
	AssignSameFieldStruct(a, b)
	assert.Equal(t, a, b)
}

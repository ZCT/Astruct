package Astruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BasicA struct {
	A1        int64
	A2        int32
	A3        string
	IntList   []int64
	StringMap map[string]string
}

type BasicB struct {
	A1        int64
	A2        int32
	A3        string
	IntList   []int64
	StringMap map[string]string
}

type AWithNestedBasicA struct {
	A1         string
	B1         *BasicA
	StructList []*BasicA
	StructMap  map[string]*BasicA
}

type BWithNestedBasicA struct {
	A1         int
	B1         *BasicA
	StructList []*BasicA
	StructMap  map[string]*BasicA
}

func initBasicA() *BasicA {
	return &BasicA{
		A1:        int64(1),
		A2:        int32(2),
		A3:        "hello",
		IntList:   []int64{1, 2},
		StringMap: map[string]string{"hello": "world"},
	}
}

func initAWithNestedBasicA() *AWithNestedBasicA {
	a := initBasicA()
	return &AWithNestedBasicA{
		A1:         "hello",
		B1:         a,
		StructList: []*BasicA{initBasicA()},
		StructMap:  map[string]*BasicA{"hello": initBasicA()},
	}
}

func TestBasicStruct(t *testing.T) {
	a := initBasicA()
	b := &BasicB{}
	AssignSameFieldStruct(a, b)
	assert.Equal(t, a.A2, b.A2)
	assert.Equal(t, a.A3, b.A3)
	assert.Equal(t, a.IntList, b.IntList)
	assert.Equal(t, a.StringMap, b.StringMap)
}

func TestNestedStruct(t *testing.T) {
	a := initAWithNestedBasicA()
	b := &BWithNestedBasicA{}
	AssignSameFieldStruct(a, b)
	assert.Equal(t, a.B1, b.B1)
	assert.Equal(t, a.StructList, b.StructList)
	assert.Equal(t, a.StructMap, b.StructMap)
}

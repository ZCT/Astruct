# Astruct

## Description
Auto assign struct field. If two struct have same field name and type, then auto assign the field value

## Examples

```
type BasicA struct {
    DiffA  string
	IntList   []int64
	StringMap map[string]string
}


type BasicB struct {
    DiffB string
	IntList   []int64
	StringMap map[string]string
}


var a = &BasicA {
    DiffA :"hello",
    IntList: []int64{1},
    StringMap: map[string]string{"hello":"World"},
}
var b = &BasicB{}

AssignSameFieldStruct(a, b)

fmt.Printf("%#v", b)

```

 output:
 ```
 {
    DiffB: "",
    IntList: {1}
    StringMap: {"hello":"world"}
 }
 ```



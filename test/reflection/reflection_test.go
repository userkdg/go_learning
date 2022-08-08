package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect1(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

}

func TestReflect1_1(t *testing.T) {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())                                       // v.Uint returns a uint64.

	type MyInt int
	var x2 MyInt = 7
	v2 := reflect.ValueOf(x2)
	fmt.Println("v2 type:", v2.Type(), "value:", v2)
}

func TestReflect3(t *testing.T) {

	type T struct {
		A int
		B string
	}
	t1 := T{23, "skidoo"}
	s := reflect.ValueOf(&t1).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(1)
	s.Field(1).SetString("kdg")
	fmt.Println("CanSet T:", s)
}

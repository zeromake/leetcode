package reflect

import (
	"reflect"
	"testing"
)

func TestReflectValue(t *testing.T)  {
	var num1 int
	var num2 int
	rValue1 := reflect.ValueOf(num1)
	rValue2 := reflect.ValueOf(&num2).Elem()
	t.Logf("%+v:\t%+v\t%v\n", rValue1.Type(), rValue1, rValue1.CanSet())
	t.Logf("%+v:\t%+v\t%v\n", rValue2.Type(), rValue2, rValue2.CanSet())
	rValue2.SetInt(5)
	t.Log(num2)
}


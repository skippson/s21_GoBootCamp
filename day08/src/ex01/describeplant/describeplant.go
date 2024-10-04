package describeplant

import (
	"fmt"
	"reflect"
)

func DescribePlant(obj interface{}) {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fmt.Print(field.Name)
		if len(field.Tag) != 0 {
			fmt.Print("(", field.Tag, ")")
		}
		fmt.Print(":", objVal.FieldByName(field.Name))
		fmt.Println()
	}
}

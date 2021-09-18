package libraries

import (
	"fmt"
	"strconv"

	u "github.com/bcicen/go-units"
)

type FuncMapLibrary struct {
}

func NewFuncMapLibrary() FuncMapLibrary {
	return FuncMapLibrary{}
}

func (lib FuncMapLibrary) Add(value int) string {
	value++
	return strconv.Itoa(value)
}

func (lib FuncMapLibrary) ConvertUnit(value float64, fromConv, toConv string) string {
	var val u.Value
	unitFrom, err := u.Find(fromConv)
	if err != nil {
		fmt.Println(err.Error())
	}
	unitTo, err := u.Find(toConv)
	if err != nil {
		fmt.Println(err.Error())
	}
	val = u.NewValue(value, unitFrom)
	return val.MustConvert(unitTo).String()
}

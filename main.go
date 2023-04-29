package main

import (
	"fmt"
)

type CustomErrorType struct {
	s string
}

func (ce *CustomErrorType) Error() string {
	return ce.s
}

func NewCustomErrorType(s string) error {
	return &CustomErrorType{s: s}
}

func main() {

	someThing := new(int)
	anotherThing := new(int)
	//*anotherThing = 12
	//*someTHing = 12
	//someThing = anotherThing
	//
	fmt.Println(*someThing)
	fmt.Println(*anotherThing)

}

func CHeckStringPointer(s *string) {
	fmt.Println(s)
	*s = "a new one"
}

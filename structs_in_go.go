package main

import "fmt"

type Foo struct {
	A int
	b string
}

func main() {
	teststructs := Foo{}
	fmt.Println(teststructs)

	test2 := Foo{10, "testing"}
	fmt.Println(test2)

	test3 := Foo{b: "testing a"}
	fmt.Println(test3)

}

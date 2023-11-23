package runners

import (
	"fmt"
)

func prettyPrint(m string, v any) {
	fmt.Println()
	fmt.Println("----------")
	fmt.Println(m)
	fmt.Printf("%+v\n", v)
}

func panicErr(m string, err error) {
	fmt.Println()
	fmt.Println("----------")
	fmt.Println(m)
	panic(err)
}

package main

import (
	"fmt"
)

func main() {
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func whatAmI(i interface{})  {
	switch t := i.(type) { case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm a int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}


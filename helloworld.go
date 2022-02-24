package helloworld

import (
	"fmt"
)

func SayHelloTo(who string) {
	fmt.Printf("hello, %v\n", who)
}

func SayHello() {
	SayHelloTo("world")
}

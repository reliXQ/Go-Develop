package gotest //和包名相同

import (
	"fmt"
)

func Hello() {
	fmt.Println("hello from test")
	hi()
}

func hi() {
	fmt.Println("hi from test")
}

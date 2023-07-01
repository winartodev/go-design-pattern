package main

import (
	"fmt"

	firstsingleton "github.com/winartodev/go-design-pattern/creational-pattern/05-singleton/firstSingleton"
	secondsingleton "github.com/winartodev/go-design-pattern/creational-pattern/05-singleton/secondSingleton"
)

func main() {
	firstsingleton.GetInstance()
	firstsingleton.GetInstance().PrintFirstInstanceWithCustomParameter("1")
	firstsingleton.GetInstance().PrintFirstInstanceWithCustomParameter("2")

	fmt.Println()
	secondsingleton.GetInstance()
	secondsingleton.GetInstance()
	secondsingleton.GetInstance()
}

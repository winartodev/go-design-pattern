package secondsingleton

import (
	"fmt"
	"sync"
)

type SecondSingleton struct {
}

var once sync.Once

var instance *SecondSingleton

func GetInstance() *SecondSingleton {

	if instance == nil {
		once.Do(
			func() {
				fmt.Println("create new instance SecondSingleton")
				instance = &SecondSingleton{}
			})
	} else {
		fmt.Println("instance SecondSingleton already exist")

	}

	return instance
}

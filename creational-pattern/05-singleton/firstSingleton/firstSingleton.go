package firstsingleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type FirstSingleton struct {
}

var instance *FirstSingleton

func GetInstance() *FirstSingleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("create new instance first singleton")
				instance = &FirstSingleton{}
			})
	} else {
		fmt.Println("instance first singleton already exist")
	}

	return instance
}

func (f *FirstSingleton) PrintFirstInstanceWithCustomParameter(cutom string) {
	fmt.Printf("Print First Instance %v\n", cutom)
}

package transportfactory

import "fmt"

type Truck struct {
	Transport
}

func TruckFactory() ITransport {
	return &Truck{
		Transport{
			Name: "Truck",
		},
	}
}

func (t *Truck) Deliver() {
	fmt.Println("Deliver by land in a box")
}

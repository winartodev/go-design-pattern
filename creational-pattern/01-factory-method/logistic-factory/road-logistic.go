package logisticfactory

import transportfactory "github.com/winartodev/go-design-pattern/creational-pattern/01-factory-method/transport-factory"

type RoadLogistic struct {
	Logistic
}

func RoadLogisticFactory() ILogistic {
	return &RoadLogistic{
		Logistic{
			Deliver: "Deliver By Road",
		},
	}
}

func (rl *RoadLogistic) CreateTransport() transportfactory.ITransport {
	transport, _ := transportfactory.NewTransportFactory("truck")
	return transport
}

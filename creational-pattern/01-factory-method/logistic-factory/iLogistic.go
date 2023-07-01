package logisticfactory

import transportfactory "github.com/winartodev/go-design-pattern/creational-pattern/01-factory-method/transport-factory"

type ILogistic interface {
	PlanDelivery() string
	CreateTransport() transportfactory.ITransport
}

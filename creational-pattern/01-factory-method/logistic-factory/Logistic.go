package logisticfactory

import transportfactory "github.com/winartodev/go-design-pattern/creational-pattern/01-factory-method/transport-factory"

type Logistic struct {
	Deliver string
}

func (l *Logistic) CreateTransport() transportfactory.ITransport {
	return nil
}

func (l *Logistic) PlanDelivery() string {
	return ""
}

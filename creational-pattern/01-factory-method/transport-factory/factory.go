package transportfactory

import "fmt"

func NewTransportFactory(transportType string) (ITransport, error) {
	if transportType == "truck" {
		return TruckFactory(), nil
	}

	if transportType == "ship" {
		return nil, nil
	}

	return nil, fmt.Errorf("transport type can handle truck and ship")
}

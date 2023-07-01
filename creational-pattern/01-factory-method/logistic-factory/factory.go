package logisticfactory

import "fmt"

func CreateLogistic(logisticType string) (ILogistic, error) {
	if logisticType == "sea" {
		return nil, nil
	}

	if logisticType == "road" {
		return RoadLogisticFactory(), nil
	}

	return nil, fmt.Errorf("logistic can handle road and sea transport")
}

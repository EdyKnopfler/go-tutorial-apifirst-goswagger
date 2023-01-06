package fakedb

import (
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/edyknopfler/apifirst/models"
)

// Database "de brinquedo"
var (
	sensor1 = "sensor1"
	sensor2 = "sensor2"
	sensor3 = "sensor3"
	data = strfmt.DateTime(time.Date(2023, 11, 01, 10, 10, 10, 0, time.UTC))
	medida = "333"
	Items = []*models.Measurement{
		&models.Measurement{&sensor1, &data, &medida},
		&models.Measurement{&sensor2, &data, &medida},
		&models.Measurement{&sensor3, &data, &medida},
	}
)

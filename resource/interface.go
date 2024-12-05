package resource

import "time"

type Interface interface {
	SaveValueReceived(value float32, date time.Time, description string, bank string)
}

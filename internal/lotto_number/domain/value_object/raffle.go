package value_object

import "time"

type Raffle struct {
	numbers []Number
	date    time.Time
}

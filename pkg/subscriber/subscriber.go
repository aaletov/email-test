package subscriber

import (
	"time"
)

type Subscriber struct {
	Email      string
	Name       string
	SecondName string
	Surname    string
	Age        int
	Birthday   time.Time
}

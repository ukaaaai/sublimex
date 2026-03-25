package entity

import (
	"time"
)

type Memory struct {
	Text      string
	Id        string
	Relation  []string
	TimeStamp time.Time
}

package messagebird

import "time"

type HLR struct {
	Id              string
	HRef            string
	MSISDN          int
	Network         int
	Reference       string
	Status          string
	CreatedDatetime *time.Time
	StatusDatetime  *time.Time
	Errors          []Error
}

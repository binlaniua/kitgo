package kitgo

import (
	"time"
)

//-------------------------------------
//
//  Date is ext time object
//
//-------------------------------------
type Date struct {
	t time.Time
}

//
//
// NewDate default time is now
//
//
func NewDate() *Date  {
	d := &Date{
		t: time.Now(),
	}
	return d
}

//
//
// Second return second value
//
//
func (d *Date) Second(second int) int {
	if second != 0 {
		d.t = d.t.Add(time.Duration(second) * time.Second)
	}
	return d.t.Second()
}

//
//
// Hour return hour value
//
//
func (d *Date) Hour(hour int) int  {
	if hour != 0 {
		d.t = d.t.Add(time.Duration(hour) * time.Hour)
	}
	return d.t.Hour()
}

//
//
// Day return day value
//
//
func (d *Date) Minute(m int) int  {
	if m != 0 {
		d.t = d.t.Add(time.Duration(m) * time.Minute)
	}
	return d.t.Minute()
}

//
//
// Format return format date
//
//
func (d *Date) Format(f string) string  {
	return d.t.Format(f)
}
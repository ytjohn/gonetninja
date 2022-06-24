package models

import (
	"encoding/json"
	"github.com/gobuffalo/validate/v3"
	"time"
)

// Quicknet is not tied to any database table
type Quicknet struct {
	Opened          time.Time `json:"opened"`
	Closed          time.Time `json:"closed"`
	NetControl      string    `json:"netcontrol"`
	EarlyCheckins   string    `json:"early_checkins"`
	RegularCheckins string    `json:"regular_checkins"`
}

// String is not required by pop and may be deleted
func (q Quicknet) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// Quicknets is not required by pop and may be deleted
type Quicknets []Quicknet

// String is not required by pop and may be deleted
func (q Quicknets) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

func (q *Quicknet) Validate() (*validate.Errors, error) {

	verrs := validate.NewErrors()
	if q.NetControl == "" {
		verrs.Add("netcontrol", "Netcontrol must not be blank!")
	}
	if q.Opened.IsZero() {
		verrs.Add("opened", "Opened must not be blank")
	}
	// force Closed (if set) to be at least 5 minutes after Open
	if !q.Closed.IsZero() && !q.Opened.IsZero() {
		if q.Closed.Before(q.Opened.Add(time.Minute * 5)) {
			q.Closed = q.Opened.Add(time.Minute * 5)
		}
	}
	//verrs.Add(&validators.StringIsPresent{Field: n.Name, Name: "Name", Message: "Name can not be blank"})
	return verrs, nil
}

//
//// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
//// This method is not required and may be deleted.
//func (q *Quicknet) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
//
//// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
//// This method is not required and may be deleted.
//func (q *Quicknet) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}

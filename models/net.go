package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
)

// Net is used by pop to map your nets database table to your go code.
type Net struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Name         string    `json:"name" db:"name"`
	PlannedStart time.Time `json:"planned_start" db:"planned_start"`
	PlannedEnd   time.Time `json:"planned_end" db:"planned_end"`
}

// String is not required by pop and may be deleted
func (n Net) String() string {
	jn, _ := json.Marshal(n)
	return string(jn)
}

// Nets is not required by pop and may be deleted
type Nets []Net

// String is not required by pop and may be deleted
func (n Nets) String() string {
	jn, _ := json.Marshal(n)
	return string(jn)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (n *Net) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.NewErrors()
	if n.Name == "" {
		verrs.Add("name", "Name must not be blank!")
	}
	//verrs.Add(&validators.StringIsPresent{Field: n.Name, Name: "Name", Message: "Name can not be blank"})
	return verrs, nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (n *Net) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {

	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (n *Net) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
)

// Activity is used by pop to map your activities database table to your go code.
type Activity struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Net         uuid.UUID `json:"net" db:"net"`
	Action      string    `json:"action"  db:"action"`
	Name        string    `json:"name" db:"name"`
	TimeAt      time.Time `json:"time_at" db:"time_at"`
	Description string    `json:"description" db:"description"`
}

// TableName overrides the table name used by Pop.
func (a Activity) TableName() string {
	return "activity"
}

// String is not required by pop and may be deleted
func (a Activity) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Activities is not required by pop and may be deleted
type Activities []Activity

// String is not required by pop and may be deleted
func (a Activities) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Activity) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Activity) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Activity) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

package domain

import (
	strings
	time

	github.com/spy16/droplets/pkg/errors
)

// Meta represents metadata about different entities.
type Meta struct {
	// Name represents a unique name/identifier for the object.
	Name string 

	// Tags can contain additional metadata about the object.
	Tags []string 

	// CreateAt represents the time at which this object was created.
	CreatedAt time.Time 

	// UpdatedAt represents the time at which this object was last
	// modified.
	UpdatedAt time.Time 
}

// SetDefaults sets sensible defaults on meta.
func (meta *Meta) SetDefaults() {
	if meta.CreatedAt.IsZero() {
		meta.CreatedAt = time.Now()
		meta.UpdatedAt = time.Now()
	}
}

// Validate performs basic validation of the metadata.
func (meta Meta) Validate() error {
	switch {
	case empty(meta.Name):
		return errors.MissingField(Name)
	}
	return nil
}

func empty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
welcome }

welcome murali 999

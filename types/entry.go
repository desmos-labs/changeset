package types

import (
	"fmt"
	"time"
)

const (
	CategoryChange = "change"
	CategoryFix    = "fix"
)

// Category represents a category of edit
type Category string

// Validate returns an error if the config contains some invalid values
func (c Category) Validate() error {
	if c != CategoryChange && c != CategoryFix {
		return fmt.Errorf("invalid category: %s", c)
	}
	return nil
}

type TypeCode string

func (code TypeCode) String() string {
	return string(code)
}

// Type defines a possible type for a changeset entry
type Type struct {
	Code        TypeCode `yaml:"code"`
	Description string   `yaml:"description"`
	Category    Category `yaml:"category"`
}

// NewType returns a new Type instance
func NewType(category Category, code TypeCode, description string) *Type {
	return &Type{
		Category:    category,
		Code:        code,
		Description: description,
	}
}

// Validate returns an error if the config contains some invalid values
func (t *Type) Validate() error {
	return t.Category.Validate()
}

// --------------------------------------------------------------------------------------------------------------------

type ModuleID string

// Module contains the details of a single module
type Module struct {
	ID          ModuleID `yaml:"id"`
	Description string   `yaml:"description"`
}

// NewModule allows to build a new Module instance
func NewModule(id ModuleID, description string) *Module {
	return &Module{
		ID:          id,
		Description: description,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// Entry represents a single changeset entry
type Entry struct {
	Type                 TypeCode  `yaml:"type"`
	Module               ModuleID  `yaml:"module"`
	PullRequestID        int       `yaml:"pull_request"`
	Description          string    `yaml:"description"`
	IsBackwardCompatible bool      `yaml:"backward_compatible"`
	Time                 time.Time `yaml:"date"`
}

// NewEntry allows to build a new Entry instance
func NewEntry(
	typeCode TypeCode, moduleID ModuleID, pr int, description string, backwardCompatible bool, time time.Time,
) *Entry {
	return &Entry{
		Type:                 typeCode,
		Module:               moduleID,
		PullRequestID:        pr,
		Description:          description,
		IsBackwardCompatible: backwardCompatible,
		Time:                 time,
	}
}

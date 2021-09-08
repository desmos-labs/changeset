package types

import (
	"time"
)

// TypeCode represents the code of the change type
type TypeCode string

// String implements fmt.Stringer
func (code TypeCode) String() string {
	return string(code)
}

// --------------------------------------------------------------------------------------------------------------------

// Type defines a possible type for a changeset entry
type Type struct {
	Code        TypeCode `yaml:"code"`        // Code of the type
	Title       string   `yaml:"title"`       // Title of the type to be used inside the changelog
	Description string   `yaml:"description"` // Description of the type change
	Hide        bool     `yaml:"hide"`        // Tells whether the change should be hidden from the changelog
}

// NewType returns a new Type instance
func NewType(code TypeCode, title, description string, hide bool) *Type {
	return &Type{
		Title:       title,
		Code:        code,
		Description: description,
		Hide:        hide,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// ModuleCode represents the Code of the module
type ModuleCode string

var (
	ModuleNone = NewModule("none", "None", "none")
)

// Module contains the details of a single module
type Module struct {
	Code        ModuleCode `yaml:"id"`          // Code of the module
	Title       string     `yaml:"title"`       // Title of the module to be used inside the changelog
	Description string     `yaml:"description"` // Description of the module to be used inside the selector
}

// NewModule allows to build a new Module instance
func NewModule(id ModuleCode, title, description string) *Module {
	return &Module{
		Code:        id,
		Title:       title,
		Description: description,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// Entry represents a single changeset entry
type Entry struct {
	Type                 TypeCode   `yaml:"type"`
	Module               ModuleCode `yaml:"module"`
	PullRequestID        int        `yaml:"pull_request"`
	Description          string     `yaml:"description"`
	IsBackwardCompatible bool       `yaml:"backward_compatible"`
	Time                 time.Time  `yaml:"date"`
}

// NewEntry allows to build a new Entry instance
func NewEntry(
	typeCode TypeCode, moduleID ModuleCode, pr int, description string, backwardCompatible bool, time time.Time,
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

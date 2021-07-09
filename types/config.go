package types

import (
	"fmt"
	"strings"
)

// Config contains the data of the configuration
type Config struct {
	GitHubRepo     string    `yaml:"github_repo"`
	CurrentVersion *Version  `yaml:"version"`
	Types          []*Type   `yaml:"types"`
	Modules        []*Module `yaml:"modules"`
}

// DefaultConfig returns the default instance of the configuration
func DefaultConfig(githubRepo string, version *Version) *Config {
	return &Config{
		GitHubRepo:     githubRepo,
		CurrentVersion: version,
		Types: []*Type{
			NewType(CategoryChange, "added", "Added a new feature"),
			NewType(CategoryChange, "changed", "Changed a feature"),
			NewType(CategoryChange, "deprecated", "Deprecated a feature"),
			NewType(CategoryChange, "removed", "Removed a feature"),
			NewType(CategoryFix, "fixed", "Fixed a bug"),
			NewType(CategoryFix, "security", "Fix a security issue"),
		},
		Modules: []*Module{
			NewModule("External", "External"),
		},
	}
}

// GetTypeByCode returns the type having the given code, or an error if such type cannot be found
func (c *Config) GetTypeByCode(code TypeCode) (*Type, error) {
	for _, t := range c.Types {
		if t.Code == code {
			return t, nil
		}
	}
	return nil, fmt.Errorf("invalid type code: %s", code)
}

// GetModuleByID returns the module having the given ID, or an error if such type cannot be found
func (c *Config) GetModuleByID(id ModuleID) (*Module, error) {
	for _, m := range c.Modules {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, fmt.Errorf("invalid module id: %s", id)
}

// Validate returns an error if the config contains some invalid values
func (c *Config) Validate() error {
	if strings.TrimSpace(c.GitHubRepo) == "" {
		return fmt.Errorf("invalid GitHub repo URL")
	}

	for _, t := range c.Types {
		err := t.Validate()
		if err != nil {
			return fmt.Errorf("invalid types %s: %s", t.Code, err)
		}
	}

	return nil
}

package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	DefaultTypes []*Type
)

const (
	TypeCodeFeat        = "feat"
	TypeCodeFix         = "fix"
	TypeCodePerformance = "perf"
	TypeCodeRefactor    = "refactor"
	TypeCodeRevert      = "revert"
)

func init() {
	type CommitType struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var typesData map[string]CommitType
	err := json.Unmarshal([]byte(typesJSON), &typesData)
	if err != nil {
		panic(err)
	}

	for code, commitType := range typesData {
		shouldBeHidden := code != TypeCodeFeat && code != TypeCodeFix && code != TypeCodePerformance && code != TypeCodeRevert
		DefaultTypes = append(DefaultTypes, NewType(TypeCode(code), commitType.Title, commitType.Description, shouldBeHidden))
	}
}

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
		Types:          DefaultTypes,
		Modules:        []*Module{},
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

// GetModuleByCode returns the module having the given Code, or an error if such type cannot be found
func (c *Config) GetModuleByCode(code ModuleCode) (*Module, error) {
	if code == ModuleNone.Code {
		return nil, nil
	}

	for _, m := range c.Modules {
		if m.Code == code {
			return m, nil
		}
	}
	return nil, fmt.Errorf("invalid module code: %s", code)
}

// Validate returns an error if the config contains some invalid values
func (c *Config) Validate() error {
	if strings.TrimSpace(c.GitHubRepo) == "" {
		return fmt.Errorf("invalid GitHub repo URL")
	}

	return nil
}

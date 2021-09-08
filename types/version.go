package types

import "fmt"

// Version contains the data of a specific software version
type Version struct {
	Major int
	Minor int
	Patch int
}

// NewVersion returns a new Version instance
func NewVersion(major, minor, patch int) *Version {
	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

// NextMajor returns the next major version starting from the current one
func (v *Version) NextMajor() *Version {
	return NewVersion(v.Major+1, 0, 0)
}

// NextMinor returns the next minor version starting from the current one
func (v *Version) NextMinor() *Version {
	return NewVersion(v.Major, v.Minor+1, 0)
}

// NextPatch returns the next patch version starting from the current one
func (v *Version) NextPatch() *Version {
	return NewVersion(v.Major, v.Minor, v.Patch+1)
}

// String implements fmt.Stringer
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// MarshalYAML implements yaml.Marshaler
func (v *Version) MarshalYAML() (interface{}, error) {
	return v.String(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler
func (v *Version) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	err := unmarshal(&value)
	if err != nil {
		return err
	}

	version, err := ParseVersion(value)
	if err != nil {
		return err
	}

	*v = *version
	return nil
}

// ParseVersion parses the given version from a string
func ParseVersion(version string) (*Version, error) {
	var major, minor, patch int
	_, err := fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	if err != nil {
		return nil, fmt.Errorf("invalid version format. Must be x.y.z")
	}

	return NewVersion(major, minor, patch), nil
}

// GetNextVersion returns the next version given the current config and changeset entries
func GetNextVersion(config *Config, entries []*Entry) (*Version, error) {
	var containsBreakingChanges, containsBackChanges, containsBackBugFixes bool
	for _, entry := range entries {
		if !entry.IsBackwardCompatible {
			containsBreakingChanges = true
		}

		entryType, err := config.GetTypeByCode(entry.Type)
		if err != nil {
			return nil, err
		}

		if entryType.Code == TypeCodeRefactor {
			containsBackChanges = true
		}

		if entryType.Code == TypeCodeFix {
			containsBackBugFixes = true
		}
	}

	currentVersion := config.CurrentVersion
	if containsBreakingChanges {
		return currentVersion.NextMajor(), nil
	} else if containsBackChanges {
		return currentVersion.NextMinor(), nil
	} else if containsBackBugFixes {
		return currentVersion.NextPatch(), nil
	}

	return currentVersion, nil
}

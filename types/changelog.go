package types

// ModuleChanges represents all the changes associated to a module
type ModuleChanges map[ModuleCode][]*Entry

func (m ModuleChanges) GetEntriesByModule(code ModuleCode) []*Entry {
	for moduleCode, moduleEntries := range m {
		if moduleCode == code {
			return moduleEntries
		}
	}
	return nil
}

// TypeChanges represents all the changes associated to a specific type
type TypeChanges map[TypeCode]ModuleChanges

// ChangeLog contains the details of a version that will be added to the Changelog
type ChangeLog struct {
	Version *Version
	Changes TypeChanges
}

// NewChangeLog allows to build a new ChangeLog version
func NewChangeLog(version *Version, changes TypeChanges) *ChangeLog {
	return &ChangeLog{
		Version: version,
		Changes: changes,
	}
}

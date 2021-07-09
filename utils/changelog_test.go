package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
)

func TestCollectChanges(t *testing.T) {
	var entries = []*types.Entry{
		types.NewEntry("added", "profiles", 10, "Added profiles", false),
		types.NewEntry("added", "profiles", 11, "Added other profiles", false),
		types.NewEntry("added", "subspaces", 20, "Added subspace", false),
		types.NewEntry("removed", "posts", 30, "Removed posts", false),
	}

	changes := utils.CollectChanges(entries)
	require.Equal(t, changes, types.TypeChanges{
		"added": types.ModuleChanges{
			"profiles": {
				types.NewEntry("added", "profiles", 10, "Added profiles", false),
				types.NewEntry("added", "profiles", 11, "Added other profiles", false),
			},
			"subspaces": {
				types.NewEntry("added", "subspaces", 20, "Added subspace", false),
			},
		},
		"removed": types.ModuleChanges{
			"posts": {
				types.NewEntry("removed", "posts", 30, "Removed posts", false),
			},
		},
	})
}

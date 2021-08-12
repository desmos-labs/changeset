package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
)

func TestCollectChanges(t *testing.T) {
	date := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
	var entries = []*types.Entry{
		types.NewEntry("added", "profiles", 10, "Added profiles", false, date),
		types.NewEntry("added", "profiles", 11, "Added other profiles", false, date),
		types.NewEntry("added", "subspaces", 20, "Added subspace", false, date),
		types.NewEntry("removed", "posts", 30, "Removed posts", false, date),
	}

	changes := utils.CollectChanges(entries)
	require.Equal(t, changes, types.TypeChanges{
		"added": types.ModuleChanges{
			"profiles": {
				types.NewEntry("added", "profiles", 10, "Added profiles", false, date),
				types.NewEntry("added", "profiles", 11, "Added other profiles", false, date),
			},
			"subspaces": {
				types.NewEntry("added", "subspaces", 20, "Added subspace", false, date),
			},
		},
		"removed": types.ModuleChanges{
			"posts": {
				types.NewEntry("removed", "posts", 30, "Removed posts", false, date),
			},
		},
	})
}

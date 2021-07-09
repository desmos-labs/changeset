package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/desmos-labs/changeset/types"
)

func TestVersion_MarshalUnmarshal(t *testing.T) {
	version := types.NewVersion(1, 0, 0)

	bz, err := yaml.Marshal(version)
	require.NoError(t, err)
	require.Equal(t, "1.0.0", string(bz))

	var serialized types.Version
	err = yaml.Unmarshal(bz, &serialized)
	require.NoError(t, err)

	require.Equal(t, version, &serialized)
}

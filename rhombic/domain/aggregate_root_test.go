package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAggregateRoot_GetAggregateRootID(t *testing.T) {
	rootID := "NewAggregateRoot"
	aggregate := NewAggregateRoot(rootID)
	require.Equal(t, aggregate.GetAggregateRootID(), rootID)
}

func TestAggregateRoot_SetAggregateRootID(t *testing.T) {
	rootID := ""
	aggregate := NewAggregateRoot(rootID)
	id := "SetAggregateRootID"
	aggregate.SetAggregateRootID(id)
	require.Equal(t, aggregate.GetAggregateRootID(), id)
}

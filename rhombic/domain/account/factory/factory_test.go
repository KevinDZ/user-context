package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactory_InstanceOf(t *testing.T) {
	rootID := ""
	require.Error(t, InstanceAccountAggregate(rootID).InstanceOf(), "")
	return
}

func TestFactory_Registered(t *testing.T) {
	rootID := ""
	require.Error(t, InstanceAccountAggregate(rootID).Registered(), "")
	return
}

package service

import (
	"github.com/stretchr/testify/require"
	"testing"
	"user-context/rhombic/domain/account/entity"
)

func TestNewAccountService(t *testing.T) {
	require.Nil(t, NewAccountService(), "service is nil")
}

func TestService_Registered(t *testing.T) {
	require.Error(t, NewAccountService().Registered(entity.Entity{ID: "rootID"}), "Register failed")
}

func TestService_Query(t *testing.T) {
	rootID := ""
	require.Equal(t, NewAccountService().Query(rootID).ID, rootID, "Query is failed")
}

func TestService_Verify(t *testing.T) {
	rootID := ""
	require.Error(t, NewAccountService().Verify(entity.Entity{ID: rootID}), "Verify is failed")
}

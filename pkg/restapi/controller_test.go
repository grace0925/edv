/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package restapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hyperledger/aries-framework-go/pkg/mock/storage"
	"github.com/stretchr/testify/require"

	"github.com/trustbloc/edv/pkg/edvprovider/memedvprovider"
	"github.com/trustbloc/edv/pkg/restapi/operation"
)

func TestController_New(t *testing.T) {
	controller, err := New(&operation.Config{Provider: memedvprovider.NewProvider()})
	require.NoError(t, err)
	require.NotNil(t, controller)

	controller, err = New(&operation.Config{StorageProvider: &storage.MockStoreProvider{
		ErrOpenStoreHandle: fmt.Errorf("failed to open")}, AuthEnable: true})
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to open")
	require.Nil(t, controller)
}

func TestController_GetOperations(t *testing.T) {
	controller, err := New(&operation.Config{Provider: memedvprovider.NewProvider()})
	require.NoError(t, err)
	require.NotNil(t, controller)

	ops := controller.GetOperations()

	require.Equal(t, 7, len(ops))

	require.Equal(t, "/encrypted-data-vaults", ops[0].Path())
	require.Equal(t, http.MethodPost, ops[0].Method())
	require.NotNil(t, ops[0].Handle())

	require.Equal(t, "/encrypted-data-vaults/{vaultID}/query", ops[1].Path())
	require.Equal(t, http.MethodPost, ops[1].Method())
	require.NotNil(t, ops[1].Handle())

	require.Equal(t, "/encrypted-data-vaults/{vaultID}/documents", ops[2].Path())
	require.Equal(t, http.MethodPost, ops[2].Method())
	require.NotNil(t, ops[2].Handle())

	require.Equal(t, "/encrypted-data-vaults/{vaultID}/documents", ops[3].Path())
	require.Equal(t, http.MethodGet, ops[3].Method())
	require.NotNil(t, ops[3].Handle())

	require.Equal(t, "/encrypted-data-vaults/{vaultID}/documents/{docID}", ops[4].Path())
	require.Equal(t, http.MethodGet, ops[4].Method())
	require.NotNil(t, ops[4].Handle())
}

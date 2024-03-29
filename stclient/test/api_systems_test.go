/*
SpaceTraders API

Testing SystemsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package stclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_stclient_SystemsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemsApiService GetJumpGate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsApi.GetJumpGate(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GetMarket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsApi.GetMarket(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GetShipyard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsApi.GetShipyard(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GetSystem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string

		resp, httpRes, err := apiClient.SystemsApi.GetSystem(context.Background(), systemSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GetSystemWaypoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string

		resp, httpRes, err := apiClient.SystemsApi.GetSystemWaypoints(context.Background(), systemSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GetSystems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemsApi.GetSystems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GetWaypoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsApi.GetWaypoint(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

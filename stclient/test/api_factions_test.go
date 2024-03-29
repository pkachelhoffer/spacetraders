/*
SpaceTraders API

Testing FactionsApiService

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

func Test_stclient_FactionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FactionsApiService GetFaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var factionSymbol string

		resp, httpRes, err := apiClient.FactionsApi.GetFaction(context.Background(), factionSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FactionsApiService GetFactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FactionsApi.GetFactions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

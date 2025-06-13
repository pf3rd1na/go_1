package api

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/config"
)

func getAPIKey(t *testing.T) string {
	cfgPath := "../config/.env"
	key := config.NewConfig(&cfgPath).Key
	require.NotEmpty(t, key, "API key not set in .env file")
	return key
}

func TestCreateBinAPI(t *testing.T) {
	apiKey := getAPIKey(t)
	// Arrange
	bin := &bins.Bin{
		Private:   false,
		CreatedAt: time.Now(),
		Name:      "test-bin-create",
	}
	// Act
	binID, err := CreateBinAPI(apiKey, bin)
	// Assert
	require.NoError(t, err, "CreateBinAPI failed")
	assert.NotEmpty(t, binID, "CreateBinAPI returned empty bin ID")
	// Clean up
	err = DeleteBinAPI(apiKey, binID)
	assert.NoError(t, err, "Cleanup failed")
}

func TestUpdateBinAPI(t *testing.T) {
	apiKey := getAPIKey(t)
	// Arrange
	bin := &bins.Bin{
		Private:   false,
		CreatedAt: time.Now(),
		Name:      "test-bin-update",
	}
	binID, err := CreateBinAPI(apiKey, bin)
	require.NoError(t, err, "Failed to create bin for update")
	bin.Name = "test-bin-updated"
	// Act
	err = UpdateBinAPI(apiKey, binID, bin)
	// Assert
	require.NoError(t, err, "UpdateBinAPI failed")
	// Clean up
	err = DeleteBinAPI(apiKey, binID)
	assert.NoError(t, err, "Cleanup failed")
}

func TestGetBinAPI(t *testing.T) {
	apiKey := getAPIKey(t)
	// Arrange
	bin := &bins.Bin{
		Private:   false,
		CreatedAt: time.Now(),
		Name:      "test-bin-get",
	}
	binID, err := CreateBinAPI(apiKey, bin)
	require.NoError(t, err, "Failed to create bin for get")
	// Act
	data, err := GetBinAPI(apiKey, binID)
	// Assert
	require.NoError(t, err, "GetBinAPI failed")
	var got struct {
		Record bins.Bin `json:"record"`
	}
	err = json.Unmarshal(data, &got)
	require.NoError(t, err, "Failed to unmarshal bin")
	assert.Equal(t, bin.Name, got.Record.Name, "Expected name %s, got %s", bin.Name, got.Record.Name)
	// Clean up
	err = DeleteBinAPI(apiKey, binID)
	assert.NoError(t, err, "Cleanup failed")
}

func TestDeleteBinAPI(t *testing.T) {
	apiKey := getAPIKey(t)
	// Arrange
	bin := &bins.Bin{
		Private:   false,
		CreatedAt: time.Now(),
		Name:      "test-bin-delete",
	}
	binID, err := CreateBinAPI(apiKey, bin)
	require.NoError(t, err, "Failed to create bin for delete")
	// Act
	err = DeleteBinAPI(apiKey, binID)
	// Assert
	require.NoError(t, err, "DeleteBinAPI failed")
	// Try to get deleted bin (should fail)
	_, err = GetBinAPI(apiKey, binID)
	assert.Error(t, err, "Expected error when getting deleted bin, got nil")
}

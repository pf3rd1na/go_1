package api

import (
	"encoding/json"
	"testing"
	"time"

	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/config"
)

func TestBinAPI(t *testing.T) {
	cfgPath := "../config/.env"
	apiKey := config.NewConfig(&cfgPath).Key

	t.Run("CreateBinAPI", func(t *testing.T) {
		// Arrange
		bin := &bins.Bin{
			Private:   false,
			CreatedAt: time.Now(),
			Name:      "test-bin-create",
		}
		// Act
		binID, err := CreateBinAPI(apiKey, bin)
		// Assert
		if err != nil {
			t.Fatalf("CreateBinAPI failed: %v", err)
		}
		if binID == "" {
			t.Fatal("CreateBinAPI returned empty bin ID")
		}
		// Clean up
		if err := DeleteBinAPI(apiKey, binID); err != nil {
			t.Errorf("Cleanup failed: %v", err)
		}
	})

	t.Run("UpdateBinAPI", func(t *testing.T) {
		// Arrange
		bin := &bins.Bin{
			Private:   false,
			CreatedAt: time.Now(),
			Name:      "test-bin-update",
		}
		binID, err := CreateBinAPI(apiKey, bin)
		if err != nil {
			t.Fatalf("Failed to create bin for update: %v", err)
		}
		bin.Name = "test-bin-updated"
		// Act
		err = UpdateBinAPI(apiKey, binID, bin)
		// Assert
		if err != nil {
			t.Fatalf("UpdateBinAPI failed: %v", err)
		}
		// Clean up
		if err := DeleteBinAPI(apiKey, binID); err != nil {
			t.Errorf("Cleanup failed: %v", err)
		}
	})

	t.Run("GetBinAPI", func(t *testing.T) {
		// Arrange
		bin := &bins.Bin{
			Private:   false,
			CreatedAt: time.Now(),
			Name:      "test-bin-get",
		}
		binID, err := CreateBinAPI(apiKey, bin)
		if err != nil {
			t.Fatalf("Failed to create bin for get: %v", err)
		}
		// Act
		data, err := GetBinAPI(apiKey, binID)
		// Assert
		if err != nil {
			t.Fatalf("GetBinAPI failed: %v", err)
		}
		var got struct {
			Record bins.Bin `json:"record"`
		}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("Failed to unmarshal bin: %v", err)
		}
		if got.Record.Name != bin.Name {
			t.Errorf("Expected name %s, got %s", bin.Name, got.Record.Name)
		}
		// Clean up
		if err := DeleteBinAPI(apiKey, binID); err != nil {
			t.Errorf("Cleanup failed: %v", err)
		}
	})

	t.Run("DeleteBinAPI", func(t *testing.T) {
		// Arrange
		bin := &bins.Bin{
			Private:   false,
			CreatedAt: time.Now(),
			Name:      "test-bin-delete",
		}
		binID, err := CreateBinAPI(apiKey, bin)
		if err != nil {
			t.Fatalf("Failed to create bin for delete: %v", err)
		}
		// Act
		err = DeleteBinAPI(apiKey, binID)
		// Assert
		if err != nil {
			t.Fatalf("DeleteBinAPI failed: %v", err)
		}
		// Try to get deleted bin (should fail)
		_, err = GetBinAPI(apiKey, binID)
		if err == nil {
			t.Error("Expected error when getting deleted bin, got nil")
		}
	})
}

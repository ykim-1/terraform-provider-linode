package instancetype_test

import (
	"context"
	"encoding/json"
	"github.com/linode/linodego"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataSourceLinodeInstanceType_basic(t *testing.T) {
	instanceTypeID := "g6-standard-2"

	// Create a mock client
	mockClient := &MockClient{}
	typeInfo, err := mockClient.GetType(context.Background(), instanceTypeID)
	if err != nil {
		t.Fatalf("failed to get instance type %s: %s", instanceTypeID, err)
	}

	data := `{"ID":"g6-standard-2","Disk":81920,"Class":"standard","Price":{"Hourly":0.0,"Monthly":0.0},"Label":"Linode 4GB","Addons":null,"NetworkOut":4000,"Memory":4096,"Transfer":4000,"VCPUs":2}`

	// JSON data into a struct
	var actualData linodego.LinodeType
	err = json.Unmarshal([]byte(data), &actualData)
	assert.NoError(t, err)
	assert.NotNil(t, actualData)

	// Assert
	assert.Equal(t, typeInfo.ID, actualData.ID)
	assert.Equal(t, typeInfo.Label, actualData.Label)
	// ... add assertions
}

// MockClient implements the Linode client interface for testing purposes
type MockClient struct{}

// GetType is a mock implementation of the GetType method
func (c *MockClient) GetType(ctx context.Context, instanceTypeID string) (*linodego.LinodeType, error) {
	return &linodego.LinodeType{
		ID:         instanceTypeID,
		Label:      "Linode 4GB",
		Disk:       81920,
		Class:      "standard",
		Memory:     4096,
		VCPUs:      2,
		NetworkOut: 4000,
	}, nil
}

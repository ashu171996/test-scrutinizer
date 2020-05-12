package handler

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/magiconair/properties/assert"
)

type mockCloudWatchClient struct {
	cloudwatchiface.CloudWatchAPI
}

func (m *mockCloudWatchClient) PutMetricData(input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	Message := &cloudwatch.PutMetricDataOutput{}
	return Message, nil
}
func TestPutmetricData(t *testing.T) {
	mockSvc := &mockCloudWatchClient{}
	_, err := PutmetricData(mockSvc, "Test", "Test", "Count", 11.09, "Test", "Test")
	assert.Equal(t, nil, err)

}

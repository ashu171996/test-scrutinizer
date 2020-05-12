package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// PutmetricData : send metric over cloudwatch
func PutmetricData(client cloudwatchiface.CloudWatchAPI, metricName string, namesapce string, unit string, value float64, DimensionName string, DimensionValue string) (*cloudwatch.PutMetricDataOutput, error) {
	Res, err := client.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String(namesapce),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String(metricName),
				Unit:       aws.String(unit),
				Value:      aws.Float64(value),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String(DimensionName),
						Value: aws.String(DimensionValue),
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return Res, nil
}

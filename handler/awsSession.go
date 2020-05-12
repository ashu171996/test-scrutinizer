package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"

	"test-scrutinizer/config"
	"test-scrutinizer/log"
)

// CloudClient : This variable store common AWS client for cloudwatch
var CloudClient cloudwatchiface.CloudWatchAPI
var err error

// AWSClientCloudwatch : Function which return client for cloudwatch
func AWSClientCloudwatch() cloudwatchiface.CloudWatchAPI {
	return CloudClient
}
func init() {
	localConfig := config.Config()
	CloudClient, err = CreateCloudClient(localConfig.GetString("AWS_REGION"))
	if err != nil {
		log.Fatalf("Could not create AWS client for cloudwatch : %s", err)
	}
}

// CreateCloudClient : This function create session and client for AWS cloudwatch.
func CreateCloudClient(region string) (cloudwatchiface.CloudWatchAPI, error) {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	svccloud := cloudwatch.New(sess)
	if err != nil {
		return nil, err
	}
	return svccloud, nil
}

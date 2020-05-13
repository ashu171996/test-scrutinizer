package cmd

import (
	"fmt"
	"test-scrutinizer/config"
	database "test-scrutinizer/driver"
	"test-scrutinizer/log"

	handle "test-scrutinizer/handler"
	"test-scrutinizer/model"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "test-scrutinizer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: scrutinizerCmd,
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
func scrutinizerCmd(cmd *cobra.Command, args []string) {
	localConfig := config.Config()
	dbC, err := database.DbConn(localConfig)
	if err != nil {
		log.Fatalf("could not connect to db : %s", err)
	}
	defer dbC.Close()
	object := model.NewDummyStructure(dbC)
	object.InsertNewRecords("Ashutosh", "Dhanbad")
	CloudwatchClient := handle.AWSClientCloudwatch()
	_, err = handle.PutmetricData(CloudwatchClient, "Scrutinizer-Test", "Test-Scrutinizer", "Count", 21.32, "Name", "value")
	if err != nil {
		log.Fatalf("Could not connect to AWS : %s", err)
	}
	log.Info("Metrics successfully sent to cloudwatch")

	recursion()
}

func Abc() {
	localConfig := config.Config()
	dbC, err := database.DbConn(localConfig)
	if err != nil {
		log.Fatalf("could not connect to db : %s", err)
	}
	defer dbC.Close()
	object := model.NewDummyStructure(dbC)
	fmt.Println("object", object)
}
func recursion() {
	recursion() /* function calls itself */
}

/*
Copyright © 2022 Abbas Gheydi

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	KAFKA_SERVER    *string
	TOPICS          *[]string
	GROUP_ID        *string
	AutoOffsetReset *string
	ProduceCount    *int
	SessionTimeOut  *int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafka-tiny-tools",
	Short: "A tiny tool to consume and produce msg into kafka",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	KAFKA_SERVER = rootCmd.PersistentFlags().String("server", "localhost", "kafka server address")
	GROUP_ID = rootCmd.PersistentFlags().String("groupid", "myGroup", "consumer group name")
	AutoOffsetReset = rootCmd.PersistentFlags().String("autooffsetreset", "earliest", "autoOffsetReset policy")
	ProduceCount = rootCmd.PersistentFlags().Int("producecount", 1, "number of msgs to produce")
	SessionTimeOut = rootCmd.PersistentFlags().Int("sessiontimeout", 45000, "number of msgs to produce")

	TOPICS = rootCmd.PersistentFlags().StringSlice("topics", []string{"test"}, "comma separated list of topics")

}

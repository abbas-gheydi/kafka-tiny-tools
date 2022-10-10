/*
Copyright © 2022 Abbas Gheydi


*/
package cmd

import (
	"github.com/Abbas-gheydi/kafka-tiny-tools/internal/kafkat"
	"github.com/spf13/cobra"
)

// info represents the topicInfo command
var info = &cobra.Command{
	Use:   "info",
	Short: "list topics and brokers",

	Run: func(cmd *cobra.Command, args []string) {
		kafkat.TopicInfo(*TOPICS, *KAFKA_SERVER, *GROUP_ID, *AutoOffsetReset, *SessionTimeOut)
	},
}

func init() {
	rootCmd.AddCommand(info)

}

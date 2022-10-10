/*
Copyright © 2022 Abbas Gheydi


*/
package cmd

import (
	"fmt"

	"github.com/Abbas-gheydi/kafka-tiny-tools/internal/kafkat"
	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "consume messages",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consume called")
		kafkat.Consumer(*TOPICS, *KAFKA_SERVER, *GROUP_ID, *AutoOffsetReset, *SessionTimeOut)
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)

}

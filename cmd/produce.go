/*
Copyright © 2022 Abbas Gheydi


*/
package cmd

import (
	"fmt"

	"github.com/Abbas-gheydi/kafka-tiny-tools/internal/kafkat"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "produce messages",

	Run: func(cmd *cobra.Command, args []string) {

		topics := *TOPICS
		fmt.Println("produce called", topics[0])
		kafkat.Producer(*KAFKA_SERVER, topics[0], *ProduceCount)
	},
}

func init() {
	rootCmd.AddCommand(produceCmd)

}

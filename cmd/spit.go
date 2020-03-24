package cmd

import (
	"github.com/commitdev/kafka-connect-cli/pkg/spit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(spitCmd)
}

var spitCmd = &cobra.Command{
	Use:   "spit",
	Short: "Spit out what the actual kafka connect config would look like.",
	Run:   Spit,
}

func Spit(cmd *cobra.Command, args []string) {

	spit.Spit(kafkaConnectorConfig)
}

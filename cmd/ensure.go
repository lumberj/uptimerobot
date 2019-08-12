package cmd

import (
	"fmt"
	"log"

	uptimerobot "github.com/bitfield/uptimerobot/pkg"
	"github.com/spf13/cobra"
)

var ensureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "add a new monitor if not present",
	Long:  `Create a new monitor with the specified URL and friendly name, if the monitor does not already exist`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		m := uptimerobot.Monitor{
			URL:           args[0],
			FriendlyName:  args[1],
			Type:          uptimerobot.TypeHTTP,
			AlertContacts: contacts,
		}
		result, err := client.EnsureMonitor(m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Monitor ID %d\n", result.ID)
	},
}

func init() {
	ensureCmd.Flags().StringSliceVarP(&contacts, "contacts", "c", []string{}, "Comma-separated list of contact IDs to notify")
	RootCmd.AddCommand(ensureCmd)
}

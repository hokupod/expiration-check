/*
Copyright © 2022 Hokuto TAKEMIYA <hokupod@outlook.com>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/hokupod/expiration-check/expiration-check/whois"
	"github.com/spf13/cobra"
)

// whoisCmd represents the expiration command
var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "Extracts expiration dates for whois",
	Long:  `Extracts expiration dates from the results of whois queries.`,
	Run: func(cmd *cobra.Command, args []string) {
		expirationDate, err := whois.ExpirationDate(args[0])
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Printf("expirationDate: %v", expirationDate)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires domain")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(whoisCmd)

	whoisCmd.Flags().BoolP("duration", "d", false, "Returns the number of days until the expiration date.")
}
/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hokupod/expiration-check/holder"
	"github.com/hokupod/expiration-check/holder/ssl"
	"github.com/hokupod/expiration-check/holder/whois"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Extracts expiration dates for all supported source",
	Long:  `Extracts expiration dates for all supported source.(JSON output)`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sh ssl.Holder
			wh whois.Holder
		)

		h := holder.ExpirationCheckerNew(args[0])
		h.AddHolder(sh)
		h.AddHolder(wh)
		res := h.RunAll()

		jsonStr, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		var buf bytes.Buffer
		err = json.Indent(&buf, []byte(jsonStr), "", "  ")
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Println(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}

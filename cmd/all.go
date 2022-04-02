/*
Copyright © 2022 hokupod <hokupod@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/hokupod/expiration-check/expchk"
	"github.com/hokupod/expiration-check/expchk/domain"
	"github.com/hokupod/expiration-check/expchk/ssl"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Extracts expiration dates for all supported source",
	Long: `Extracts expiration dates for all supported source.(JSON output)

Example for:
  expiration-check all example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sh ssl.Holder
			dh domain.Holder
		)

		ec := expchk.New(args[0])
		ec.AddHolder(sh)
		ec.AddHolder(dh)
		res := ec.Run()
		for _, ex := range res.Expirations {
			if ex.Error != nil {
				fmt.Printf("Error: %v: %v\n", ex.Name, ex.Error)
				os.Exit(1)
			}
		}

		jsonStr, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}

		var buf bytes.Buffer
		err = json.Indent(&buf, []byte(jsonStr), "", "  ")
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
		fmt.Println(buf.String())
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires domain")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}

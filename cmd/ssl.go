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
	"errors"
	"fmt"
	"os"

	"github.com/hokupod/expiration-check/expchk"
	"github.com/hokupod/expiration-check/expchk/ssl"
	"github.com/spf13/cobra"
)

// sslCmd represents the ssl command
var sslCmd = &cobra.Command{
	Use:   "ssl",
	Short: "Extracts expiration dates for ssl",
	Long: `Extracts expiration dates from the results of ssl queries.

Example for:
  expiration-check ssl [-d] example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		var sh ssl.Holder

		h := expchk.New(args[0])
		h.AddHolder(sh)
		res := h.Run()
		errors := res.Expirations[0].Errors
		if errors != nil {
			for _, err := range errors {
				fmt.Printf("Error: %v: %v\n", res.Expirations[0].Name, err)
			}
			os.Exit(1)
		}

		if o.durationFlg {
			fmt.Println(*res.Expirations[0].Duration)
		} else {
			fmt.Println(res.Expirations[0].ExpirationDate)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires domain")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sslCmd)

	sslCmd.Flags().BoolVarP(&o.durationFlg, "duration", "d", false, "Returns the number of days until the expiration date.")
}

// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(htmlEncodeCmd)
}

var htmlEncodeCmd = &cobra.Command{
	Use:     "html-encode",
	Short:   "Escape your HTML",
	Aliases: []string{"html-enc", "html-escape"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		in, out := "", ""

		if len(args) == 0 {
			all, err := ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
			in = string(all)
		} else {
			in = args[0]
		}

		p := processors.HTMLEncode{}
		flags := make([]processors.Flag, 0)

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}

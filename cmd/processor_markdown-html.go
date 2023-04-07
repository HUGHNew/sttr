// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/abhimanyu003/sttr/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markdownHtmlCmd)
}

var markdownHtmlCmd = &cobra.Command{
	Use:     "markdown-html [string]",
	Short:   "Convert Markdown to HTML",
	Aliases: []string{"md-html"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var in []byte
		var out string

		if len(args) == 0 {
			in = []byte(utils.ReadMultilineInput())
		} else {
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := os.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = d
			} else {
				in = []byte(args[0])
			}
		}

		flags := make([]processors.Flag, 0)
		p := processors.Markdown{}

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}

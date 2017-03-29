package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yaskoo/strest/play"

	"github.com/spf13/cobra"
)

var playFile string
var testMode bool
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a play described in a YAML file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if playFile == "" {
			fmt.Println("A play file is required")
			os.Exit(1)
		}

		client := &http.Client{}
		p := play.New(playFile)
		ctx := &play.Context{
			Responses: make([]*play.Response, len(p.Steps)),
		}

		if testMode {
			p.ExecTest(ctx, client)
		} else {
			p.Exec(ctx, client)
		}
	},
}

func init() {
	RootCmd.AddCommand(playCmd)

	playCmd.Flags().StringVarP(&playFile, "file", "f", "", "YAML file to play")
	playCmd.Flags().BoolVarP(&testMode, "test", "t", false, "Play the play in test mode i.e check expectations")
}

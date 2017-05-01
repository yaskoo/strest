package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/yaskoo/strest/play"
)

var playCmd = &cobra.Command{
	Use:   "play FILE",
	Short: "Play a play described in a YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("The 'play' command requires a single argument, which is the file to play")
			os.Exit(1)
		}

		client := &http.Client{}
		ctx := &types.Context{}

		p := types.NewPlay(args[0])
		p.Play(ctx, client)
	},
}

func init() {
	RootCmd.AddCommand(playCmd)
}

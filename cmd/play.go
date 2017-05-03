package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/yaskoo/strest/play"
	"github.com/yaskoo/strest/player"
)

var playCmd = &cobra.Command{
	Use:   "play FILE",
	Short: "Play a play described in a YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("The 'play' command requires a single argument, which is the file to play")
			os.Exit(1)
		}

		var p play.Play
		if err := p.Load(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		player.New().Play(&p)
	},
}

func init() {
	RootCmd.AddCommand(playCmd)
}

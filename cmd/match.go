package cmd

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/gonnect/internal/app"
)

func init() {
	matchCmd := &cobra.Command{
		Use:     "match",
		Short:   "Runs a matchmaking worker",
		Long:    "Runs a matchmaking worker that will create a new match whenever enough players are searching",
		Example: "gonnect match --batch 5",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			batch, err := cmd.Flags().GetInt("batch")
			if err != nil {
				cmd.PrintErr(err)
				return
			}

			app.MatchPlayers(batch)
		},
	}

	matchCmd.Flags().IntP("batch", "b", 10, "Number of players per match")
	rootCmd.AddCommand(matchCmd)
}

package commands

import (
    "github.com/dabarov/zoom-joiner/pkg/meetings"
    "github.com/spf13/cobra"
)

var JoinCmd = &cobra.Command{
    Use:   "join [alias]",
    Short: "Join the Zoom meeting for the given alias",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        alias := args[0]
        meetings.JoinMeeting(alias)
    },
}


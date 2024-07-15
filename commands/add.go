package commands

import (
    "github.com/dabarov/zoom-joiner/pkg/meetings"
    "github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
    Use:   "add [alias] [meeting number] [password]",
    Short: "Add a new Zoom meeting with alias, meeting number, and password",
    Args:  cobra.ExactArgs(3),
    Run: func(cmd *cobra.Command, args []string) {
        alias := args[0]
        meetingNumber := args[1]
        password := args[2]
        meetings.AddMeeting(alias, meetingNumber, password)
    },
}


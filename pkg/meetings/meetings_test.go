package meetings

import (
    "testing"
)

func TestAddMeeting(t *testing.T) {
    alias := "team_meeting"
    meetingNumber := "123456789"
    password := "12345"

    AddMeeting(alias, meetingNumber, password)
}

func TestJoinMeeting(t *testing.T) {
    alias := "team_meeting"

    JoinMeeting(alias)
}


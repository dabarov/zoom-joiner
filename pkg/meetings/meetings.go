package meetings

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "runtime"
)

const fileName = "meetings.json"

type Meeting struct {
    MeetingNumber string `json:"meeting_number"`
    Password      string `json:"password"`
}

func AddMeeting(alias, meetingNumber, password string) {
    meetings := loadMeetings()

    meetings[alias] = Meeting{
        MeetingNumber: meetingNumber,
        Password:      password,
    }

    saveMeetings(meetings)
    fmt.Println("Meeting added successfully!")
}

func JoinMeeting(alias string) {
    meetings := loadMeetings()

    if meeting, ok := meetings[alias]; ok {
        url := fmt.Sprintf("zoommtg://zoom.us/join?confno=%s&pwd=%s", meeting.MeetingNumber, meeting.Password)
        openURL(url)
    } else {
        fmt.Println("No meeting found for alias:", alias)
    }
}

func loadMeetings() map[string]Meeting {
    file, err := ioutil.ReadFile(fileName)
    if err != nil {
        if os.IsNotExist(err) {
            return make(map[string]Meeting)
        }
        panic(err)
    }

    var meetings map[string]Meeting
    if err := json.Unmarshal(file, &meetings); err != nil {
        panic(err)
    }

    return meetings
}

func saveMeetings(meetings map[string]Meeting) {
    file, err := json.MarshalIndent(meetings, "", "  ")
    if err != nil {
        panic(err)
    }

    if err := ioutil.WriteFile(fileName, file, 0644); err != nil {
        panic(err)
    }
}

func openURL(url string) {
    var cmd *exec.Cmd

    switch runtime.GOOS {
    case "darwin":
        cmd = exec.Command("open", url)
    case "linux":
        cmd = exec.Command("xdg-open", url)
    case "windows":
        cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
    default:
        fmt.Println("Unsupported platform")
        return
    }

    err := cmd.Start()
    if err != nil {
        fmt.Println("Failed to open URL:", err)
    }
}


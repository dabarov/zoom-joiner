# Zoom Joiner CLI

Effortlessly save and join Zoom meetings using simple aliases

## Features

- Save Zoom meeting details with an alias.
- Join a Zoom meeting using the alias.
- Supports macOS, Linux, and Windows.

## Installation

### Option 1: Build from Source

#### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)

#### Steps

1. Clone the repository:

   ```sh
   git clone https://github.com/dabarov/zoom-joiner.git
   cd zoom-joiner
   ```

2. Build the project:

   ```sh
   go build -o zoom-joiner
   ```

3. (Optional) Move the binary to your PATH:

   ```sh
   sudo mv zoom-joiner /usr/local/bin/
   ```

### Option 2: Download Released Version

1. Go to the [Releases](https://github.com/dabarov/zoom-joiner/releases) page.
2. Download the appropriate binary for your operating system:
   - [zoom-joiner-darwin-amd64](https://github.com/dabarov/zoom-joiner/releases/download/v0.0.1/zoom-joiner-darwin-amd64) for macOS
   - [zoom-joiner-linux-amd64](https://github.com/dabarov/zoom-joiner/releases/download/v0.0.1/zoom-joiner-linux-amd64) for Linux
   - [zoom-joiner-windows-amd64.exe](https://github.com/dabarov/zoom-joiner/releases/download/v0.0.1/zoom-joiner-windows-amd64.exe) for Windows

#### Steps

1. Once downloaded, make the binary executable (if necessary):
   
   ```sh
   chmod +x zoom-joiner-darwin-amd64  # For macOS or Linux
   ```

2. Move the binary to a directory included in your PATH to run it from anywhere:

   ```sh
   sudo mv zoom-joiner-darwin-amd64 /usr/local/bin/zoom-joiner  # Example for macOS

   ```

   Replace /usr/local/bin/zoom-joiner with the appropriate directory for your system.

## Usage

### Adding a Meeting

To add a meeting, use the `add` command followed by the alias, meeting number, and password:

```sh
zoom-joiner add [alias] [meeting_number] [password]
```

**Example:**

```sh
zoom-joiner add team_meeting 123456789 12345
```

### Joining a Meeting

To join a meeting, use the join command followed by the alias:

```sh
zoom-joiner join [alias]
```

**Example:**

```sh
zoom-joiner join team_meeting
```

This will open the Zoom meeting in your default browser or Zoom application.


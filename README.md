# Zoom Joiner CLI

Effortlessly save and join Zoom meetings using simple aliases

## Features

- Save Zoom meeting details with an alias.
- Join a Zoom meeting using the alias.
- Supports macOS, Linux, and Windows.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)

### Steps

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


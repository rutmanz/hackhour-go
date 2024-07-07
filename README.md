# hackhour-go

Golang CLI tool and API wrapper for HackClub's [HackHour API](https://github.com/hackclub/hack-hour)


## CLI Tool
### Installation
```
$ go install github.com/rutmanz/hackhour-go
```
### Usage
First locate your HackHour API Key, or generate one using `/api` in Slack, then use it to login
```
hackhour-go login 'your-api-key'
```
To see available commands, checkout the help page.
```
hackhour-go help
```
### Demo
[![asciicast](https://asciinema.org/a/0n5osvDq0d0CUFsVyI6OBwkWS.svg)](https://asciinema.org/a/0n5osvDq0d0CUFsVyI6OBwkWS)

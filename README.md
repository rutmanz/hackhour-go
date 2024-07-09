# hackhour-go

Golang CLI tool and API wrapper for HackClub's [HackHour API](https://github.com/hackclub/hack-hour)


## CLI Tool
### Installation
```
$ go install github.com/rutmanz/hackhour-go@latest
```
### Usage
First locate your HackHour API Key, or generate one using `/api` in Slack, then use it to login
```
hackhour-go login 'your-api-key'
```

To enable the use of `hackhour-go session send`, you'll need to [create a slack application](https://api.slack.com/apps), add the `chat:write` user scope, and install it to the Hack Club workspace. Then configure hackhour-go with the token (you must have already run `hackhour-go login`) to authorize it.
```
hackhour-go authslack 'your-slack-token'
```


To see available commands, checkout the help page.
```
hackhour-go help
```

### Send to Session Thread
`hackhour-go session send` supports a `-g` flag, which automatically appends the github or gitlab link to your latest commit to the message. It uses the first git remote for the url and uses `HEAD` for the commit hash. Keep in mind this will use the current checked-out commit locally, and you will need to push the commits before people can visit the links

You can also automatically send commits as you make them, using a post-commit [git hook](https://www.atlassian.com/git/tutorials/git-hooks)

### Demo
[![asciicast](https://asciinema.org/a/0n5osvDq0d0CUFsVyI6OBwkWS.svg)](https://asciinema.org/a/0n5osvDq0d0CUFsVyI6OBwkWS)

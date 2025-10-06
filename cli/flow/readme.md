# flow-go

## Setup

This puts `f` command into your path to use. I use this CLI to do various things but you are free to create a similar CLI for your use cases.

If there is something you want to contribute to this CLI too, feel free to PR.

Current output of CLI is:

```
f h
flow is CLI to do things fast

Usage:
  flow [command]

Available Commands:
  help             Help about any command
  install-flow     Install the Flow CLI into ~/bin and optionally add it to PATH
  commit           Generate a commit message with GPT-5 nano and create the commit
  commitPush       Generate a commit message, commit, and push to the default remote
  commitReviewAndPush Generate a commit message, review it interactively, commit, and push
  branchFromClipboard Create a git branch from the clipboard name
  clone            Clone a GitHub repository into ~/gh/<owner>/<repo>
  cloneAndOpen     Clone a GitHub repository and open it in Cursor (falls back to Safari tab)
  gitCheckout      Check out a branch from the remote, creating a local tracking branch if needed
  updateGoVersion  Upgrade Go using the workspace script
  youtubeToSound   Download audio from a YouTube URL into ~/.flow/youtube-sound using yt-dlp (falls back to Safari tab)
  version          Reports the current version of flow

Flags:
  -h, --help   help for flow

Use "flow [command] --help" for more information about a command.
```

Running `f` without any arguments opens an embedded fzf palette so you can fuzzy-search commands and read their descriptions before executing them.

For `f commit`, export `OPENAI_API_KEY` in your shell profile (e.g. fish config) so the CLI can talk to OpenAI. This environment variable is the only requirement, so the command works in local shells and CI alike.

For `f youtubeToSound`, the CLI now automatically passes `--cookies-from-browser` using Safari cookies. Override this by setting `FLOW_YOUTUBE_COOKIES_BROWSER` (e.g. `firefox`), set it to `none` to skip cookies entirely, or pass your own `--cookies*` flags after the URL—they are forwarded directly to `yt-dlp`.

If you run `f youtubeToSound` without arguments, the command grabs the frontmost Safari tab URL automatically.

# Go [![Docs](https://img.shields.io/badge/-Docs-0a0a0a.svg?style=flat&colorA=0a0a0a)](https://pkg.go.dev/github.com/nikitavoloboev/go?tab=doc)

> Small Go packages & CLIs

## Install

There is [cmd](cmd) dir that contains Go CLIs. The rest are libraries.

Can install commands/libraries using:

`go get -u github.com/nikitavoloboev/go/<path-to-lib-or-cmd>`

For example: `go get -u github.com/nikitavoloboev/go/cmd/savelink` will install [this tool](cmd/savelink/main.go) to save links.

## Run

The repo consists of go libraries and CLIs written in Go. Assuming all dependencies are installed (opening repo with [Go extension](https://github.com/golang/vscode-go) & [VSCode](https://code.visualstudio.com/download) should do it automatically).

Running `go run .` in folder with `main.go` will build & run the library/tool.

## Tips

[watchexec](https://github.com/watchexec/watchexec) is a nice tool for developing. Can make this alias:

`alias wg=watchexec --exts go "echo -- && go run ."`

When `wg` is then run, it will run `go run .` automatically whenever go file changes.

## Discuss / help

Search for [existing issues](../../issues) or open [new one](../../issues/new/choose).

## Thank you

You can support me on [GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects) I shared.

I also have [personal Discord](https://discord.com/invite/TVafwaD23d) you can join for more indepth discussions.

[![MIT](http://bit.ly/mitbadge)](https://choosealicense.com/licenses/mit/) [![Twitter](http://bit.ly/nikitatweet)](https://twitter.com/nikitavoloboev)

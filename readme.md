# Go [![Docs](https://img.shields.io/badge/-Docs-0a0a0a.svg?style=flat&colorA=0a0a0a)](https://pkg.go.dev/github.com/nikitavoloboev/go?tab=doc)

> Go libraries + testing code

## File structure

- [cli](cli) - CLIs in go

## Setup

Everything is driven using [bun](https://bun.sh/) commands.

Assumes [go](https://go.dev/doc/install) is installed.

## Run

```
bun run dev
```

Runs: `go run .`

## Install CLI

There is [cli](cli) dir that contains Go CLIs.

`go get -u github.com/nikitavoloboev/go/<path-to-cmd>`

For example: `go get -u github.com/nikitavoloboev/go/cli/savelink` will install [this CLI](cli/savelink/main.go) to save links.

## Install package

> TODO:

<!-- `go get -u github.com/nikitavoloboev/go/lib/<path-to-lib>` -->

## Contribute

The tasks to do are outlined in [existing issues](../../issues) and in [tasks below](#tasks) (sorted by priority).

If issue/idea you have is not there, [open new issue](../../issues/new/choose) or [start discussion](../../discussions).

Any PR with code/doc improvements is welcome. ✨

Join [Discord](https://discord.com/invite/TVafwaD23d) for more indepth discussions on this repo and [others](https://github.com/nikitavoloboev#src).

## Tasks

- setup [goreleaser](https://goreleaser.com) and have fast setup to write, share and deploy code
  - both as libraries, CLIs, http services
- do some CLIs with [bubble tea](https://github.com/charmbracelet/bubbletea)
- move `learn.rs` into a folder in `src/learn` and have many files
  - `slices.rs`, `structs.rs` etc. rather than have it all in one `learn.rs` file

### ♥️

[Support on GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects).

[![MIT](http://bit.ly/mitbadge)](https://choosealicense.com/licenses/mit/) [![Twitter](http://bit.ly/nikitatweet)](https://twitter.com/nikitavoloboev)

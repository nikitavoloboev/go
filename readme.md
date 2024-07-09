# Go [![Docs](https://img.shields.io/badge/-Docs-0a0a0a.svg?style=flat&colorA=0a0a0a)](https://pkg.go.dev/github.com/nikitavoloboev/go?tab=doc)

> Go libraries and other code

## Files

- [cli](cli) - CLIs in go

## Setup

<!-- TODO: add nix setup -->

<!-- [Bun](https://bun.sh/) is used to run things. -->

Assumes [go](https://go.dev/doc/install) is installed.

## Run

For now I go inside each of the go projects inside. Then I have these fish functions:

```
function g
    watchexec --no-vcs-ignore --restart --exts go "tput reset && go run ."
end

function G
    watchexec --no-vcs-ignore --restart --exts go "tput reset && go test"
end
```

`g` runs the code, and `G` runs the tests. Will improve setup with time.

<!-- ```
bun dev
```

Runs: `go run .` -->

## Install CLI

There is [cli](cli) dir that contains Go CLIs.

`go get -u github.com/nikitavoloboev/go/<path-to-cmd>`

For example: `go get -u github.com/nikitavoloboev/go/cli/savelink` will install [this CLI](cli/savelink/main.go) to save links.

<!-- ## Install package

> TODO: -->

<!-- `go get -u github.com/nikitavoloboev/go/lib/<path-to-lib>` -->

## Contribute

Always open to useful ideas or fixes in form of issues or PRs.

Can [open new issue](../../issues/new/choose) (search [existing issues](../../issues) first) or [start discussion](../../discussions).

It's okay to submit draft PR as you can get help along the way to make it merge ready.

Join [Discord](https://discord.com/invite/TVafwaD23d) for more indepth discussions on this repo and [others](https://github.com/nikitavoloboev#src).

### 🖤

[Support on GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects).

[![Discord](https://img.shields.io/badge/Discord-100000?style=flat&logo=discord&logoColor=white&labelColor=black&color=black)](https://discord.com/invite/TVafwaD23d) [![X](https://img.shields.io/badge/nikitavoloboev-100000?logo=X&color=black)](https://x.com/nikitavoloboev) [![nikiv.dev](https://img.shields.io/badge/nikiv.dev-black)](https://nikiv.dev)

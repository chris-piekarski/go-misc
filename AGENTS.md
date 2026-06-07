# AGENTS.md

Context for AI agents working in this repository.

## What this is

`go-misc` — a personal collection of small, mostly independent Go exercises and
experiments (algorithms, data structures, and small HTTP services). Originally
written 2011–2013 under the old GOPATH layout; modernized to Go modules.

It is a learning/playground repo, not a production system. Keep changes small,
idiomatic, and self-contained. There is no shared framework to honor — each
package stands alone.

## Module & toolchain

- Single Go module at the root: `module github.com/chris-piekarski/go-misc`, `go 1.23`.
- **Standard library only** — the root module has no third-party dependencies, so
  there is no `go.sum`. Keep it that way unless there's a strong reason; adding a
  dependency is a notable decision, not a default.
- `go` may not be on `PATH` on every machine. It is a normal Go install when present.

## Layout

```
fibo/  pell/  sort/  tree/  myhttp/   library packages (with tests)
cmd/                                  runnable programs (package main)
  calcfibo  calcpell  myserver  gowiki  hello
legacy/                               old experiments that DO NOT build — see legacy/AGENTS.md
```

- Library packages live at the repo root, one directory per package.
- Every `package main` lives under `cmd/<name>/`.
- Import paths are module-relative, e.g. `github.com/chris-piekarski/go-misc/fibo`.

### Package map
- `fibo` — Fibonacci generator; `Fibo` takes a pluggable operator func and returns a closure.
- `pell` — Pell-number generator over a channel (`Pell(ch chan<- uint64)`).
- `sort` — selection / insertion / quick / heap sorts on `[]int`.
- `tree` — binary search tree (`BinaryNode`): insert, search, delete, min/max, successor, and in/pre/post-order walks.
- `myhttp` — serves a file's contents through ROT13; two ROT13 impls (table + modulo).
- `cmd/calcfibo`, `cmd/calcpell` — flag-driven CLIs (`-i N` iterations) over fibo/pell.
- `cmd/myserver` — HTTP server using `myhttp` (ROT13) on `:12345`.
- `cmd/gowiki` — the classic golang.org wiki tutorial (view/edit/save `.txt` pages).
- `cmd/hello` — hello world.

## Build / test / run

```sh
go build ./...
go test ./...
go test ./fibo -bench=.        # benchmarks (fibo, myhttp have them)
gofmt -l .                     # must print nothing; run `gofmt -w .` to fix

go run ./cmd/calcfibo -i 10
go run ./cmd/calcpell -i 10
go run ./cmd/myserver          # run from repo root; reads myhttp/data.html
cd cmd/gowiki && go run .      # run from its own dir so it finds *.html templates
```

`./...` does NOT include `legacy/` — it is a separate nested module (intentional).

## CI

GitHub Actions runs on every push to `master` and every PR
(`.github/workflows/ci.yml`): gofmt check, `go vet ./...`, `go build ./...`, and
`go test -race -cover ./...`. Keep all four green. The fmt check fails the build on
any non-gofmt'd file, so always run `gofmt -w .` before pushing.

## Conventions & expectations

- **Always `gofmt`** before finishing. The repo is gofmt-clean; keep it that way.
- Run `go vet ./...`, `go build ./...`, and `go test ./...` and confirm green before
  claiming a change is done. If `go` isn't installed, say so rather than guessing.
- Match the surrounding style of the file you edit (it is mostly tabs, idiomatic Go).
- New `main` programs go under `cmd/`. New libraries go at the repo root.
- Don't add third-party dependencies casually — stdlib-only is a deliberate property.
- Templates/static files (gowiki, magnetize) are loaded by relative path from the
  process working directory; that fragility is pre-existing. `//go:embed` would be a
  reasonable modernization if asked.

## Known issues (don't mistake these for your bug)

- `cmd/*` packages have no tests (0% coverage); their `main` funcs aren't unit
  tested. The library packages carry the meaningful coverage (~98–100%, tree 100%).

## Security / secrets

- Do not commit secrets. A FullContact API key was previously hardcoded in
  `legacy/magnetize/mag.go`; it now reads `FULLCONTACT_API_KEY` from the env. The old
  key still exists in git history — flag this if asked to harden, but don't rewrite
  history without explicit instruction.

## Git — required conventions

This repo enforces the following. Agents MUST follow them.

- **Signed commits are required.** Every commit must be signed (`git commit -S`,
  GPG or SSH signing). Ensure `user.signingkey` / `gpg.format` / `commit.gpgsign`
  are configured before committing; if signing isn't set up, stop and ask rather
  than producing unsigned commits.
- **Conventional Commits** for commit messages, branch names, and PR titles.
  - Commit message / PR title: `type(scope): subject`, e.g.
    `refactor(layout): migrate to go modules and cmd/ structure`.
  - Branch name: `type/short-kebab-description`, e.g. `refactor/go-modules`.
  - Common types: `feat`, `fix`, `refactor`, `chore`, `build`, `docs`, `test`, `ci`.
    Use `!` or a `BREAKING CHANGE:` footer for breaking changes.
- Default branch: `master`. Branch before committing; commit/push only when asked.
- Use `git mv` for moves so history is preserved.

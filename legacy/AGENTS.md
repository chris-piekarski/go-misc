# AGENTS.md — legacy/

**This subtree does not build. Do not try to make `go build ./...` cover it.**

## Why it's isolated

`magnetize/` and `gojson/` depend on `labix.org/v2/mgo`, an abandoned MongoDB
driver whose host no longer resolves. `magnetize/` also targets the FullContact v2
API, which has been discontinued. To keep the dead dependency out of the root
module's `go build ./...` / `go test ./...`, this directory is its **own Go
module** (`legacy/go.mod`, module path `github.com/chris-piekarski/go-misc/legacy`).
Nested modules are skipped by `./...` in the parent — that isolation is the point.

## Contents

- `magnetize/` — gift-giving web service: queues requests in MongoDB, enriches
  contacts via FullContact, emails a confirmation over SMTP. (`package main`)
- `gojson/` — minimal MongoDB insert/find example. (`package main`)

## Rules for agents

- It is expected and correct that these do not compile. Don't "fix" that with a
  build tag hack or by re-merging into the root module.
- Do not add `labix.org/v2/mgo` (or any abandoned driver) to the root module.
- To genuinely revive this code: migrate to the official driver
  `go.mongodb.org/mongo-driver` and a current contact-enrichment API. That is a
  real, separate effort that needs a running MongoDB to verify — not a quick edit.
- No secrets in source. The FullContact key now comes from `FULLCONTACT_API_KEY`.

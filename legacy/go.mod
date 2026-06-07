// Separate module so the abandoned labix.org/v2/mgo dependency stays out of
// the root module's `go build ./...`. This code is kept for reference and is
// NOT expected to build. See legacy/README.md.
module github.com/chris-piekarski/go-misc/legacy

go 1.23

require labix.org/v2/mgo v0.0.0-00010101000000-000000000000

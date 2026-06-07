# legacy

Older experiments kept for reference. **These do not build.**

They depend on `labix.org/v2/mgo`, an abandoned MongoDB driver whose host no
longer resolves, and (in `magnetize`) on the FullContact v2 API, which has been
discontinued. They live in a separate Go module so the dead dependency stays
out of the root module's `go build ./...` / `go test ./...`.

To revive them, migrate to the official driver
[`go.mongodb.org/mongo-driver`](https://github.com/mongodb/mongo-go-driver)
and a current contact-enrichment API.

- `magnetize/` — gift-giving web service: queues requests in MongoDB, enriches
  contacts via FullContact, emails a confirmation over SMTP.
- `gojson/` — minimal MongoDB insert/find example.

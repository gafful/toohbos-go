# toohbos-go
A toolbox of common go components.

Install modules:
$ go mod init

## Modules
### REST Protobuf
REST API with Protobuf requests and responses
- Client:
 - /: Sends JSON payload via POST request, converts payload to Protobuf and sends to Server.
 - Server receives the Protobuf payload, and sends it back in the response.
- Server:
 - /: POST requests receives a Protobuf payload and sends it back encoded in Protobuf.
 - /: GET request returns a DummyResponse object over Protobuf.
 - /page: Both GET and POST requests return a paged response ProtoResponsePage
- Add "?json=true" to get a JSON response instead of a Protobuf response
- /complex: Sends a POST request via Protobuf and receives a paged response via Protobuf

### Run
$ go run rest-ptbuf/cmd/client/main.go

### Dev
Run from module root:
$ protoc --go_out=. proto/facto.proto

Add new modules to the workspace:
$ go work use <module-name>

#### References:
- https://protobuf.dev/getting-started/gotutorial/
- https://dev.to/andyjessop/building-a-basic-http-server-in-go-a-step-by-step-tutorial-ma4
- https://dev.to/chandrapenugonda/vscode-multiple-go-projects-in-a-directorygo-118-3l7i
- https://www.kirandev.com/http-post-golang
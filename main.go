// We need two main method for our lsp to work
// which is initialize and shutdown for the client usage

package main

import (
    "github.com/tliron/commonlog" // for logging
    "github.com/tliron/glsp" // lsp SDK written in go
    protocol "github.com/tliron/glsp/protocol_3_16" // the protocol
    "github.com/tliron/glsp/server" // the server

    _ "github.com/tliron/commonlog/simple"
)

const lsName = "Emoji Autocomplete Language Server"

var version string = "0.0.1"
var handler protocol.Handler

func main() {
    commonlog.Configure(2, nil) // write the log to stderr

    handler = protocol.Handler{
        Initialize: initialize,
        Shutdown: shutdown,
    }

    server := server.NewServer(&handler, lsName, true)

    server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
    commonlog.NewInfoMessage(0, "Initializing server...") // log for initialize

    capabilities := handler.CreateServerCapabilities()

    return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

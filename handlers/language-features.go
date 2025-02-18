package handlers

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/wreckitral/go-lsp/mappers"

	_ "github.com/tliron/commonlog/simple"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
    var completionItems []protocol.CompletionItem

    for word, emoji := range mappers.EmojiMapper {
        emojiCopy := emoji // create copy of emoji
        completionItems = append(completionItems, protocol.CompletionItem{
            Label: word, // the label of what emoji to look for
            Detail: &emojiCopy,
            InsertText: &emojiCopy,
        })
    }

    return completionItems, nil
}

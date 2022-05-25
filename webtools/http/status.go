package http

import (
	"golang.org/x/net/html"
)

// Reference: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml

const (
	statusURL = `https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml`
)

type statusCode struct {
	id          int
	name        string
	description string
	reference   string
}

func PageLines(url string) ([]string, error) {
	r, err := GetReader(statusURL)
	if err != nil {
		return nil, err
	}

	h := html.NewTokenizer(r)
	for {
		token := h.Next()
		switch token {
		case html.ErrorToken:

		case html.TextToken:
		case html.StartTagToken:
		case html.EndTagToken:
		case html.SelfClosingTagToken:
		case html.CommentToken:
		case html.DoctypeToken:

		}

	}

	// lines := bytes.Split(b, []byte("\n"))
}

// Table Name: table-http-status-codes-1

/* html Token types
type TokenType uint32

const (
	// ErrorToken means that an error occurred during tokenization.
	ErrorToken TokenType = iota
	// TextToken means a text node.
	TextToken
	// A StartTagToken looks like <a>.
	StartTagToken
	// An EndTagToken looks like </a>.
	EndTagToken
	// A SelfClosingTagToken tag looks like <br/>.
	SelfClosingTagToken
	// A CommentToken looks like <!--x-->.
	CommentToken
	// A DoctypeToken looks like <!DOCTYPE x>
	DoctypeToken
)
*/

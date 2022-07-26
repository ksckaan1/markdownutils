package markdownutils

import (
	"bytes"

	"github.com/alecthomas/chroma/quick"
)

func highlightCode(code, lang string) (string, error) {
	buf := new(bytes.Buffer)
	err := quick.Highlight(buf, code, lang, "html", "monokai")

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

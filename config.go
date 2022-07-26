package markdownutils

import (
	"strings"

	"github.com/gomarkdown/markdown"
)

func EditMarkdown(input string) string {

	op := runComps(
		input,

		//Custom Markdown Components
		EditCodeWithoutTitle,
		EditCodeWithTitle,
		EditShortCodes,
		EditHint,

		//At last, standart md comps
		normalMdToHtml,

		//fixes
		fixHTMLChars,
	)

	return op
}

func runComps(input string, comps ...func(string) string) string {
	for _, comp := range comps {
		input = comp(input)
	}
	return input
}

func normalMdToHtml(input string) string {
	op := markdown.ToHTML([]byte(input), nil, nil)
	return string(op)
}

func fixHTMLChars(input string) string {
	input = strings.ReplaceAll(input, "&amp;lt;", "&lt;")
	input = strings.ReplaceAll(input, "&amp;gt;", "&gt;")
	return input
}

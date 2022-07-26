package markdownutils

import (
	"log"
	"regexp"
)

func EditCodeWithoutTitle(fullCode string) string {
	re := regexp.MustCompile(CodesWithoutTitle)

	op := re.ReplaceAllStringFunc(fullCode, findAllCodeWithoutTitle)

	return op
}

func findAllCodeWithoutTitle(singleCode string) string {
	s, err := inspectSingleCodeWithoutTitle(singleCode)
	if err != nil {
		log.Printf("Regexp işinde hata oluştu: %s\n", err.Error())
		return ""
	}

	return s
}

func inspectSingleCodeWithoutTitle(singleCode string) (string, error) {
	re := regexp.MustCompile(CodesWithoutTitle)

	submatchs := re.FindStringSubmatch(singleCode)

	langName := submatchs[1]
	codeArea := submatchs[2]

	htmlOutput, err := genStyledCodeWithoutTitle(codeArea, langName)
	if err != nil {
		return "", err
	}
	return htmlOutput, nil
}

func genStyledCodeWithoutTitle(codeArea, langName string) (string, error) {
	op, err := highlightCode(codeArea, langName)
	if err != nil {
		return "", err
	}

	styleRgx := regexp.MustCompile(OnlyCodeClasses)

	submatchs := styleRgx.FindStringSubmatch(op)

	htmlOutput := generateHTMLOutputWithoutTitle(langName, submatchs[1])

	return htmlOutput, nil
}

func generateHTMLOutputWithoutTitle(langName, styledCode string) string {
	op := `<div class="customcode">
<code class="lang-` + langName + `">
<pre>` + styledCode + `</pre>
</code>
</div>`
	return op
}

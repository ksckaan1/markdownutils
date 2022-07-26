package markdownutils

import (
	"fmt"
	"log"
	"regexp"
)

func EditCodeWithTitle(fullCode string) string {
	re := regexp.MustCompile(CodesWithTitle)

	op := re.ReplaceAllStringFunc(fullCode, findAllCodeWithTitle)

	fmt.Println(op)
	return op
}

func findAllCodeWithTitle(singleCode string) string {

	s, err := inspectSingleCodeWithTitle(singleCode)
	if err != nil {
		log.Printf("Regexp işinde hata oluştu: %s\n", err.Error())
		return ""
	}

	return s
}

func inspectSingleCodeWithTitle(singleCode string) (string, error) {
	re := regexp.MustCompile(CodesWithTitle)

	submatchs := re.FindStringSubmatch(singleCode)

	langName := submatchs[1]
	codeTitle := submatchs[2]
	codeArea := submatchs[3]

	htmlOutput, err := genStyledCodeWithTitle(codeArea, langName, codeTitle)
	if err != nil {
		return "", err
	}
	return htmlOutput, nil
}

func genStyledCodeWithTitle(codeArea, langName, codeTitle string) (string, error) {
	op, err := highlightCode(codeArea, langName)
	if err != nil {
		return "", err
	}

	styleRgx := regexp.MustCompile(OnlyCodeClasses)

	submatchs := styleRgx.FindStringSubmatch(op)

	htmlOutput := generateHTMLOutputWithTitle(langName, codeTitle, submatchs[1])

	return htmlOutput, nil
}

func generateHTMLOutputWithTitle(langName, codeTitle, styledCode string) string {
	op := `<div class="customcode">
<div class="title">` + codeTitle + `</div>
<code class="lang-` + langName + `">
<pre>` + styledCode + `</pre>
</code>
</div>`
	return op
}

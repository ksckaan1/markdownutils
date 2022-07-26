package markdownutils

import "regexp"

func EditHint(fullCode string) string {
	re := regexp.MustCompile(HintAreas)

	op := re.ReplaceAllStringFunc(fullCode, findAllHints)

	return op
}

func findAllHints(singleCode string) string {
	s := inspectSingleHint(singleCode)

	return s
}

func inspectSingleHint(singleCode string) string {
	re := regexp.MustCompile(HintAreas)

	submatchs := re.FindStringSubmatch(singleCode)

	hintType := submatchs[1]
	hintTitle := submatchs[2]
	hintText := submatchs[3]

	htmlOutput := genStyledHint(hintType, hintTitle, hintText)
	return htmlOutput
}

func genStyledHint(hintType, hintTitle, hintText string) string {
	op := `<div class="hint ` + hintType + `">
<div class="top">
	<div class="hintIcon"></div>
	<div class="title">` + hintTitle + `</div>
</div>

<div class="hintText">` + hintText + `</div>

</div>`

	return op
}

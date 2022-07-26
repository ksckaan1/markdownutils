package markdownutils

import (
	"regexp"
)

func EditShortCodes(fullCode string) string {
	re := regexp.MustCompile(ShortCodes)

	op := re.ReplaceAllStringFunc(fullCode, findAllShortCodes)

	return op
}

func findAllShortCodes(singleCode string) string {
	s := inspectSingleShortCode(singleCode)

	return s
}
func inspectSingleShortCode(singleCode string) string {
	re := regexp.MustCompile(ShortCodes)

	submatchs := re.FindStringSubmatch(singleCode)

	shortCodeType := submatchs[1]
	shortCodeValue := submatchs[2]

	htmlOutput := generateByShortCodeType(shortCodeType, shortCodeValue)

	return htmlOutput
}

func generateByShortCodeType(shortCodeType, shortCodeValue string) string {
	var op string
	switch shortCodeType {
	case "playground":
		op = genPlayground(shortCodeValue)
	case "youtube":
		op = genYouTubeVideo(shortCodeValue)
	}
	return op
}

func genPlayground(shortCodeValue string) string {
	op := `<div class="goplayground">
	<iframe src="https://goplay.tools/snippet/` + shortCodeValue + `" title="playground-short-code"></iframe>
</div>`
	return op
}

func genYouTubeVideo(shortCodeValue string) string {
	op := `<div class="youtube-video">
	<iframe src="https://www.youtube.com/embed/` + shortCodeValue + `" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
</div>`
	return op
}

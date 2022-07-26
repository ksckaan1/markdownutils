package markdownutils

const (

	// detect all codes in document
	// non-escaped: \x60\x60\x60([\w]+)\s([^\n]*?)\n([\s\S]*?)\x60\x60\x60$
	CodesWithTitle = "(?m)\x60\x60\x60([\\w]+)\\s([^\\n]*?)\\n([\\s\\S]*?)\x60\x60\x60$"

	// non-escaped: \x60\x60\x60([\w]+)\n([\s\S]*?)\x60\x60\x60$
	CodesWithoutTitle = "(?m)\x60\x60\x60([\\w]+)\\n([\\s\\S]*?)\x60\x60\x60$"

	OnlyCodeClasses = "(?m)<code>([\\s\\S]+)*?</code>"

	HintAreas = "(?m):::([\\w]+)\\s([^\\n]*?)\\n([\\s\\S]*?):::"

	ShortCodes = ":::([\\w]*?)\\s([\\w]*?):::"
)

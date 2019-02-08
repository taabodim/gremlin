package gremlin

const (
	singleQuote rune = '\''
	doubleQuote rune = '"'
	backslash   rune = '\\'
	pctSymbol   rune = '%'

	// Gremlin client does not allow Other characters
	ARG_REGEX = "\\p{C}+"

	// Gremlin stack defaults
	DEFAULT_MAX_CAP             = 10
	DEFAULT_MAX_GREMLIN_RETRIES = 2
	DEFAULT_VERBOSE_LOGGING     = false
	DEFAULT_PING_INTERVAL       = 5
)

var ESCAPE_CHARS_GREMLIN = CharSliceToMap([]rune{
	singleQuote,
	backslash,
	pctSymbol,
	doubleQuote,
})

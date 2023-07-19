package sess

import "errors"

var (
	ErrParsingBase64 = errors.New("base64 parsing error")
	ErrParsingJson   = errors.New("json parsing error")
	ErrCookieReading = errors.New("cookie reading error")
	ErrCookieParsing = errors.New("cookie parsing error")
)

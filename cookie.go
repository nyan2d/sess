package sess

import (
	"net/http"
	"time"
)

func ReadFromCookie(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, ErrCookieReading
	}

	session, err := ParseBase64(cookie.Value)
	if err != nil {
		return nil, ErrCookieParsing
	}

	return session, nil
}

func WriteToCookie(w http.ResponseWriter, session *Session) {
	c := http.Cookie{
		Name:    "session",
		Value:   session.Base64(),
		Path:    "/",
		Expires: time.Unix(session.Valid, 0),
	}
	http.SetCookie(w, &c)
}

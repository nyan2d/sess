package sess

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

type Sess struct {
	secret []byte
}

func New(secret []byte) *Sess {
	return &Sess{
		secret: secret,
	}
}

func (ss *Sess) SignSession(session *Session) {
	session.Sign = ss.Sing(session)
}

func (ss *Sess) Sing(session *Session) string {
	hm := hmac.New(sha256.New, ss.secret)
	hm.Write([]byte(session.RawString()))
	sum := hm.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}

func (ss *Sess) IsSessionValid(session *Session) bool {
	if session == nil {
		return false
	}

	if time.Now().Unix() > session.Valid {
		return false
	}

	return session.Sign == ss.Sing(session)
}

func (ss *Sess) CreateSession(userID int64, expiresIn time.Duration) *Session {
	s := &Session{
		UserID:  userID,
		GUID:    GenerateUUID(),
		Created: time.Now().Unix(),
		Valid:   time.Now().Add(expiresIn).Unix(),
	}
	ss.SignSession(s)
	return s
}

func (ss *Sess) UpdateSession(session *Session, expiresIn time.Duration) {
	session.Valid = time.Now().Add(expiresIn).Unix()
	ss.SignSession(session)
}

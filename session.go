package sess

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type Session struct {
	UserID  int64  `json:"user_id"`
	GUID    string `json:"guid"`
	Created int64  `json:"created"`
	Valid   int64  `json:"valid"`
	Sign    string `json:"sign"`
}

func (s *Session) RawString() string {
	return strconv.FormatInt(s.UserID, 10) + ";" + s.GUID + ";" +
		strconv.FormatInt(s.Created, 10) + ";" + strconv.FormatInt(s.Valid, 10)
}

func (s *Session) Base64() string {
	j, _ := json.Marshal(s)
	return base64.StdEncoding.EncodeToString(j)
}

func ParseBase64(b64 string) (*Session, error) {
	buf, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, ErrParsingBase64
	}

	var jobj Session
	err = json.Unmarshal(buf, &jobj)
	if err != nil {
		return nil, ErrParsingJson
	}

	return &jobj, nil
}

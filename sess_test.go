package sess

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"testing"
	"time"
)

func TestSignSession(t *testing.T) {
	secret := []byte{0x01, 0x02, 0x03, 0x04}

	s := New(secret)
	session := s.CreateSession(0, 1*time.Second)
	s.SignSession(session)

	hm := hmac.New(sha256.New, secret)
	hm.Write([]byte(session.RawString()))
	sum := hm.Sum(nil)
	str := base64.StdEncoding.EncodeToString(sum)

	if session.Sign == "" {
		t.Error("sign is empty")
	}
	if session.Sign != str {
		t.Error("wrong sign")
	}
}

func TestSign(t *testing.T) {
	secret := []byte{0x01, 0x02, 0x03, 0x04}

	s := New(secret)
	session := s.CreateSession(0, 1*time.Second)
	sign1 := s.Sing(session)

	hm := hmac.New(sha256.New, secret)
	hm.Write([]byte(session.RawString()))
	sum := hm.Sum(nil)
	sign2 := base64.StdEncoding.EncodeToString(sum)

	if sign1 == "" {
		t.Error("sign is empty")
	}
	if sign1 != sign2 {
		t.Error("wrong sign")
	}
}

func TestIsSessionValid(t *testing.T) {
	secret := []byte{0x01, 0x02, 0x03, 0x04}
	s := New(secret)

	validSession := s.CreateSession(0, time.Second)
	s.SignSession(validSession)

	invalidTime := s.CreateSession(0, -1*time.Second)
	s.SignSession(invalidTime)

	invalidSign := s.CreateSession(1, time.Second)
	s.SignSession(invalidSign)
	invalidSign.UserID = 0

	if !s.IsSessionValid(validSession) {
		t.Fail()
	}
	if s.IsSessionValid(invalidTime) {
		t.Fail()
	}
	if s.IsSessionValid(invalidSign) {
		t.Fail()
	}
}

func TestUpdateSession(t *testing.T) {
	secret := []byte{0x01, 0x02, 0x03, 0x04}
	s := New(secret)

	outdatedSession := s.CreateSession(0, -1*time.Second)
	s.SignSession(outdatedSession)

	updatedSession := *outdatedSession
	s.UpdateSession(&updatedSession, time.Second)

	if s.IsSessionValid(outdatedSession) {
		t.Fail()
	}
	if !s.IsSessionValid(&updatedSession) {
		t.Fail()
	}
}

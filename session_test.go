package sess

import (
	"testing"
)

func TestSessionRawString(t *testing.T) {
	s := &Session{
		UserID:  0,
		GUID:    "12345678-1234-1234-1234-123456789012",
		Created: 12345678,
		Valid:   12345679,
		Sign:    "sign",
	}
	want := `0;12345678-1234-1234-1234-123456789012;12345678;12345679`
	have := s.RawString()
	if want != have {
		t.Fatal("wrong output")
	}
}

func TestSessionBase64(t *testing.T) {
	s := &Session{
		UserID:  0,
		GUID:    "12345678-1234-1234-1234-123456789012",
		Created: 12345678,
		Valid:   12345679,
		Sign:    "sign",
	}

	want := `eyJ1c2VyX2lkIjowLCJndWlkIjoiMTIzNDU2NzgtMTIzNC0xMjM0LTEyMzQtMTIzNDU2Nzg5MDEyIiwiY3JlYXRlZCI6MTIzNDU2NzgsInZhbGlkIjoxMjM0NTY3OSwic2lnbiI6InNpZ24ifQ==`
	have := s.Base64()

	if want != have {
		t.Fatal("wrong base64 output")
	}
}

func TestParseBase64(t *testing.T) {
	data := `eyJ1c2VyX2lkIjowLCJndWlkIjoiMTIzNDU2NzgtMTIzNC0xMjM0LTEyMzQtMTIzNDU2Nzg5MDEyIiwiY3JlYXRlZCI6MTIzNDU2NzgsInZhbGlkIjoxMjM0NTY3OSwic2lnbiI6InNpZ24ifQ==`
	parsed, err := ParseBase64(data)
	if err != nil {
		t.Fatal(err)
	}
	if parsed.UserID != 0 {
		t.Fatal("wrong userid")
	}
	if parsed.GUID != "12345678-1234-1234-1234-123456789012" {
		t.Fatal("wrong guid")
	}
	if parsed.Created != 12345678 {
		t.Fatal("wrong created")
	}
	if parsed.Valid != 12345679 {
		t.Fatal("wrong valid")
	}
}

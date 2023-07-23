package sess

import "testing"

func TestGenerateUUID(t *testing.T) {
	uuidA := GenerateUUID()
	uuidB := GenerateUUID()
	if len(uuidA) != 36 || len(uuidB) != 36 {
		t.Fatal("wrong uuid length")
	}
	if uuidA == uuidB {
		t.Fatal("uuids are not unique")
	}
}

package serverutil

import "testing"

func TestDecodebase64(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"dGVzdCBleGFtcGxl", "test example"},
		{"", ""},
	}

	for _, c := range tests {
		got := string(DecodeBase64(c.in))
		if got != c.out {
			t.Error("Expected content", c.out, "and got", got)
		}
	}
}

func TestEncodebase64(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"test example", "dGVzdCBleGFtcGxl"},
		{"", ""},
	}

	for _, c := range tests {
		got := EncodeBase64([]byte(c.in))
		if got != c.out {
			t.Error("Expected content", c.out, "and got", got)
		}
	}

}

func TestGetOriginalMessage(t *testing.T) {
	tests := []struct {
		inOut string
	}{
		{"text to cipher"},
		{""},
	}
	key := []byte("This is the test key of 32 bits.")

	for _, c := range tests {
		got := Decrypt(key, Encrypt(key, []byte(c.inOut)))
		if got != c.inOut {
			t.Error("Expected content", c.inOut, "and got", got)
		}
	}
}

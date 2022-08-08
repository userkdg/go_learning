package crypto

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	s := "sha1 test"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	// bs 16进制字符串 42aae338eae1f985c9a2188977e11a190002afe3
	fmt.Printf("%x\n", bs)
}

package encoding

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	data := "abc123!?$*&()'-=@~中国"
	eStr := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(eStr)
	dStr, err := base64.StdEncoding.DecodeString(eStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dStr))

	uEStr := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEStr)
	udStr, err := base64.URLEncoding.DecodeString(uEStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(udStr))
}

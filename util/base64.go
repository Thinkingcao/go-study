package main

import (
	"encoding/base64"
	"fmt"
)


func Base64Encode(rawStr string) string {
	return base64.StdEncoding.EncodeToString([]byte(rawStr))
}

func Base64Decode(endcodeStr string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(endcodeStr)
	if err == nil {
		return string(b), nil
	}
	return "", err
}

func main() {
	fmt.Println("sdw211111")
	a:=Base64Encode("123")
	b,_:=Base64Decode(a)
	fmt.Println(b)

}
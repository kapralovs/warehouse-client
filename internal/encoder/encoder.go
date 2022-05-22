package encoder

import "encoding/base64"

func Encode(username string,password string) string{
	credsString:=username+":"+password
	encodedCreds:=base64.StdEncoding.EncodeToString([]byte(credsString))
	return encodedCreds
}
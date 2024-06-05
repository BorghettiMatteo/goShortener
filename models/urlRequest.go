package models

import (
	"crypto"
	"encoding/hex"
)

type UrlRequestBody struct {
	ApiVersion string `json:"apiversion"`
	Plain      string `json:"plain"`
}

func (ur *UrlRequestBody) UrlEncoder() string {
	sha1 := crypto.SHA1.New()
	sha1.Write([]byte(ur.Plain))
	dump := sha1.Sum(nil)
	encodedUrl := hex.EncodeToString(dump)
	return string(encodedUrl)
}

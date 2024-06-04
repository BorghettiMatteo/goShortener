package models

import (
	"crypto"
	"encoding/hex"
)

type UrlRequest struct {
	Shortened string `json:"shortened"`
	Plain     string `json:"plain"`
}

func (ur *UrlRequest) UrlEncoder(apiVersion string) {
	sha1 := crypto.SHA1.New()
	sha1.Write([]byte(ur.Plain))
	dump := sha1.Sum(nil)
	encodedUrl := hex.EncodeToString(dump)
	ur.Shortened = "0.0.0.0:8080" + apiVersion + "url/" + string(encodedUrl)
}

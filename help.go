package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return string(bs)
}

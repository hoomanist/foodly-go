package main

import (
  "net/http"
  "log"
  "net/http/httptest"
  "crypto/md5"
	"crypto/sha1"
	"encoding/hex"
  "encoding/json"
  "golang.org/x/crypto/bcrypt"
  "bytes"
)

// generate a token based on password with a random salt
func GenerateToken(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

// create a hash
func Hash(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}

func PerformPostRequest(r http.Handler, path string, data map[string]string) *httptest.ResponseRecorder{
  body, err := json.Marshal(data)
  if err != nil {
    log.Fatalln(err)
  }
  b := bytes.NewBuffer(body)
  req, _ := http.NewRequest("POST", path, b)
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)
  return w
}

func PerformEmptyRequest(r http.Handler, method, path string)* httptest.ResponseRecorder{
  req, _ := http.NewRequest(method, path, nil)
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)
  return w
}

package client

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"net/http"
	"time"
)

const cookiePKCE = "isPKCE"

var (
	pkceCodeVerifier string

	pkceCodeChallenge string
)

const codeVerifierLenMin = 43
const codeVerifierLenMax = 128
const codeVerifierAllowedLetters = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ._~"

func generateCodeVerifier(n int) string {
	if n < codeVerifierLenMin {
		n = codeVerifierLenMin
	}
	if n > codeVerifierLenMax {
		n = codeVerifierLenMax
	}

	b := make([]byte, n)
	for i := range b {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(len(codeVerifierAllowedLetters))))
		b[i] = codeVerifierAllowedLetters[j.Int64()]
	}

	return string(b)
}

func generateCodeChallenge(codeVerifier string) string {
	s256 := sha256.New()
	s256.Write([]byte(codeVerifier))

	return base64.RawURLEncoding.EncodeToString(s256.Sum(nil))
}

func isPKCE(r *http.Request) bool {
	cookie, err := r.Cookie(cookiePKCE)
	if err != nil {
		return false
	}

	return cookie.Value == "true"
}

func resetPKCE(w http.ResponseWriter) (codeVerifier string) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookiePKCE,
		Path:    "/",
		Expires: time.Unix(0, 0),
	})

	codeVerifier = pkceCodeVerifier
	pkceCodeVerifier = ""

	return codeVerifier
}

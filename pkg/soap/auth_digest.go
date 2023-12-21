package soap

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// createAuthenticationDigestResponse creates the auth header for digest authorization.
// It's not a feature complete implementation see https://en.wikipedia.org/wiki/Digest_access_authentication
// AuthHeader example:
// WWW-Authenticate: Digest realm="testrealm@host.com",
//
//	 algorithm=MD5,
//		qop="auth,auth-int",
//		nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
//		opaque="5ccc069c403ebaf9f0171e9517f40e41"
func createAuthenticationDigestResponse(username string, password string, reqAuthHeader string, httpMethod string, fullUrl string) string {
	if !strings.HasPrefix(reqAuthHeader, "Digest") {
		log.Fatal("Can handle digest auth only!")
	}
	digestChallenge := parseAuthParam(strings.TrimPrefix(reqAuthHeader, "Digest "))
	realm := digestChallenge["realm"]
	nonce := digestChallenge["nonce"]
	algorithm := digestChallenge["algorithm"]
	qop := digestChallenge["qop"]

	if strings.ToLower(algorithm) != "md5" {
		panic(errors.New("unsupported hash algorithm=" + algorithm))
	}
	nonceCount := 00000001
	cnonce := createEightRandomBytes()
	response := "undefined"

	if strings.ToLower(qop) == "auth" {
		ha1 := md5Sum(username + ":" + realm + ":" + password)
		ha2 := md5Sum(httpMethod + ":" + fullUrl)
		response = md5Sum(fmt.Sprintf("%s:%s:%v:%s:%s:%s", ha1, nonce, nonceCount, cnonce, qop, ha2))
	} else {
		panic(errors.New("unsupported quality of protection (qop):" + algorithm))
	}

	return fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", Uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`,
		username, realm, nonce, fullUrl, cnonce, nonceCount, qop, response)
}

func md5Sum(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func createEightRandomBytes() string {
	b := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		log.Fatal("Can't create random bytes: \n" + err.Error())
	}
	return fmt.Sprintf("%x", b)[:16]
}

func parseAuthParam(allParams string) map[string]string {
	params := strings.Split(allParams, ",")
	result := map[string]string{}
	for _, param := range params {
		paramPair := strings.Split(param, "=")
		result[paramPair[0]] = strings.Trim(paramPair[1], "\"")
	}
	return result
}

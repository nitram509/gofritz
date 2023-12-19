package tr064

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strings"
)

type digestAuthParameter struct {
	username  string
	password  string
	realm     string
	nonce     string
	uri       string
	algorithm string
	qop       string
	cnonce    string
	response  string
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

// createAuthenticationDigest creates the auth header for digest authorization.
// It's not a feature complete implementation see https://en.wikipedia.org/wiki/Digest_access_authentication
// authHeader example:
// WWW-Authenticate: Digest realm="testrealm@host.com",
//
//	qop="auth,auth-int",
//	nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
//	opaque="5ccc069c403ebaf9f0171e9517f40e41"
func createAuthenticationDigest(username string, password string, authHeader string, method string, url string) string {
	if !strings.HasPrefix(authHeader, "Digest") {
		log.Fatal("Can handle digest auth only!")
	}
	authParams := parseAuthParam(strings.TrimPrefix(authHeader, "Digest "))
	d := digestAuthParameter{
		username:  username,
		password:  password,
		realm:     authParams["realm"],
		nonce:     authParams["nonce"],
		uri:       url,
		algorithm: authParams["algorithm"],
		qop:       authParams["qop"],
		cnonce:    createEightRandomBytes(),
		response:  "",
	}
	ha1 := md5Sum(d.username + ":" + d.realm + ":" + d.password)
	ha2 := md5Sum(method + ":" + url)
	nonceCount := 00000001
	d.response = md5Sum(fmt.Sprintf("%s:%s:%v:%s:%s:%s", ha1, d.nonce, nonceCount, d.cnonce, d.qop, ha2))
	return fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", Uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`,
		d.username, d.realm, d.nonce, d.uri, d.cnonce, nonceCount, d.qop, d.response)
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

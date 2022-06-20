package key

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Token struct {
	AuthTime int64                  `json:"auth_time"`
	Issuer   string                 `json:"iss"`
	Audience string                 `json:"aud"`
	Expires  int64                  `json:"exp"`
	IssuedAt int64                  `json:"iat"`
	Subject  string                 `json:"sub,omitempty"`
	UID      string                 `json:"uid,omitempty"`
	Claims   map[string]interface{} `json:"-"`
}

type KeyGenService interface {
	Encrypt(token Token) (string, error)
	Decrypt(cipherText string) (Token, error)
	IsCustomToken(cipherTExt string) bool
}

type keyGenService struct {
}

func NewKeyGenService() KeyGenService {
	return &keyGenService{}
}

const TOKEN_PREFIX string = "ONQLAVE-"

func (s *keyGenService) IsCustomToken(cipherTExt string) bool {
	return strings.HasPrefix(cipherTExt, TOKEN_PREFIX)
}

func (s *keyGenService) Encrypt(token Token) (string, error) {
	key := []byte(os.Getenv("TOKEN_GEN_KEY"))
	plaintext, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// convert to base64
	return TOKEN_PREFIX + base64.URLEncoding.EncodeToString(ciphertext), nil
}

func (s *keyGenService) Decrypt(cipherText string) (Token, error) {
	key := []byte(os.Getenv("TOKEN_GEN_KEY"))
	cipherText = strings.Replace(cipherText, TOKEN_PREFIX, "", -1)
	ciphertext, _ := base64.URLEncoding.DecodeString(cipherText)
	var token Token
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return token, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return token, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	err = json.Unmarshal(ciphertext, &token)
	if err != nil {
		return token, err
	}
	t := time.Now().UTC().Unix()
	if token.Expires <= t {
		return token, errors.New("token is already expired")
	}
	return token, err
}

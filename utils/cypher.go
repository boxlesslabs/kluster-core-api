//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================
 
/**
 **
 * @struct GeneralUtil
 * @EncodeBase64() return encoded base64 string
 * @DecodeBase64() return decoded base64 string
 * @Encrypt() return encrypted string
 * @Decrypt() return decrypted string
 **
**/

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"fmt"
)

func (util *GeneralUtil) EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (util *GeneralUtil) DecodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil { panic(err) }
	return data
}

func (util *GeneralUtil) Encrypt(key []byte, text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	cypher := make([]byte, aes.BlockSize+len(plaintext))
	iv := cypher[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cypher[aes.BlockSize:], plaintext)
	return base64.URLEncoding.EncodeToString(cypher)
}

func (util *GeneralUtil) Decrypt(key []byte, cryptoText string) string {
	cypher, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(cypher) < aes.BlockSize {
		panic("cypher too short")
	}
	iv := cypher[:aes.BlockSize]
	cypher = cypher[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cypher, cypher)

	return fmt.Sprintf("%s", cypher)
}
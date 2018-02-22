package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"io"
)

const saltSize = 32
const aesKeySize = 32
const argonTimes = 3
const argonMemory = 32 * 1024
const argonThreads = 4

func getRandomBytes(number int) ([]byte, error) {
	randomBytes := make([]byte, number)
	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func passwordToKey(password []byte, salt []byte) []byte {
	return argon2.Key(password, salt, argonTimes, argonMemory, argonThreads, aesKeySize)
}

func EncryptAES256withGCM(data []byte, password []byte) (encryptedData []byte, salt []byte, nonce []byte, err error) {

	if salt, err = getRandomBytes(saltSize); err != nil {
		return
	}

	key := passwordToKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	if nonce, err = getRandomBytes(aesgcm.NonceSize()); err != nil {
		return
	}

	encryptedData = aesgcm.Seal(nil, nonce, data, nil)
	return
}

func DecryptAES256withGCM(encryptedData []byte, salt []byte, nonce []byte, password []byte) (data []byte, err error) {
	key := passwordToKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	data, err = aesgcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return
	}

	return
}

func HexEncryptAES256(data string, password string) (encodedEncryptedData string, encodedSalt string, encodedNonce string, err error) {
	encryptedData, salt, nonce, err := EncryptAES256withGCM([]byte(data), []byte(password))
	if err != nil {
		return
	}
	encodedEncryptedData = hex.EncodeToString(encryptedData)
	encodedSalt = hex.EncodeToString(salt)
	encodedNonce = hex.EncodeToString(nonce)
	return

}

func HexDecryptAES256(encodedEncryptedData string, encodedSalt string, encodedNonce string, password string) (data []byte, err error) {
	decodedEncryptedData, err := hex.DecodeString(encodedEncryptedData)
	if err != nil {
		return
	}
	decodedSalt, err := hex.DecodeString(encodedSalt)
	if err != nil {
		return
	}
	decodedNonce, err := hex.DecodeString(encodedNonce)
	if err != nil {
		return
	}

	data, err = DecryptAES256withGCM(decodedEncryptedData, decodedSalt, decodedNonce, []byte(password))
	return
}

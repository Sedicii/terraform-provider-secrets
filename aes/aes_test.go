package aes

import (
	"fmt"
	"testing"
)

func TestEncryptDecryptCycle(t *testing.T) {
	data := "test_data"
	password := "test_password"

	encryptedData, salt, nonce, err := EncryptAES256withGCM([]byte(password), []byte(data))
	if err != nil {
		t.Fail()
	}

	decryptedData, err := DecryptAES256withGCM(encryptedData, salt, nonce, []byte(password))
	if err != nil {
		t.Fail()
	}
	if string(decryptedData) != data {
		t.Fail()
	}
}

func TestHexEncryptDecryptCycle(t *testing.T) {
	data := "test_data"
	password := "test_password"

	encryptedData, salt, nonce, err := HexEncryptAES256(data, password)
	if err != nil {
		t.Fail()
	}

	decryptedData, err := HexDecryptAES256(encryptedData, salt, nonce, password)
	if err != nil {
		t.Fail()
	}
	if string(decryptedData) != data {
		t.Fail()
	}
	fmt.Println("original data : ", data)
	fmt.Println("decrypted data : ", string(decryptedData))
}

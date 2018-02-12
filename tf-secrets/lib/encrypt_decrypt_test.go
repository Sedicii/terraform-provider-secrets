package lib

import "testing"

func TestEncryptDecryptCycle(t *testing.T) {
	password := "test_password"

	vars := map[string]string{
		"data1": "test_data1",
		"data2": "test_data2",
		"data3": "test_data3",
	}
	encryptedVars, err := EncryptVars(&vars, password)

	if err != nil {
		t.Fail()
	}

	decryptedVars, err := DecryptVars(encryptedVars, password)

	if err != nil {
		t.Fail()
	}

	for k, v := range vars {
		if v != decryptedVars[k] {
			t.Fail()
		}
	}
}

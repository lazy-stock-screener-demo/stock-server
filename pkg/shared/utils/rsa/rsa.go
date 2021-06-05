package rsahelper

import (
	"io/ioutil"
)

func GetPrivatePemCert() ([]byte, error) {
	keyData, _ := ioutil.ReadFile("login.rsa")
	return keyData, nil
}

func GetPublicPemCert() ([]byte, error) {
	keyData, _ := ioutil.ReadFile("login.rsa.pub")
	// if keyData == "" {
	// 	err := errors.New("unable to find appropriate key")
	// 	return keyData, err
	// }
	return keyData, nil
}

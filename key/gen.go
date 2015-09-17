package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func Gen(outFile string, length int) {
	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, length)
	if err != nil {
		log.Println("ERROR: key.Gen generating private:", err)
		return
	}

	// Note : Only used for 3rd and subsequent primes
	//log.Printf("CRTValues : Exp[%s]\n Coeff[%s]\n R[%s]\n", CRTVal[2].Exp.String(), CRTVal[2].Coeff.String(), CRTVal[2].R.String())

	// Note : if you want to have multi primes,
	//        use rsa.GenerateMultiPrimeKey() function instead of
	//        rsa.GenerateKey() function
	//        see http://golang.org/pkg/crypto/rsa/#GenerateMultiPrimeKey

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Precompute
	privatekey.Precompute()

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Validate
	err = privatekey.Validate()
	if err != nil {
		log.Println("ERROR: key.Gen validating private:", err)
		return
	}

	privateFile, err := os.Create(outFile)
	if err != nil {
		log.Println("ERROR: key.Gen creating private file:", err)
		return
	}
	defer privateFile.Close()
	err = pem.Encode(
		privateFile,
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
		},
	)
	if err != nil {
		log.Println("ERROR: key.Gen writing private:", err)
		return
	}

	var publickey *rsa.PublicKey
	publickey = &privatekey.PublicKey
	publicMarshal, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		log.Println("ERROR: key.Gen marshaling public key:", err)
		return
	}
	publicFile, err := os.Create(outFile + ".pub")
	if err != nil {
		log.Println("ERROR: key.Gen creating public file:", err)
		return
	}
	defer publicFile.Close()
	err = pem.Encode(
		publicFile,
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicMarshal,
		},
	)
	if err != nil {
		log.Println("ERROR: key.Gen writing public:", err)
		return
	}
}

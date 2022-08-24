package adapters

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type EncoderAndDecoder struct {
	Key   []byte
	block cipher.Block
}

func NewEncoderAndDecode(key string) *EncoderAndDecoder {
	byteKey := []byte(key)
	//Create a new AES cipher using the key
	block, err := aes.NewCipher(byteKey)

	//IF NewCipher failed, exit:
	if err != nil {
		panic(err)
	}

	return &EncoderAndDecoder{
		Key:   byteKey,
		block: block,
	}
}

func (e EncoderAndDecoder) Encrypt(message string) (encoded string, err error) {
	plainText := []byte(message)
	cipherText, err := e.encryptRawMessageWithAESStrategy(plainText)
	if err != nil {
		panic(err)
	}
	return base64.RawStdEncoding.EncodeToString(cipherText), nil

}
func (e EncoderAndDecoder) encryptRawMessageWithAESStrategy(plainText []byte) (encoded []byte, err error) {
	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return encoded, errors.New("invalid reader " + err.Error())
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(e.block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return cipherText, err
}

func (e EncoderAndDecoder) Decrypt(message string) (decoded string, err error) {
	cipherText, err := base64.RawStdEncoding.DecodeString(message)
	if err != nil {
		return
	}
	return e.decryptWithAESStrategy(cipherText)
}

func (e EncoderAndDecoder) decryptWithAESStrategy(cipherText []byte) (string, error) {
	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext block size is too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(e.block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}

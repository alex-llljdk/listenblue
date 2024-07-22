package encrypt

import (
	"crypto/md5"
	"encoding/base32"
	"encoding/hex"
	"fmt"

	"github.com/zeromicro/go-zero/core/codec"
	"golang.org/x/crypto/scrypt"
)

const (
	passwordEncryptSeed = "(zjhx)@#$"
	roomNumberAesKey    = "5A2E746B08D846502F37A6E2D85D583B"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func RoomNumberEnctypy(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func EncRoomNumber(RoomNumber string) (string, error) {
	data, err := codec.EcbEncrypt([]byte(roomNumberAesKey), []byte(RoomNumber))
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.WithPadding(-1).EncodeToString(data), nil
}

func DecRoomNumber(RoomNumber string) (string, error) {
	originalData, err := base32.StdEncoding.WithPadding(-1).DecodeString(RoomNumber)
	if err != nil {
		return "", err
	}
	data, err := codec.EcbDecrypt([]byte(roomNumberAesKey), originalData)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Md5Sum(data []byte) string {
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}

func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp[16:]
}

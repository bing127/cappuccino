package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Hmac(key string, data string) string {
	hmacHash := hmac.New(md5.New, []byte(key))
	hmacHash.Write([]byte(data))
	return hex.EncodeToString(hmacHash.Sum([]byte("")))
}

func IsStringEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}

func GetUUID() string {
	u := uuid.NewV4()
	return strings.ReplaceAll(u.String(),"-","")
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func Base64ToImage(imageBase64 string) ([]byte, error) {
	image, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func GetDirFiles(dir string) ([]string, error) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesRet := make([]string, 0)

	for _, file := range dirList {
		if file.IsDir() {
			files, err := GetDirFiles(dir + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}

			filesRet = append(filesRet, files...)
		} else {
			filesRet = append(filesRet, dir+string(os.PathSeparator)+file.Name())
		}
	}

	return filesRet, nil
}

func GetCurrentTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}
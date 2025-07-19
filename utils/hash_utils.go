package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)


func GetHash(input string, hashType string) string {

	if hashType == "md5" {
		md5Hash := md5.New()
		io.WriteString(md5Hash, input)

		hashBytes := md5Hash.Sum(nil)

		retString := fmt.Sprintf("%x", hashBytes)

		return retString

	} else {
		sha256Hash := sha256.New()
		io.WriteString(sha256Hash, input)

		hashBytes := sha256Hash.Sum(nil)

		retString := fmt.Sprintf("%x", hashBytes)

		return retString
	}
}


func ReadFileBytes(filename string) string {

	file,err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening the file")
	}
	defer file.Close() // Ensures the file is closed when the function exits

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(data))
	return string(data)
}

func StrMatch(str1 string, str2 string) bool {


		if str1 == str2 {
			return true
		} else {
			return false
		}
}
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getHash(input string, hashType string) string {

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


func readFileBytes(filename string) string {

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

func strMatch(str1 string, str2 string) bool {


		if str1 == str2 {
			return true
		} else {
			return false
		}
}

func containsHelpFlag(args []string) bool {
    for _, arg := range args {
        if arg == "-h" || arg == "--help" {
            return true
        }
    }
    return false
}

func main() {
	
	// Arg flags
	var inputStr string
	var hashType string
	var cmpValue string

	// Specify args and their usage
	flag.StringVar(&inputStr, "filename", "test.exe", "The file you want to check the hash of.")
	flag.StringVar(&hashType, "hashType", "md5", "ex: md5 or sha256")
	flag.StringVar(&cmpValue, "cmpHash", "NONE", "The hash you are comparing to.")



	if len(os.Args) == 1 || containsHelpFlag(os.Args) {
        fmt.Println("Usage:")
        flag.PrintDefaults()
        return
    }

	flag.Parse()


	fmt.Println("Input file:", inputStr)
	fmt.Println("-----------------------------------------------------------------------------------------------------------")

	hashType = strings.ToLower(hashType)

	if hashType == "md5" {
		fmt.Println("Hash Type: MD5 ", hashType)
		fmt.Println("-----------------------------------------------------------------------------------------------------------")

	}
	if hashType == "sha256" {
		fmt.Println("Hash Type: SHA256 ", hashType)
		fmt.Println("-----------------------------------------------------------------------------------------------------------")

	}

	// 1 md5
	// 2 sha256

	fileToHash := readFileBytes(inputStr)

	hashStr := getHash(fileToHash, hashType)

	fmt.Println("Input file hash: " + hashType + ": " + hashStr)


	fmt.Println("-----------------------------------------------------------------------------------------------------------")

	if strMatch(hashStr,cmpValue) {
		fmt.Println("Hashes \033[32mMATCH\033[0m")
	} else {
		fmt.Println("Hashes \033[31mDONT MATCH\033[0m ")
	}

	}	
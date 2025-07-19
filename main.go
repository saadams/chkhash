package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/saadams/chkhash/utils"
)



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

	fileToHash := utils.ReadFileBytes(inputStr)

	hashStr := utils.GetHash(fileToHash, hashType)


	fmt.Println("Desired hash: " + hashType + ": " + cmpValue)


	fmt.Println("-----------------------------------------------------------------------------------------------------------")

	fmt.Println("Input file hash: " + hashType + ": " + hashStr)


	fmt.Println("-----------------------------------------------------------------------------------------------------------")

	if utils.StrMatch(hashStr,cmpValue) {
		fmt.Println("Hashes \033[32mMATCH\033[0m")
	} else {
		fmt.Println("Hashes \033[31mDONT MATCH\033[0m ")
	}

	}	
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 0 for encryption or 1 for decryption:\n")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	fmt.Print("Enter file path: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	fmt.Print("Enter key: ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Print("Enter output path: ")
	outputPath, _ := reader.ReadString('\n')
	outputPath = strings.TrimSpace(outputPath)

	if choice == "0" {
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		encryptedContent := encrypt(fileContent, key)
		err = ioutil.WriteFile(outputPath, encryptedContent, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("File encrypted successfully!")
	} else if choice == "1" {
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		decryptedContent, err := decrypt(fileContent, key)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ioutil.WriteFile(outputPath, decryptedContent, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("File decrypted successfully!")
	} else {
		fmt.Println("Invalid choice!")
	}
}

func encrypt(data []byte, key string) []byte {
	// encryption logic
	return data
}

func decrypt(data []byte, key string) ([]byte, error) {
	// decryption logic
	return data, nil
}

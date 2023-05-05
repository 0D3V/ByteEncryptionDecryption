package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func XOREncrypt(data []byte, key string) []byte {
    CL := []byte(key)
    FP := 112
    BA := make([]byte, len(data)+1)
    KA := 0
    for MP := 0; MP < len(data); MP++ {
        BA[MP] += byte((int(data[MP]) ^ int(CL[KA])) ^ FP)
        if KA == len(key)-1 {
            KA = 0
        } else {
            KA = KA + 1
        }
    }
    BA[len(data)] = byte(112 ^ FP)
    return BA
}

func XORDecrypt(data []byte, key string) []byte {
    X1 := []byte(key)
    A1 := int(data[len(data)-1]) ^ 112
    BN := make([]byte, len(data))
    CC := 0
    for EE := 0; EE < len(data)-1; EE++ {
        BN[EE] = byte((int(data[EE]) ^ A1) ^ int(X1[CC]))
        if CC == len(key)-1 {
            CC = 0
        } else {
            CC = CC + 1
        }
    }
    return BN
}

func main() {
    fmt.Println("Enter 0 for encryption or 1 for decryption:")
    scanner := bufio.NewScanner(os.Stdin)
    var choice int
    for scanner.Scan() {
        input := scanner.Text()
        if input == "0" || input == "1" {
            choice = int(input[0] - '0')
            break
        }
        fmt.Println("Invalid choice! Please enter 0 for encryption or 1 for decryption:")
    }
    fmt.Println("Enter the file path:")
    scanner.Scan()
    filePath := scanner.Text()
    fmt.Println("Enter the key:")
    scanner.Scan()
    key := scanner.Text()
    fmt.Println("Enter the output file name:")
    scanner.Scan()
    outFilePath := scanner.Text()

    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }

    var outputData []byte
    if choice == 0 {
        outputData = XOREncrypt(data, key)
    } else if choice == 1 {
        outputData = XORDecrypt(data, key)
    } else {
        fmt.Println("Invalid choice!")
        return
    }

    err = ioutil.WriteFile(outFilePath, outputData, 0644)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Done!")
}

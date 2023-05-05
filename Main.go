package main

import (
    "crypto/rand"
    "fmt"
    "io/ioutil"
)

func XOREncrypt(up []byte, BB2 []byte) []byte {
    FP := make([]byte, 1)
    rand.Read(FP)
    BA := make([]byte, len(up)+1)
    KA := 0
    for MP := 0; MP < len(up); MP++ {
        BA[MP] = (up[MP] ^ BB2[KA]) ^ FP[0]
        if KA == len(BB2)-1 {
            KA = 0
        } else {
            KA++
        }
    }
    BA[len(up)] = 112 ^ FP[0]
    return BA
}

func XOR_Decrypt(P1 []byte, K1 []byte) []byte {
    A1 := P1[len(P1)-1] ^ 112
    BN := make([]byte, len(P1))
    CC := 0
    for EE := 0; EE < len(P1); EE++ {
        BN[EE] = (P1[EE] ^ A1) ^ K1[CC]
        if CC == len(K1)-1 {
            CC = 0
        } else {
            CC++
        }
    }
    BN = BN[:len(P1)-1]
    return BN
}

func main() {
    var operation string
    fmt.Println("Enter 'encrypt' or 'decrypt'")
    fmt.Scanln(&operation)
    var inputFilePath, outputFilePath, key string
    fmt.Println("Enter input file path:")
    fmt.Scanln(&inputFilePath)
    fmt.Println("Enter output file path:")
    fmt.Scanln(&outputFilePath)
    fmt.Println("Enter key:")
    fmt.Scanln(&key)
    inputBytes, err := ioutil.ReadFile(inputFilePath)
    if err != nil {
        panic(err)
    }
    var outputBytes []byte
    if operation == "encrypt" {
        outputBytes = XOREncrypt(inputBytes, []byte(key))
    } else if operation == "decrypt" {
        outputBytes = XOR_Decrypt(inputBytes, []byte(key))
    } else {
        fmt.Println("Invalid operation")
        return
    }
    err = ioutil.WriteFile(outputFilePath, outputBytes, 0644)
    if err != nil {
        panic(err)
    }
}

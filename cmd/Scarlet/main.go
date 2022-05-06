package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)






func main() {

	provider := "https://mainnet.infura.io"

    client, err := ethclient.Dial(provider)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Provider Connection Made Successfully")
    _ = client // we'll use this in the upcoming sections 

    fmt.Println("Successfully Ran")
    log.Println("Success")
}
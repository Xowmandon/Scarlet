package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)


func LogFatalError(err error ) {
    log.Fatalf("An Error Occured. Err: %s", err)
}



func main() {

    //Load Client Link/Auth
    err := godotenv.Load("./client/client.env")
    if err != nil {LogFatalError(err)}

    authLink := os.Getenv("PROVIDER_AUTH")

    client, err := ethclient.Dial(authLink)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Provider Connection Made Successfully")
    _ = client // we'll use this in the upcoming sections 

    fmt.Println("Successfully Ran")

}
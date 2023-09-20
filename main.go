// main.go
package main

func main() {
  println("Ba dum, tss!")
}

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("SARIF test")

    secretKey := os.Getenv("SECRET_KEY")
    fmt.Println("Secret Key:", secretKey)

    password := "supersecret"
    fmt.Println("Password:", password)

    userInput := "userInputValue"
    query := "SELECT * FROM users WHERE username='" + userInput + "'"
    fmt.Println("SQL Query:", query)

    // encryptedData := weakEncrypt("sensitive data")
    // fmt.Println("Encrypted Data:", encryptedData)

    username := "admin"
    pwd := "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"

    fmt.Println("Doing something with: ", username, pwd)
}

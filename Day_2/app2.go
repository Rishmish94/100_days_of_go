package main

import (
    "fmt"
    "strconv"
)

func main() {
    name := "Rishabh"
    age := 30
    message := "My name is " + name + " and I am " + strconv.Itoa(age) + " years old."
    fmt.Println(message)
}

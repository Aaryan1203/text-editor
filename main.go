package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: goeditor <filename>")
        return
    }

    filename := os.Args[1]
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error opening file %s: %v\n", filename, err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) 
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file %s: %v\n", filename, err)
    }
}

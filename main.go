package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
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

	reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        cmd, _ := reader.ReadString('\n')
        cmd = strings.ToLower(strings.TrimSpace(cmd))

        switch cmd {
        case "exit":
            fmt.Println("Exiting...")
            return
        case "save":
            fmt.Println("Saving file... (not implemented)")
        default:
            fmt.Println("Unknown command:", cmd)
        }
    }

}

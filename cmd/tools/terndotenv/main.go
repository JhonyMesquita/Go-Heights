package main

import (
    "fmt"
    "os/exec"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        panic(err)
    }

    cmd := exec.Command("tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf")
    output, err := cmd.CombinedOutput() // Captura tanto stdout quanto stderr
    if err != nil {
        fmt.Printf("Command output: %s\n", string(output)) // Exibe a saída completa
        panic(err)
    }
}

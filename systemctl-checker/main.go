package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
    cmd := exec.Command("bash", "-c", "systemctl | cat")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        fmt.Printf("wow, got error %s", err)
    }

    outputLines := strings.Split(string(output), "\r\n")

    fmt.Println(outputLines)
}

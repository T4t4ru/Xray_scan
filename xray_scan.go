package main

import (
	"bufio"
	"os"
	"os/exec"
	"fmt"
	"syscall"
)

func main() {
	maxMemory := int64(3 * 512)

	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		cmd := exec.Command("./xray_linux_amd64", "ws", "--browser", fmt.Sprintf("http://%s/", s), "--html-output", fmt.Sprintf("%s.html", s))
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to run command: %s\n", err)
		}
	}

	// Note: SetRLimit may require additional privileges
	var rlimit syscall.Rlimit
	rlimit.Max = uint64(maxMemory)
	rlimit.Cur = uint64(maxMemory)
	err = syscall.Setrlimit(syscall.RLIMIT_RSS, &rlimit)
	if err != nil {
		fmt.Printf("Error Setting Rlimit: %s", err)
		os.Exit(1)
	}
}

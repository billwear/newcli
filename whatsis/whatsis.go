package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getProcessList() (string, error) {
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
func filterProcesses(processList, filter string) string {
	var filteredLines []string
	lines := strings.Split(processList, "\n")
	for _, line := range lines {
		if strings.Contains(line, filter) {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n")
}
func pageOutput(output string) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	const pageSize = 20 // Number of lines per page
	lineCount := 0

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCount++
		if lineCount%pageSize == 0 {
			fmt.Print("Press 'Enter' for more...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}
func main() {
	processList, err := getProcessList()
	if err != nil {
		fmt.Println("Error fetching processes:", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		filter := os.Args[1]
		processList = filterProcesses(processList, filter)
	}

	pageOutput(processList)
}

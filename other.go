package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFullLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Notify() {
	fmt.Printf("\n======================================================\n[🤷🏻] No notes available, please make at least one note\n======================================================\n\n")
}

func NotifyCRUD(s string) {
	fmt.Printf("\n============================\n[✅] %s successfully\n============================\n\n", s)
}

func NotifyNotFound() {
	fmt.Printf("\n===================\n[🥀] Ups, Not Found\n===================\n\n")
}

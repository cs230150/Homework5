package main

import (
	"bufio"
	"fmt"
	"homework5/cmd/internal/commission"
	"os"
	"strconv"
	"strings"
)

const (
	minAmount = 500
	maxAmount = 15_000_000
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter amont: ")
	sumStr, _ := reader.ReadString('\n')
	sumStr = strings.TrimSpace(sumStr)
	amount, err := strconv.Atoi(sumStr)
	if err != nil {
		fmt.Println("Error: enter valid amount")
		return
	}
	if amount < minAmount {
		fmt.Printf("Error: minimum transfer amount %d sum\n", minAmount)
		return
	}
	if amount > maxAmount {
		fmt.Printf("Error: maximum transfer amount %d sum\n", maxAmount)
		return
	}

	fmt.Printf("Is that Alif card? (1-yes/0-no): ")
	cardStr, _ := reader.ReadString('\n')
	cardStr = strings.TrimSpace(cardStr)
	isAlifCard := cardStr == "1"

	comm := commission.Calculate(isAlifCard, amount)
	total := amount + comm

	fmt.Println("\n====== Receipt ===========")
	fmt.Printf("Service: XIKMATOVA NOZIMA\n")
	fmt.Printf("Summa: %d сум\n", amount)
	fmt.Printf("Commission: %d сум\n", comm)
	fmt.Printf("Total: %d сум\n", total)
	fmt.Println("Status: Done")
	fmt.Println("===========================")
}

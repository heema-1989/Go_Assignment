package main

import (
	"BankingApp/app/functions"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	users := []functions.Users{{Username: "Heema Goratela", CardNo: 6351545, PinNo: 1905, AccountType: "Savings", Balance: 40000},
		{Username: "Dhatri Goratela", CardNo: 9428319, PinNo: 2006, AccountType: "Fixed Deposit", Balance: 50000},
		{Username: "Khushi Rajpal", CardNo: 4567834, PinNo: 1876, AccountType: "Recurring", Balance: 30000},
		{Username: "Riya Sonara", CardNo: 3498744, PinNo: 2060, AccountType: "Savings", Balance: 40000}}
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("Welcome to our ATM.Please enter your credentials as asked.")
	fmt.Println("--------------------------------------------------------------------------------------")
	functions.ValidateUsers(&users)
MENU:
	fmt.Println("Please choose any one from the below mentioned options..")
	fmt.Println("--------------------------------------------")
	fmt.Println("1. Withdraw money from ATM.")
	fmt.Println("2. Deposit money to ATM.")
	fmt.Println("3. Check Account Balance.")
	fmt.Println("4. Exit.")
	inp, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(inp))
	if err != nil {
		fmt.Println("Please enter valid input. Enter digits only")

	}

	switch choice {
	case 1:
		functions.WithdrawMoney(&users)
		goto MENU
	case 2:
		functions.DepositMoney(&users)
		goto MENU
	case 3:
		functions.CheckBalance(&users)
		goto MENU
	case 4:
		fmt.Println("Thank you for using our service.Have a nice day!")

	}

}

package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func DepositMoney(u *[]Users) {

DEPOSIT:
	fmt.Println("Please enter amount to be deposited.")
	inp, _ := reader.ReadString('\n')
	depAmt, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		fmt.Println("Please enter valid input. Enter digits only")
		goto DEPOSIT
	}
	fmt.Println("Your current balance is: ", usr.Balance)
	usr.Balance += depAmt
	fmt.Printf("You have successfully deposited %.2f\n", depAmt)
	fmt.Printf("Your new balance is: %.2f\n", usr.Balance)

}

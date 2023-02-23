package functions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	//"errors"
)

var usr Users

func ValidateUsers(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			ValidateUsers(u)
		}
	}()
	fmt.Println("Please enter your card number")

	inp, _ := reader.ReadString('\n')

	card_no, err := strconv.Atoi(strings.TrimSpace(inp))

	if err != nil {
		panic("please enter valid input. Enter digits only")

	}

	for _, val := range *u {
		if card_no == val.CardNo {
			usr = val
			fmt.Println(val)
			break
		}
	}
	if (usr == Users{}) {
		panic("User  not found")
	}

	count := 0
PIN:
	for {

		fmt.Println("Please enter your pin number.")
		inp, _ = reader.ReadString('\n')
		pin_no, err1 := strconv.Atoi(strings.TrimSpace(inp))
		if err1 != nil {
			panic("Please enter valid input. Enter digits only")

		}
		if pin_no != usr.PinNo {

			if count >= 3 {
				log.Fatal("Maximum limit for entering pin exceeded.Please try again after some time.")
				return
			} else {
				fmt.Println("Please enter valid pin number...")
				fmt.Printf("\nOnly %d attempts left.Enter valid pin.", 3-count)
				count++
				goto PIN
			}

		} else {
			fmt.Println("--------------------------------------------")
			fmt.Println("Successfully logged in to your account..")
			fmt.Println("--------------------------------------------")
			fmt.Printf("Welcome %v\n", usr.Username)
			fmt.Println("--------------------------------------------")
			break
		}
	}
}

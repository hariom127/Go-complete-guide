package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.ReadFloatFromFile(balanceFile)
	fmt.Println("Help Desk 24/7 :", randomdata.PhoneNumber())
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Can't processed 👾")
	}
	for {
		getWelcomeMsg()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is : %.2f\n", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("Enter amount: ")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount, must be grater then 0")
				return
			}
			accountBalance = accountBalance + deposit
			fmt.Println("Deposit success, your balance is: ", accountBalance)
			fileops.WriteValueToFile(balanceFile, accountBalance)
		case 3:
			var withdraw float64
			fmt.Print("Enter amount: ")
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid amount, must be grater then 0")
				return
			}
			if withdraw > accountBalance {
				fmt.Println("Insufficient balance")
				return
			}
			accountBalance = accountBalance - withdraw
			fmt.Println("Withdraw success, your balance is: ", accountBalance)
			fileops.WriteValueToFile(balanceFile, accountBalance)

		default:
			fmt.Println("Goodbye 👋...")
			fmt.Println("Thanks for choosing us 💐...")
			return
		}

		//========================
		// if choice == 1 {
		// 	fmt.Printf("Your account balance is : %.2f\n", accountBalance)
		// 	continue
		// } else if choice == 2 {
		// 	var deposit float64
		// 	fmt.Print("Enter amount: ")
		// 	fmt.Scan(&deposit)
		// 	if deposit <= 0 {
		// 		fmt.Println("Invalid amount, must be grater then 0")
		// 		return
		// 	}
		// 	accountBalance = accountBalance + deposit
		// 	fmt.Println("Deposit success, your balance is: ", accountBalance)
		// 	continue
		// } else if choice == 3 {
		// 	var withdraw float64
		// 	fmt.Print("Enter amount: ")
		// 	fmt.Scan(&withdraw)

		// 	if withdraw <= 0 {
		// 		fmt.Println("Invalid amount, must be grater then 0")
		// 		return
		// 	}
		// 	if withdraw > accountBalance {
		// 		fmt.Println("Insufficient balance")
		// 		return
		// 	}

		// 	accountBalance = accountBalance - withdraw
		// 	fmt.Println("Withdraw success, your balance is: ", accountBalance)
		// 	continue
		// } else {
		// 	fmt.Println("Goodbye 👋...")
		// 	break
		// }
	}

}

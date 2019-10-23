package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards > 50 {
		fmt.Println("Looks like you've managed to     make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better disguise next time?")
		isHeistOn = false
	}

	if openedVault := rand.Intn(100); openedVault > 70 && isHeistOn {
		fmt.Println("Vault is open")
		fmt.Println("Grab and GO!")
	} else if openedVault < 70 && !isHeistOn {
		isHeistOn = false
		fmt.Println("Vault cannot be opened")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("failed")
		case 1:
			isHeistOn = false
			fmt.Println("failed1")
		case 2:
			isHeistOn = false
			fmt.Println("failed2")
		case 3:
			isHeistOn = false
			fmt.Println("failed3")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("You've gotten away with ", amtStolen)
	}

}

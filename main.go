package main

import (
	"fmt"
	"math/rand"
)

func main() {

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if openedVault >= 70 && isHeistOn == true {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Vault could not be opened.")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist failed. Start over!")
		case 1:
			isHeistOn = false
			fmt.Println("When did they start raising dogs in vaults??")
		case 2:
			isHeistOn = false
			fmt.Println("Looks like this fingerprint scanner won't accept any fingerprint...")
		case 3:
			isHeistOn = false
			fmt.Println("Did I even pack the burlap bags?")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(10000000)
		fmt.Println(amtStolen)
	}

	fmt.Println(isHeistOn)

}

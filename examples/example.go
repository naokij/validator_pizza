package main

import (
	"fmt"

	"github.com/naokij/validator_pizza"
)

func main() {
	resp, _ := validator_pizza.ValidateEmail("ditydi@cars2.club")
	fmt.Println("ditydi@cars2.club", resp)

	resp2, _ := validator_pizza.ValidateDomain("gamil.com")
	fmt.Println("gamil.com", resp2)
	fmt.Println("DidYouMean", *resp2.DidYouMean)
}

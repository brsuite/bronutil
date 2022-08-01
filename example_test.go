package bronutil_test

import (
	"fmt"
	"math"

	"github.com/brsuite/bronutil"
)

func ExampleAmount() {

	a := bronutil.Amount(0)
	fmt.Println("Zero Bronees:", a)

	a = bronutil.Amount(1e8)
	fmt.Println("100,000,000 Broneess:", a)

	a = bronutil.Amount(1e5)
	fmt.Println("100,000 Broneess:", a)
	// Output:
	// Zero Bronees: 0 BRON
	// 100,000,000 Broneess: 1 BRON
	// 100,000 Broneess: 0.001 BRON
}

func ExampleNewAmount() {
	amountOne, err := bronutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := bronutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := bronutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := bronutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 BRON
	// 0.01234567 BRON
	// 0 BRON
	// invalid brocoin amount
}

func ExampleAmount_unitConversions() {
	amount := bronutil.Amount(44433322211100)

	fmt.Println("Bronees to kBRON:", amount.Format(bronutil.AmountKiloBRON))
	fmt.Println("Bronees to BRON:", amount)
	fmt.Println("Bronees to MilliBRON:", amount.Format(bronutil.AmountMilliBRON))
	fmt.Println("Bronees to MicroBRON:", amount.Format(bronutil.AmountMicroBRON))
	fmt.Println("Bronees to Bronees:", amount.Format(bronutil.AmountBronees))

	// Output:
	// Bronees to kBRON: 444.333222111 kBRON
	// Bronees to BRON: 444333.222111 BRON
	// Bronees to MilliBRON: 444333222.111 mBRON
	// Bronees to MicroBRON: 444333222111 μBRON
	// Bronees to Bronees: 44433322211100 Bronees
}

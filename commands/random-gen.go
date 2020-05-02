package command

import (
	"fmt"
	"math/rand"
	"time"
	str "strconv"

	"github.com/Josh012422/gocharm/misc"
)

func RanNum (min int, max int) {

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn((max) - min)
	rS := str.Itoa(r)
	txt := misc.Bold(misc.Cyan("Gocharm says:"))
	rT := misc.Bold(rS)

	fmt.Printf("\n\n%s\n\n   %s \n\n\n", txt, rT)

	return
}

func RanDecision (dec1, dec2 string) {

	gocharmTxt := misc.Bold(misc.Cyan("Gocharm says: "))
	someTxt := misc.Bold(misc.Cyan("Gocharm speaked. Gocharm out."))

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn((2))

	switch r {
	   case 0:
		   decTxt := misc.Bold(misc.Yellow("Go for "))
		   dec := misc.Bold(misc.Green(dec1 + "."))
		   fmt.Printf("\n\n%s \n\n   %s%s\n\n%s \n\n\n", gocharmTxt, decTxt, dec, someTxt)
		   return
	   case 1:
		   decTxt := misc.Bold(misc.Yellow("Your destiny is "))
		   dec := misc.Bold(misc.Green(dec2 + "."))
		   fmt.Printf("\n\n%s \n\n   %s%s\n\n%s \n\n\n", gocharmTxt, decTxt, dec, someTxt)
		   return
	}
}

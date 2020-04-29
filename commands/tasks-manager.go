package command

import (
	"fmt"
	str "strconv"

	"github.com/spf13/viper"
	"github.com/Josh012422/gocharm/misc"
)

func Add (title string) {
	viper.New()
	success := misc.Green("Succesfully saved task with title: ")
	titleColor := misc.Cyan(title)
	numberColor := misc.Green("and number: ")

	sum := 1
	numberRaw := viper.GetInt("tasks.future_number")
	numberStr := str.Itoa(numberRaw)

	viper.Set("tasks.task." + numberStr, title)
	fmt.Printf("%s%s %s%d\n", success, titleColor, numberColor,  numberRaw)
	viper.Set("tasks.current_number", numberRaw)
	viper.WriteConfig()
	numberRaw += sum;
	numberStr = str.Itoa(numberRaw)
	viper.Set("tasks.future_number", numberStr)
	viper.WriteConfig()

	return
}

func List () {
	viper.New()
	numberRaw := viper.GetInt("tasks.current_number")
	numberStr := str.Itoa(numberRaw)
	numRaw := 1
	fmt.Sprintf(numberStr)

	for i := 0; i < numberRaw; i++{
		sum := 1
	//	numRaw := 1
	//	fmt.Println(numberRaw)
		numStr := str.Itoa(numRaw)
		numRaw += sum;
		numColor := misc.Green(numStr)

		title := viper.GetString("tasks.task." + numStr)
		fmt.Printf("%s. %s\n", numColor, title)
	//	fmt.Println(numRaw, numStr, numColor)
	}

	return

}

package command

import (
	"fmt"
	str "strconv"
	"errors"

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

	viper.Set("tasks.task." + numberStr, nil)
	viper.Set("tasks.task." + numberStr + "." + "title", title)
	viper.Set("tasks.task." + numberStr + "." + "status", "uncompleted")
	fmt.Printf("%s%s %s%d\n", success, titleColor, numberColor,  numberRaw)
	viper.Set("tasks.current_number", numberRaw)
	viper.WriteConfig()
	numberRaw += sum;
	viper.Set("tasks.future_number", numberRaw)
	viper.WriteConfig()

	return
}

func List () {
	viper.New()
	numberRaw := viper.GetInt("tasks.current_number")
	numberStr := str.Itoa(numberRaw)

	if numberRaw == 0 {
		fmt.Println(misc.Red("No pending tasks"))
	}
	statusNumRaw := 1

	numRaw := 1
	fmt.Sprintf(numberStr)

	for i := 0; i < numberRaw; i++{
		statusNum := str.Itoa(statusNumRaw)
		statusViper := viper.GetString("tasks.task." + statusNum + "." + "status")
		sum := 1
		numStr := str.Itoa(numRaw)
		numRaw += sum;
		var numColor string;
		var statusTxt string;

		switch statusViper {
			case "completed":
				numColor = misc.Green("✓" + ".")
				statusTxt = misc.Green("// Completed")
			case "uncompleted":
				numColor = misc.Gray(numStr + ".")
				statusTxt = misc.Gray("// Pending")
		}

		title := viper.GetString("tasks.task." + numStr + "." + "title")
		fmt.Printf("%s %s  %s\n", numColor, title, statusTxt)
		statusNumRaw += sum;
	}

	return

}

func View (id string) error {
	viper.New()

	idViper := viper.GetString("tasks.task." + id + "." + "title")
	if idViper == "" {
		return errors.New(misc.Red("No match for number: " + id))
	}

	txt := misc.Cyan("The task number " + id + ": ")
	statusViper := viper.GetString("tasks.task." + id + "." + "status")
	status := misc.Green("Status of task number " + id + ": ") + statusViper

	fmt.Printf("%s%s\n%s\n", txt, idViper, status)

	return nil
}

func Complete (id string) {
	viper.New()
	viper.Set("tasks.task." + id + "." + "status", "completed")
	viper.WriteConfig()
	fmt.Println(misc.Green("✓ Successfully completed task, Cheers!"))
	return
}

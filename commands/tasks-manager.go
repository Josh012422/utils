package command

import (
	"fmt"
	str "strconv"

	"github.com/spf13/viper"
	"github.com/Josh012422/gocharm/misc"
)

// To add a new task

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

// To list all tasks

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

// To view a single task

func View (id string) {
	viper.New()
	taskExists := viper.Get("tasks.task." + id)

	if taskExists == nil {
		fmt.Println(misc.Red("No matching tasks for number: " + id))
		return
	}

	idViper := viper.GetString("tasks.task." + id + "." + "title")

	txt := misc.Cyan("Title of task number " + id + ": ")
	statusViper := viper.GetString("tasks.task." + id + "." + "status")
	status := misc.Green("Status of task number " + id + ": ") + statusViper

	fmt.Printf("%s%s\n%s\n", txt, idViper, status)

	return
}

// To complete tasks

func Complete (id string) {
	viper.New()
	numberRaw := viper.GetInt("tasks.current_number")

	var existingTasks bool;
	var taskExists bool;

	taskExistsViper := viper.Get("tasks.task." + id)

	if taskExistsViper == nil {
		taskExists = false
	} else {
		taskExists = true
	}

	taskAlreadyCompleted := viper.GetString("tasks.task." + id + "." + "status")

	if numberRaw == 0 {
		existingTasks = false
	} else {
	    existingTasks = true
	}

	var taskAlreadyCompletedBool bool;

	if taskAlreadyCompleted == "completed" {
		taskAlreadyCompletedBool = true
	}

	if existingTasks != false && taskAlreadyCompletedBool == false && taskExists == true {
		viper.Set("tasks.task." + id + "." + "status", "completed")
		viper.WriteConfig()

		fmt.Println(misc.Green("✓ Successfully completed task, Cheers!"))
		return
	}

	if taskExists == false {
		fmt.Println(misc.Red("Error: Task does not exists."))
		return
	}

	if existingTasks == false {
		fmt.Println(misc.Red("No existing tasks to complete."))
		return
	}

	if existingTasks != false && taskAlreadyCompletedBool == true {
		fmt.Println(misc.Red("Task already completed ;D"))
		return
	}

	return
}

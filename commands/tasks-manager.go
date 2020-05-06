package command

import (
	"fmt"
	str "strconv"
	"os"
	"strings"

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
	viper.Set("tasks.task." + numberStr + "." + "items_future_number", 1)
	viper.Set("tasks.task." + numberStr + "." + "items_current_number", 0)
	fmt.Printf("%s%s %s%d\n", success, titleColor, numberColor,  numberRaw)
	viper.Set("tasks.current_number", numberRaw)
	viper.WriteConfig()
	numberRaw += sum;
	viper.Set("tasks.future_number", numberRaw)
	viper.WriteConfig()

	return
}

// To add a new item to a task

func AddItem (content string, idRaw int) {
	viper.New()

	id := str.Itoa(idRaw)

//	numberCurr := viper.GetString("tasks.task." + id + "items_current_number")

	numberFutRaw := viper.GetInt("tasks.task." + id + "." + "items_future_number")

	numberFut := str.Itoa(numberFutRaw)

	numberCurr := viper.GetInt("tasks.task." + id + "." + "items_current_number")

	success := misc.Bold(misc.Green("Succesfully added new item to task number:"))

	taskExistsViper := viper.Get("tasks.task." + id)
	var taskExists bool

	if taskExistsViper == nil {
		taskExists = false
		fmt.Println(misc.Red("Error: Task does not exists"))
		os.Exit(1)
		return
	} else {
		taskExists = true
	}

	if taskExists == true {

		viper.Set("tasks.task." + id + "." + "items", nil)
		viper.Set("tasks.task." + id + "." + "items.item", nil)

		viper.Set("tasks.task." + id + "." + "items.item." + numberFut + ".content", content)
		viper.Set("tasks.task." + id + ".items.item." + numberFut + ".status", "uncompleted")
		numberFutRaw += 1
		numberFut = str.Itoa(numberFutRaw)
		numberCurr += 1
		viper.Set("tasks.task." + id + "." + "items_future_number", numberFut)
		viper.Set("tasks.task." + id + "." + "items_current_number", numberCurr)
		viper.WriteConfig()

		fmt.Printf("%s %s. \n%s %s \n", success, id, misc.Bold(misc.Green("Item content:")), content)
		os.Exit(0)
		return
	}
	return

}

// To list all tasks

func List () {
	viper.New()
	numberRaw := viper.GetInt("tasks.current_number")
	numberStr := str.Itoa(numberRaw)

	if numberRaw == 0 {
		fmt.Println(misc.Red("No pending tasks"))
		os.Exit(3)
		return
	}
	statusNumRaw := 1

	numRaw := 1
	itemsNumRaw := 1
	numItemsRaw := 0
	fmt.Sprintf(numberStr)


	for i := 0; i < numberRaw; i++{

		statusNum := str.Itoa(statusNumRaw)
		statusViper := viper.GetString("tasks.task." + statusNum + "." + "status")
		sum := 1
		numStr := str.Itoa(numRaw)
		numRaw += sum;
		var numColor string;
		var statusTxt string;

		itemsNum := viper.GetInt("tasks.task." + numStr + "." + "items_current_number")
		var hasItems bool

		if itemsNum == 0 {
			hasItems = false
		} else {
			hasItems = true
		}

		title := viper.GetString("tasks.task." + numStr + "." + "title")

		items := make([]string, itemsNum)
		itemRaw := 1

		switch statusViper {
			case "completed":
				numColor = misc.Green("✓" + ".")
				statusTxt = misc.Green("// Completed")
			case "uncompleted":
				numColor = misc.Gray(numStr + ".")
				statusTxt = misc.Gray("// Pending")
		}

		// To list all items for a task

		if hasItems == true {
			numIt := 0
			for i:= 0; i < itemsNum; i++{
				numStr := str.Itoa(itemsNumRaw)
				itemStr := str.Itoa(itemRaw)
				items[numIt] = viper.GetString("tasks.task." + numStr + "." + "items.item." + itemStr + ".content")
				itemRaw += 1
				numItemsRaw += 1
				sum := 1
				numIt += sum
			}

			itemsLenght := len(items)
			fmt.Printf("%s %s  %s \n", numColor, title, statusTxt)
			fmtItems := 0
			itemStatusNumRaw := 1

			for i := 0; i < itemsLenght; i++{
				numStr := str.Itoa(itemsNumRaw)
				itemStr := str.Itoa(itemStatusNumRaw)
				bulletColor := "•"
				itemsTxt := items[fmtItems]
				itemStatusViper := viper.GetString("tasks.task." + numStr + ".items.item." + itemStr + ".status")

				if itemStatusViper == "completed" {
					bulletColor = misc.Green("✓")
					itemsTxt = misc.Green(items[fmtItems])
				}

				if itemStatusViper == "uncompleted" {
					bulletColor = "•"
					itemsTxt = items[fmtItems]
				}


				fmt.Printf("     %s %s \n", bulletColor , itemsTxt)
				fmtItems += 1
				itemRaw += 1
				itemStatusNumRaw += 1
			}

		}

		if hasItems == false {

			fmt.Printf("%s %s  %s\n", numColor, title, statusTxt)
		}
		statusNumRaw += sum;
		itemsNumRaw += sum
	}

	return

}

// To view a single task

func View (id string) {
	viper.New()
	taskExists := viper.Get("tasks.task." + id)
	itemsNumViper := viper.GetInt("tasks.task." + id + ".items_current_number")
	items := make([]string, itemsNumViper)
	var hasItems bool

	if itemsNumViper == 0 {
		hasItems = false
	} else {
		hasItems = true
	}

	if taskExists == nil {
		fmt.Println(misc.Red("No matching tasks for number: " + id))
		return
	}

	idViper := viper.GetString("tasks.task." + id + "." + "title")

	txt := misc.Cyan("Title of task number " + id + ": ")
	statusViper := viper.GetString("tasks.task." + id + "." + "status")
	statusViperStr := strings.ToUpper(statusViper)
	status := misc.Green("Status of task number " + id + ": ") + statusViperStr

	fmt.Printf("%s%s\n%s\n", txt, idViper, status)

	if hasItems == true {
		fmt.Printf("%s\n", misc.Yellow("Items:"))
		itemRaw := 0
		itemsNumRaw := 1
		fmtNumRaw := 1
		fmtItems := 0
		bulletColor := "•"
		itemsLenght := len(items)

		itemTxt := items[fmtItems]
		for i := 0; i < itemsNumViper; i++{
			itemsNumStr := str.Itoa(itemsNumRaw)
			items[itemRaw] = viper.GetString("tasks.task." + id + ".items.item." + itemsNumStr + ".content")

			itemRaw += 1
			itemsNumRaw += 1
		}

		for i := 0; i < itemsLenght; i++{
			itemsNumStr := str.Itoa(fmtNumRaw)
			itemStatusViper := viper.GetString("tasks.task." + id + ".items.item." + itemsNumStr + ".status")

			if itemStatusViper == "completed" {
				bulletColor = misc.Green("✓")
				itemTxt = misc.Green(items[fmtItems])
			}

			if itemStatusViper == "uncompleted" {
				bulletColor = "•"
				itemTxt = items[fmtItems]
			}

			fmtNumRaw += 1
			fmtItems += 1

			fmt.Printf("   %s %s\n", bulletColor, itemTxt)
		}
	}

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
		fmt.Println(misc.Green("Task already completed ;D"))
		return
	}

	return
}

// To complete tasks items

func CompleteItem (taskId, itemId string) {
	viper.New()
	taskExistsViper := viper.Get("tasks.task." + taskId)
	itemAlreadyCompleted := viper.GetString("tasks.task." + taskId + ".items.item." + itemId + ".status")

	var taskExists bool
	var itemAlreadyCompletedBool bool

	if itemAlreadyCompleted == "completed" {
		itemAlreadyCompletedBool = true
	} else {
		itemAlreadyCompletedBool = false
	}

	if taskExistsViper == nil {
		taskExists = false
	} else {
		taskExists = true
	}

	if taskExists == false {
		fmt.Println(misc.Red("Error: Task does not exists"))
		os.Exit(35)
		return
	}

	if taskExists == true && itemAlreadyCompletedBool == false {

		viper.Set("tasks.task." + taskId + ".items.item." + itemId + ".status", "completed")
		fmt.Println("tasks.task." + taskId + ".items.item." + itemId + ".status")
		viper.WriteConfig()
		fmt.Println(misc.Green("✓ Succesfully completed item of task " + taskId))
		return
	}

	return
}


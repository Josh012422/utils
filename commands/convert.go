package command

import (
	"fmt"
	"time"

	"github.com/Josh012422/utils/misc"
//	"github.com/spf13/viper"
)

func ConvertTime (tz1 string, tz2 string, hourFormat bool) error {
	loc1raw := tz1
	loc1 := time.LoadLocation(loc1raw)

	loc2raw := tz2
	loc2 := time.LoadLocation(loc2raw)

	if hourFormat == true {
		t1 := time.Now().In(loc1)
		t1.String()
		t1.Format("3:04:05 PM")

		t2 := time.Now().In(loc2)
		t2.String()
		t2.Format("3:04:05 PM")

		return fmt.Println(misc.Cyan("The time in locations are:"), t1, misc.Yellow("Second Timezone"), t2)
	}
}

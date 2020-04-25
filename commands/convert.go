package command

import (
//	"fmt"
	"time"
//	"errors"
	"os"

//	"github.com/Josh012422/utils/misc"
//	"github.com/spf13/viper"
)

func ConvertTime (tz1 string, tz2 string, hourFormat bool) (string, string, error) {
	loc1raw := tz1
	loc1, err1:= time.LoadLocation(loc1raw)

	if err1 != nil {
		return "", "", err1
		os.Exit(1)
	}

	loc2raw := tz2
	loc2, err2 := time.LoadLocation(loc2raw)

	if err2 != nil {
		return "", "", err2
		os.Exit(1)
	}

	if hourFormat == true {
		t1 := time.Now().In(loc1)
		t1.String()
//		t1.Format("3:04:05 PM")

		t2 := time.Now().In(loc2)
		t2.String()
//		t2.Format("3:04:05 PM")

		return t1.Format("3:04:05 PM"), t2.Format("3:04:05 PM"),nil
		os.Exit(0)
	}

	if hourFormat == false {
		t1 := time.Now().In(loc1)
		t1.String()

		t2 := time.Now().In(loc2)
		t2.String()

		return t1.Format("15:04:05 PM"), t2.Format("15:04:05 PM"), nil
		os.Exit(0)
	}

	return "", "", nil
}

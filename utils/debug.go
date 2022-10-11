package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Debug(args ...interface{}) {
	s := ``
	l := len(args)
	if l < 1 {
		fmt.Println("please input some data")
		os.Exit(0)
	}
	i := 1
	for _, v := range args {
		s += fmt.Sprintf("【"+strconv.Itoa(i)+"】: %#v\n", v)
		i++
	}
	s = "******************** 【DEBUG - " + time.Now().Format("2006-01-02 15:04:05") + "】 ********************\n" + s + "******************** 【DEBUG - END】 ********************\n"
	fmt.Println(s)
	os.Exit(0)
}

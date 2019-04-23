package interfaces

import (
	"time"
	"fmt"
)

type IntConv func(op int) int

func timeSpent(inner IntConv)IntConv{
	return func(n int) int {
		start:=time.Now()
		ret:=inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}
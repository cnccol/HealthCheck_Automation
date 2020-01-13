package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v\n",time.Now().AddDate(0, 0, -13).Format("2006-01-02"))
}
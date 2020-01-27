package main

import(
	/* "fmt" */
	"time"
	//"os"
	"github.com/cnccol/HealthCheck_Automation/HealthCheck"
)

func main() {
	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02") 
	name := timeStampString+".txt"
	healthCheck.Run(name)
}
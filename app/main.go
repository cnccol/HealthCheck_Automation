package main

import(
	"fmt"
	"time"
	"os"
	hc "./HealthCheck"
)

func main() {
	day := time.Now().AddDate(0, 0, -37).Format("2006-01-02")
	f, err := os.Stat("face_rec/CSVs_cnn_hog/"+ day +"_1.csv")

	if err != nil{
		panic(err)
	}

	fmt.Printf("Day = %v, Size = %v\n", day, f.ModTime())
	fmt.Printf("%v\n",hc.GetVideoHealth("test"))
}
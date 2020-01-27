package healthCheck

import(
	"time"
	"math/rand"
	"github.com/cnccol/HealthCheck_Automation/Common"
)

func Run(day string) {
	for {
		time.Sleep(5 * time.Second)
		v := Common.Data{ Id: rand.Intn(1000), VideoName: day ,VideoSize: CheckFileSize("videos/"+day) }
		err := Send(v)
		if err != nil {
			panic(err)
		}
	}
}
package healthCheck

import(
	"time"
)

func Run() {
	for {
		time.Sleep(3 * time.Second)
		Send()
	}
}
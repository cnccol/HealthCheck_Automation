package healthCheck

import(
	"time"
	"math/rand"
	"fmt"
)

type Data struct {
	Id int `json:"Id"`
	VideoSize int `json:"VideoSize"`
}
func (p Data) String() string{
	return fmt.Sprintf("{Id: %v, VideoSize: %v}", p.Id, p.VideoSize)
}
func Run() {
	for {
		time.Sleep(3 * time.Second)
		v := Data{Id: rand.Intn(1000), VideoSize: rand.Intn(1000)}
		err := Send(v)
		if err != nil {
			panic(err)
		}
	}
}
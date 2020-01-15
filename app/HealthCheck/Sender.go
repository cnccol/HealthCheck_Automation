package healthCheck

import(
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
)

func Send(){
	netClient := http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get("http://127.0.0.1:9000/data")
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	fmt.Printf("Get %v\n", string(body))

}
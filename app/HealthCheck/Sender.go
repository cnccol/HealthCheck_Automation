package healthCheck

import(
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/cnccol/HealthCheck_Automation/Common"
)

func Send(d Common.Data) error {
	
	netClient := http.Client{
		Timeout: time.Second * 10,
	}
	vals, err :=json.Marshal(d)

	if err != nil {
		return err
	}

	response, err := netClient.Post("http://127.0.0.1:9000/data", "application/json", bytes.NewBuffer(vals))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Get %v\n", string(body))
	return nil
}
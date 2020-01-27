package Common

import (
	"fmt"
)


type Data struct {
	Id int `json:"Id"`
	VideoName string `json:"VideoName"`
	VideoSize int64 `json:"VideoSize"`
}

func (p Data) String() string{
	return fmt.Sprintf("{Id: %v,VideoName: %v ,VideoSize: %v}\n", p.Id, p.VideoName, p.VideoSize)
}
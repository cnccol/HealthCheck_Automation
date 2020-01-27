package healthCheck

import (
	"os"
)

func CheckFileSize(name string) int64{
	fi, err := os.Stat(name)

	if err != nil {
	  return 0
	}

	return fi.Size()
}
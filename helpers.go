package helpers

import (
  "fmt"
  "strconv"
)

//GetInt attempts to retrieve an integer from the user
func GetInt() int {
	var (
		s string
		d int64
	)
	for ok := true; ok; ok = true {
		_, err := fmt.Scanf("%v", &s)
		if err != nil {
			fmt.Print("retry: ")
			continue
		}
		d, err = strconv.ParseInt(s, 0, 64)
		if err != nil {
			fmt.Print("retry: ")
			continue
		} else {
			break
		}
	}
	return int(d)
}
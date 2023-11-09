package sqlcHandler

import (
	"log"
	"strconv"
)

func ConvertStringFloat32(result *string) *float32 {
	if result != nil {
		number, err := strconv.ParseFloat(*result, 32)
		if err != nil {
			log.Fatal(err)
		}
		number_32 := float32(number)
		return &number_32
	}
	return nil
}

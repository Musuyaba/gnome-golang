package sqlcHandler

func ConvertStringFloat32(result *string) *float32 {
	if result != nil {
		number := float32(2.4)
		return &number
		// number, err := strconv.ParseFloat(*result, 32)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// number_32 := float32(number)
		// return &number_32
	}
	return nil
}

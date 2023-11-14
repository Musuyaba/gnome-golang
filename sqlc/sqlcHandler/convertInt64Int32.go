package sqlcHandler

func ConvertInt64Int32(result *int64) *int32 {
	number := int32(*result)
	return &number
}

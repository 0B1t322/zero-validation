package rule

type Number interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint16 | uint64
}

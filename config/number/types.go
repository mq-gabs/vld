package number

type typeUint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type typeInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		typeUint
}

type typeFloat interface {
	~float32 | ~float64
}

type typeNumber interface {
	typeInt | typeFloat
}

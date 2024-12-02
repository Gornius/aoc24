package mathutils

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[N Number](x N) N {
	if x < 0 {
		return -x
	}

	return x
}

package images

const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float32 = 1 << (10 * iota)
	MB
)

var MaxSize = 10 * MB

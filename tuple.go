package collection

// Tuple is a generic ordered pair of values, used by helpers like Zip.
type Tuple[A any, B any] struct {
	First  A
	Second B
}

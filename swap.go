package evaluation2

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

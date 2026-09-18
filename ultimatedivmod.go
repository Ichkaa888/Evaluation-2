package evaluation2

func UltimateDivMod(a *int, b *int) {
	mod := *a % *b
	*a = *a / *b
	*b = mod
}

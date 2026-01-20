package dictionary

func lookupStrategy(word string) (prefixLen int, maxDist int) {
	l := len(word)

	switch {
	case l <= 2:
		return l, 0
	case l <= 4:
		return 2, 1
	case l <= 7:
		return 3, 2
	default:
		return 4, 3
	}
}

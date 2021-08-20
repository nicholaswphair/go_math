package calc

// First letter of names must be uppercase to export
func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

func Sub(a, b int) int {
	return a - b
}

package package1

func Add(a int, b int) int {
	return a + b
}

// in the c programming language this is a bad thing
func Add_pointers(a *int, b *int) *int {
	c := *a + *b
	return &c
}

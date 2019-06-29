package iteration

func CharacterCount(s string) int {
	var count int
	for range s {
		count += 1
	}
	return count
}

package iteration

func CharacterCount(s string) int {
	count := 0
	for range s {
		count += 1
	}
	return count
}

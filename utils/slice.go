package utils

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	former := []rune(s[i])
	last := []rune(s[j])
	if len(former) > len(last) {
		return true
	}
	return false
}

func (s StringSlice) Swap(i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

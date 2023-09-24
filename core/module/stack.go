package module

type Stack struct {
	data []string
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack) Top() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1]
}

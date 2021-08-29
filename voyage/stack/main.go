package main

type StackIF interface {
	Pop() string
	Push(string)
}

type Stack struct {
	StackIF
	Limit int
	Value []string
}

var _ StackIF = &Stack{}

func (s *Stack) Pop() string {
	if len(s.Value) == 0 {
		return ""
	}
	pops := s.Value[len(s.Value)-1]
	s.Value = s.Value[:len(s.Value)-1]
	return pops
}

func (s *Stack) Push(ss string) {

	// Limitに応じて古いデータを削除
	if len(s.Value) >= s.Limit {
		s.Value = s.Value[1:len(s.Value)]
	}

	s.Value = append(s.Value, ss)
}

func main() {

}

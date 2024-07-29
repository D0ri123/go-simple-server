package cmd

// main은 cmd를 호출시킬 예정, 모든 항목들을 Cmd 스트럭쳐 안에 구성할 예정
type Cmd struct{}

func NewCmd() *Cmd {
	c := &Cmd{}
	return c
}

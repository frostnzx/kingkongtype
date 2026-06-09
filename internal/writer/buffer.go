package writer

type Buffer struct {
	Text      string
	InputText string
	Position  int
}

func NewBuffer(s string) *Buffer {
	return &Buffer{Text: s, InputText: "", Position: 0}
}

func (b *Buffer) InsertNextChar(c string) {
	b.InputText += c
	b.Position += 1
}
func (b *Buffer) Pop() {
	b.InputText = b.InputText[:len(b.InputText)-1]
	b.Position -= 1
}
func (b *Buffer) CheckPos(pos int) bool {
	if pos >= min(len(b.Text), len(b.InputText)) {
		return false
	}
	return b.Text[pos] == b.InputText[pos]
}
func (b *Buffer) CheckCurrentPos() bool {
	return b.Text[b.Position] == b.InputText[b.Position]
}

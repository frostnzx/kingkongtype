package data

import "testing"

func TestIsDifficultyMatch(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		difficulty int
		want       bool
	}{
		{name: "easy accepts short quote", text: stringOfLength(80), difficulty: 0, want: true},
		{name: "easy rejects medium quote", text: stringOfLength(81), difficulty: 0, want: false},
		{name: "medium accepts lower bound", text: stringOfLength(81), difficulty: 1, want: true},
		{name: "medium accepts upper bound", text: stringOfLength(160), difficulty: 1, want: true},
		{name: "medium rejects hard quote", text: stringOfLength(161), difficulty: 1, want: false},
		{name: "hard accepts long quote", text: stringOfLength(161), difficulty: 2, want: true},
		{name: "hard rejects medium quote", text: stringOfLength(160), difficulty: 2, want: false},
		{name: "unknown difficulty accepts any quote", text: "anything", difficulty: 99, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDifficultyMatch(tt.text, tt.difficulty)
			if got != tt.want {
				t.Fatalf("IsDifficultyMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringOfLength(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = 'a'
	}
	return string(b)
}

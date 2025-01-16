package wordsearch

import "testing"

func exist(board [][]byte, word string) bool {
	if len(word) > len(board)*len(board[0]) {
		return false
	}

	chars := []byte(word)
	visited := make([][]bool, len(board))
	for i := range len(board) {
		visited[i] = make([]bool, len(board[0]))
	}

	var check func(x, y, pos int) bool

	check = func(x, y, pos int) bool {
		if x < 0 || y < 0 || x == len(board[0]) || y == len(board) {
			return false
		}
		if visited[y][x] {
			return false
		}
		if board[y][x] == chars[pos] {
			if pos == len(chars)-1 {
				return true
			}
			visited[y][x] = true
			if check(x, y-1, pos+1) || check(x, y+1, pos+1) || check(x-1, y, pos+1) || check(x+1, y, pos+1) {
				return true
			}
			visited[y][x] = false
		}

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if check(j, i, 0) {
				return true
			}
		}
	}

	return false
}

func TestExist(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
		{
			board: [][]byte{
				{'A'},
			},
			word: "A",
			want: true,
		},
		{
			board: [][]byte{
				{'A'},
			},
			word: "B",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := exist(tt.board, tt.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

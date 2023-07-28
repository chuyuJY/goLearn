package main

const RFC3339Mill = "2006-01-02T15:04:05.999Z07:00"

type b []int

func main() {
}

func exist(board [][]byte, word string) bool {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	var dfs func(index, i, j int) bool
	dfs = func(index, i, j int) bool {
		if index == len(word)-1 {
			return true
		}
		flag := false
		for _, dir := range dirs {
			nextI, nextJ := i+dir[0], j+dir[1]
			if 0 <= nextI && nextI < len(board) && 0 <= nextJ && nextJ < len(board[0]) && !visited[nextI][nextJ] && board[nextI][nextJ] == word[index+1] {
				visited[nextI][nextJ] = true
				flag = dfs(index+1, nextI, nextJ)
				visited[nextI][nextJ] = false
			}
			if flag {
				break
			}
		}
		return flag
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = true
			if board[i][j] == word[0] && dfs(0, i, j) {
				return true
			}
			visited[i][j] = false
		}
	}
	return false
}

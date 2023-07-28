package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordGame struct {
	Word      string
	WordParts []string
}

func NewWordGame(word string) *WordGame {
	wg := &WordGame{Word: word}
	for i := 0; i < len(wg.Word); i++ {
		if i < len(wg.Word)-1 && wg.Word[i+1] == ' ' {
			wg.WordParts = append(wg.WordParts, wg.Word[i:i+2])
			i++ // 跳过 ' '
		} else {
			wg.WordParts = append(wg.WordParts, string(wg.Word[i]))
		}
	}
	return wg
}

func (wg *WordGame) TryConnect(currentIndex, currentPos, touchedIndex, touchedPos int) {
	// 判断输入是否合法
	if currentIndex == touchedIndex || currentIndex < 0 || touchedIndex < 0 || currentIndex >= len(wg.WordParts) || touchedIndex >= len(wg.WordParts) || currentPos^touchedPos >= 0 {
		fmt.Println("invalid input")
		return
	}
	// 设置拼接字符串的位置
	var leftPart, rightPart string
	if currentPos > 0 {
		leftPart, rightPart = wg.WordParts[currentIndex], wg.WordParts[touchedIndex]
	} else {
		leftPart, rightPart = wg.WordParts[touchedIndex], wg.WordParts[currentIndex]
	}
	// 判断是否可以拼接
	temp := wg.WordParts[currentIndex]
	wg.WordParts[currentIndex] = leftPart + rightPart
	wg.WordParts[touchedIndex], wg.WordParts[len(wg.WordParts)-1] = wg.WordParts[len(wg.WordParts)-1], wg.WordParts[touchedIndex]
	// 首先判断是否含有拼接成的字符串(剪枝)，然后判断拼接之后的 wordParts 是否还可以组成 word
	ok := strings.Contains(wg.Word, wg.WordParts[currentIndex]) && isValid(wg.Word, wg.WordParts[:len(wg.WordParts)-1])
	wg.WordParts[touchedIndex], wg.WordParts[len(wg.WordParts)-1] = wg.WordParts[len(wg.WordParts)-1], wg.WordParts[touchedIndex]
	if ok { // 若可以拼接，则删除原来的两个字符串，添加新的字符串
		wg.WordParts = append(wg.WordParts[:touchedIndex], wg.WordParts[touchedIndex+1:]...)
		fmt.Println("combine success")
	} else { // 若不可以拼接，则恢复原来的字符串
		wg.WordParts[currentIndex] = temp
		fmt.Println("combine failed")
	}

	// 打印当前的 wordParts
	wg.printWordParts()
}

// 判断 wordParts 是否可以组成 word
func isValid(s string, wordParts []string) bool {
	var backTrack func(s string, start int) bool
	backTrack = func(s string, start int) bool {
		if s == "" {
			return true
		}
		flag := false
		for i := start; i < len(wordParts); i++ {
			// 回溯
			if strings.HasPrefix(s, wordParts[i]) {
				wordParts[i], wordParts[start] = wordParts[start], wordParts[i]
				flag = backTrack(s[len(wordParts[start]):], start+1)
				wordParts[i], wordParts[start] = wordParts[start], wordParts[i]
			}
			// 若成立，则直接返回
			if flag {
				break
			}
		}
		return flag
	}
	return backTrack(s, 0)
}

func (wg *WordGame) printWordParts() {
	fmt.Println("current word parts:", wg.WordParts)
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	fmt.Print("input word: ")
	word, _ := sc.ReadString('\n')
	word = strings.Trim(word, "\r\n")
	wg := NewWordGame(word)
	wg.printWordParts()
	fmt.Println("input format [currentIndex, currentPos, touchedIndex, touchedPos], Position: 1 for left, -1 for right")
	for {
		var currentIndex, currentPos, touchedIndex, touchedPos int
		fmt.Print("input: ")
		fmt.Fscanf(sc, "%d, %d, %d, %d\n", &currentIndex, &currentPos, &touchedIndex, &touchedPos)
		wg.TryConnect(currentIndex, currentPos, touchedIndex, touchedPos)
		if len(wg.WordParts) == 1 {
			fmt.Println("game over")
			break
		}
	}
}

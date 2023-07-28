package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//type WordGame struct {
//	Word      string
//	WordParts []string
//	WordMap   map[string]letterMap
//}
//
//type letterMap struct {
//	left, right map[string]int
//}
//
//func NewWordGame(word string) *WordGame {
//	wg := &WordGame{Word: word}
//	wg.setWordParts()
//	wg.setWordMap()
//	return wg
//}
//
//func (wg *WordGame) TryConnect(currentIndex, currentPos, touchedIndex, touchedPos int) {
//	// 判断输入是否合法
//	if currentIndex == touchedIndex || currentIndex < 0 || touchedIndex < 0 || currentIndex >= len(wg.WordParts) || touchedIndex >= len(wg.WordParts) || currentPos^touchedPos >= 0 {
//		//fmt.Println("invalid input")
//		return
//	}
//	// 设置拼接字符串的位置
//	var leftPart, rightPart string
//	if currentPos > 0 {
//		leftPart, rightPart = wg.WordParts[currentIndex], wg.WordParts[touchedIndex]
//	} else {
//		leftPart, rightPart = wg.WordParts[touchedIndex], wg.WordParts[currentIndex]
//	}
//	// 设置拼接字符串的首尾
//	var leftEnd, rightStart string
//	if len(leftPart) > 1 && leftPart[len(leftPart)-1] == ' ' {
//		leftEnd = leftPart[len(leftPart)-2:]
//	} else {
//		leftEnd = leftPart[len(leftPart)-1:]
//	}
//	if len(rightPart) > 1 && rightPart[1] == ' ' {
//		rightStart = rightPart[:2]
//	} else {
//		rightStart = rightPart[:1]
//	}
//	// 判断是否可以拼接
//	flag := false
//	if wg.WordMap[leftEnd].right[rightStart] > 0 && wg.WordMap[rightStart].left[leftEnd] > 0 {
//		temp := wg.WordParts[currentIndex]
//		wg.WordParts[currentIndex] = leftPart + rightPart
//		wg.WordParts[touchedIndex], wg.WordParts[len(wg.WordParts)-1] = wg.WordParts[len(wg.WordParts)-1], wg.WordParts[touchedIndex]
//		if strings.Contains(wg.Word, wg.WordParts[currentIndex]) && isValid(wg.Word, wg.WordParts[:len(wg.WordParts)-1]) {
//			wg.WordParts[touchedIndex], wg.WordParts[len(wg.WordParts)-1] = wg.WordParts[len(wg.WordParts)-1], wg.WordParts[touchedIndex]
//			wg.WordMap[leftEnd].right[rightStart]--
//			wg.WordMap[rightStart].left[leftEnd]--
//			wg.WordParts = append(wg.WordParts[:touchedIndex], wg.WordParts[touchedIndex+1:]...)
//			flag = true
//		} else {
//			wg.WordParts[touchedIndex], wg.WordParts[len(wg.WordParts)-1] = wg.WordParts[len(wg.WordParts)-1], wg.WordParts[touchedIndex]
//			wg.WordParts[currentIndex] = temp
//		}
//	}
//
//	if flag {
//		//fmt.Println("combine success")
//	} else {
//		//fmt.Println("combine failed")
//	}
//	//wg.printWordParts()
//}
//func isValid(s string, wordDict []string) bool {
//	var backTrack func(s string, start int) bool
//	backTrack = func(s string, start int) bool {
//		if s == "" {
//			return true
//		}
//		flag := false
//		for i := start; i < len(wordDict); i++ {
//			if strings.HasPrefix(s, wordDict[i]) {
//				wordDict[i], wordDict[start] = wordDict[start], wordDict[i]
//				flag = backTrack(s[len(wordDict[start]):], start+1)
//				wordDict[i], wordDict[start] = wordDict[start], wordDict[i]
//			}
//			if flag {
//				break
//			}
//		}
//		return flag
//	}
//	return backTrack(s, 0)
//}
//
//func (wg *WordGame) setWordParts() {
//	for i := 0; i < len(wg.Word); i++ {
//		if i < len(wg.Word)-1 && wg.Word[i+1] == ' ' {
//			wg.WordParts = append(wg.WordParts, wg.Word[i:i+2])
//			i++ // 跳过 ' '
//		} else {
//			wg.WordParts = append(wg.WordParts, string(wg.Word[i]))
//		}
//	}
//}
//
//func (wg *WordGame) setWordMap() {
//	wg.WordMap = map[string]letterMap{}
//	for i := 0; i < len(wg.WordParts); i++ {
//		if _, ok := wg.WordMap[wg.WordParts[i]]; !ok {
//			wg.WordMap[wg.WordParts[i]] = letterMap{left: map[string]int{}, right: map[string]int{}}
//		}
//		if i > 0 {
//			wg.WordMap[wg.WordParts[i]].left[wg.WordParts[i-1]]++
//		}
//		if i < len(wg.WordParts)-1 {
//			wg.WordMap[wg.WordParts[i]].right[wg.WordParts[i+1]]++
//		}
//	}
//}
//
//func (wg *WordGame) printWordParts() {
//	fmt.Println("current word parts:", wg.WordParts)
//}
//
//func main() {
//	sc := bufio.NewReader(os.Stdin)
//	fmt.Print("input word: ")
//	word, _ := sc.ReadString('\n')
//	word = strings.Trim(word, "\r\n")
//	wg := NewWordGame(word)
//	wg.printWordParts()
//	fmt.Println("input format [currentIndex, currentPos, touchedIndex, touchedPos], Position: 1 for left, -1 for right")
//	for {
//		var currentIndex, currentPos, touchedIndex, touchedPos int
//		fmt.Print("input: ")
//		fmt.Fscanf(sc, "%d, %d, %d, %d\n", &currentIndex, &currentPos, &touchedIndex, &touchedPos)
//		wg.TryConnect(currentIndex, currentPos, touchedIndex, touchedPos)
//		if len(wg.WordParts) == 1 {
//			fmt.Println("game over")
//			break
//		}
//	}
//}

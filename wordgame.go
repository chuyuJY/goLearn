//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//type GameWord struct {
//	Content string
//	X       int
//	Y       int
//	Color   string
//}
//
//type WordGame struct {
//	Slices []GameWord
//}
//
//func NewWordGame(word string) *WordGame {
//	slices := []GameWord{}
//	for _, c := range word {
//		slices = append(slices, GameWord{Content: string(c), X: 0, Y: 0})
//	}
//	return &WordGame{Slices: slices}
//}
//
//func (wg *WordGame) TryConnect(currentIndex int, touchedIndex int) {
//	slices := make([]GameWord, len(wg.Slices))
//	copy(slices, wg.Slices)
//	currentSlice := slices[currentIndex]
//	touchedSlice := slices[touchedIndex]
//	removeIndex := -1
//	combinedContent := ""
//	positionIndex := -1
//	if currentIndex > 0 && slices[currentIndex-1].Content == touchedSlice.Content {
//		// 1. touched 在 current 的左边
//		removeIndex = currentIndex
//		combinedContent = touchedSlice.Content + currentSlice.Content
//		wg.swap(slices, touchedIndex, currentIndex-1)
//	} else if currentIndex < len(slices)-1 && slices[currentIndex+1].Content == touchedSlice.Content {
//		// 2. touched 在 current 的右边
//		removeIndex = currentIndex
//		combinedContent = currentSlice.Content + touchedSlice.Content
//		positionIndex = currentIndex
//		wg.swap(slices, touchedIndex, currentIndex+1)
//	} else if touchedIndex > 0 && slices[touchedIndex-1].Content == currentSlice.Content {
//		// 3. current 在 touched 左边
//		combinedContent = currentSlice.Content + touchedSlice.Content
//		removeIndex = touchedIndex - 1
//		positionIndex = currentIndex
//		wg.swap(slices, currentIndex, touchedIndex-1)
//		positionIndex = touchedIndex - 1
//	} else if touchedIndex < len(slices)-1 && slices[touchedIndex+1].Content == currentSlice.Content {
//		// 4. current 在 touched 右边
//		combinedContent = touchedSlice.Content + currentSlice.Content
//		removeIndex = touchedIndex + 1
//		wg.swap(slices, currentIndex, touchedIndex+1)
//	}
//	if removeIndex != -1 {
//		fmt.Println("combine success")
//		fmt.Printf("before combine: %s\n", wg.printWord())
//		slices[touchedIndex].Content = combinedContent
//		tIndex := touchedIndex
//		if touchedIndex < removeIndex {
//			tIndex = touchedIndex
//		} else {
//			tIndex = touchedIndex - 1
//		}
//		slices = append(slices[:removeIndex], slices[removeIndex+1:]...)
//		if positionIndex != -1 {
//			wg.swap(slices, positionIndex, tIndex)
//		}
//		wg.Slices = slices
//		fmt.Printf("after combine: %s\n", wg.printWord())
//	} else {
//		fmt.Println("can not combine")
//	}
//}
//
//func (wg *WordGame) swap(slices []GameWord, i int, j int) {
//	if i == j {
//		return
//	}
//	temp := slices[i]
//	slices[i] = slices[j]
//	slices[j] = temp
//}
//
//func (wg *WordGame) printWord() string {
//	var sb strings.Builder
//	sb.WriteString("[")
//	for i, slice := range wg.Slices {
//		if i > 0 {
//			sb.WriteString(",")
//		}
//		sb.WriteString(slice.Content)
//	}
//	sb.WriteString("]")
//	return sb.String()
//}
//
//func main() {
//	game := NewWordGame("Seed")
//	game.TryConnect(0, 3)
//	game.TryConnect(2, 0)
//	game.TryConnect(0, 2)
//	game.TryConnect(0, 1)
//	game.TryConnect(1, 0)
//}

package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"sort"
)
type runeSlice []rune
func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func main() {
	runeMap := make(map[rune]bool)


	srcFilePtr, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("打开文件错误")
		fmt.Println(err)
		return
	}
	defer srcFilePtr.Close()

	bufReader := bufio.NewReader(srcFilePtr)

	expFilePtr, err := os.OpenFile("./result.txt", os.O_CREATE | os.O_EXCL  | os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("当前路径存在result.txt，删除后重试")
		fmt.Println(err)
		return
	}
	defer expFilePtr.Close()

	bufWriter := bufio.NewWriter(expFilePtr)

	for {
		line, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		for _, r := range string(line){
			runeMap[r] = true
		}
	}

	result := make([]rune, 0, 1<<10)
	for r := range runeMap {
		result = append(result, r)
	}
	sort.Sort(runeSlice(result))


	bufWriter.Write([]byte(string(result)))
	if err != nil {
		fmt.Println(err)
		return
	}

	bufWriter.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("完成")
}

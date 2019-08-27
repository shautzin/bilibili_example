package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const DepthFolder = "depth45p" // 帧数据文件夹

// 颜色深度 <-> 终端颜色 映射表
var DepthMapping = map[int]int{
	0: 0, 1: 232, 2: 233, 3: 234, 4: 235, 5: 236, 6: 237, 7: 238, 8: 239, 9: 240,
	10: 241, 11: 242, 12: 243, 13: 244, 14: 245, 15: 246, 16: 247, 17: 248,
	18: 249, 19: 250, 20: 251, 21: 252, 22: 253, 23: 254, 24: 255, 25: 15,
}

// 打印单个色块，两个空格组成一个块
func PrintBlock(depth int) {
	fmt.Printf("\x1b[0;48;5;%dm  \x1b[0m", DepthMapping[depth])
}

// 读取帧数据
func ReadFrame(c int) {
	fileName := fmt.Sprintf("%s/%d.txt", DepthFolder, c)
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		PrintLine(line) // 单行输出
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

// 输出单行
func PrintLine(line string) {
	if line != "" {
		depths := strings.Split(line, ",")
		for _, depthStr := range depths {
			if depthStr != "" {
				depth, _ := strconv.Atoi(depthStr)
				PrintBlock(depth / 10) // 每 10 一个阶
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Print("\x1b[2J")   // 全屏
	fmt.Print("\x1b[?25l") // 隐藏光标

	// for 循环所有帧
	for i := 1; i < 6569; i++ {
		ReadFrame(i)
		time.Sleep(30 * time.Millisecond) // 30ms / 帧
		fmt.Print("\x1b[45A")             // 回首行
	}

	fmt.Print("\x1b[?25h") // 显示光标
}

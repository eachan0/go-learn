package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/shirou/gopsutil/process"
	"log"
	"sync"
	"time"
)

var (
	// 局数计数器
	count = 1
	// 操作鼠标的锁
	mouseLock = sync.Mutex{}
	// 资源文件夹路径
	folder = "D:\\lol\\"
	// 游戏结束标记
	gameStartFlag *bool
	// 操作游戏结束标记的锁
	gameStartLock = sync.Mutex{}
	// 游戏结束标记
	gameEndFlag *bool
	// 操作游戏结束标记的锁
	gameEndLock = sync.Mutex{}
)

func init() {
	// 配置日志
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	// 初始化游戏开始标记
	tempStart := false
	gameStartFlag = &tempStart
	// 初始化游戏开始标记
	tempEnd := false
	gameEndFlag = &tempEnd
}

func main() {
	/** 开始游戏 **/
	res := findAndClickButton("1.png")
	logger("开始寻找对局")
	if res {
		// 接受,进入游戏,拒绝
		acceptGame()
	} else {
		logger("没有找到寻找对局按钮,请确定LOL窗口是否进入云顶房间,程序退出")
		return
	}
	/** 游戏过程中 **/
	// 选秀
	go choice()
	// 捡装备
	go equipment()
	// 捡钱
	go money()
	// 买英雄
	go buyHero()
	// 升级
	go upgrade()
	/** 退出游戏 **/
	end := make(chan bool)
	// 正常退出
	go normalExit(end)
	// 兜底强制退出
	go forceExit(end)
	<-end
	// 打印对局完成次数
	logger("完成对局")
	// 返回房间
	backHouse()
	// 1s延时后进行下一把游戏
	time.Sleep(time.Second * 1)
	count++
	main()
}

// 升级
func upgrade() {
	// 游戏退出,结束协程
	if getGameEndState() {
		return
	}
	// 按钮位置固定,隔段时间点击即可
	time.Sleep(time.Minute * 2)
	// 游戏退出,结束协程
	if getGameEndState() {
		return
	}
	// todo: 填写位置
	moveAndClick(1, 1, 5)
	upgrade()
}

// 买英雄
func buyHero() {
	// 游戏退出,结束协程
	if getGameEndState() {
		return
	}
	// 按钮位置固定,隔段时间点击即可
	time.Sleep(time.Minute * 1)
	// todo: 填写商品栏1号英雄位置
	moveAndClick(1, 1)
	// todo: 填写商品栏2号英雄位置
	moveAndClick(1, 1)
	buyHero()
}

// 捡装备
func equipment() {
	// 游戏退出,结束协程
	if getGameEndState() {
		return
	}
	// 按钮位置固定,隔段时间点击即可
	time.Sleep(time.Minute * 1)
	findAndClickButton("装备.png")
	equipment()
}

// 捡钱
func money() {
	// 游戏退出,结束协程
	if getGameEndState() {
		return
	}
	// 按钮位置固定,隔段时间点击即可
	time.Sleep(time.Minute * 1)
	findAndClickButton("金币.png")
	equipment()
}

// 选秀
func choice() {
	// 游戏退出,结束协程
	if getGameEndState() {
		return
	}
	time.Sleep(time.Second * 30)
	// todo: 选秀图片
	findAndClickButton("选秀.png")
}

// 返回房间
func backHouse() {
	// 活动有OK按钮
	findAndClickButton("ok.png")
	time.Sleep(time.Second * 3)
	findAndClickButton("再来一次.png")
}

// 正常退出
func normalExit(end chan bool) {
	for true {
		if getGameEndState() {
			return
		}
		_, _, find := findImageInScreenPosition("33.png")
		if find {
			exitGame()
			end <- true
			return
		}
		time.Sleep(time.Minute * 1)
	}
}

// 强制退出
func forceExit(end chan bool) {
	// 15分钟
	time.Sleep(time.Minute * 15)
	exitGame()
	end <- true
}

// 退出游戏
func exitGame() {
	if changeGameEndState(true) {
		// 按下 ESC键
		robotgo.KeyTap(`esc`)
		time.Sleep(time.Millisecond * 500)
		// todo: 投降图片
		findAndClickButton("投降.png")
		time.Sleep(time.Second * 6)
		// todo: 确定图片
		findAndClickButton("确定.png")
	}
}

// 接受游戏
func acceptGame() {
	// 开启判断是否进入游戏协程
	go startedGame()
	for true {
		if getGameStartState() {
			return
		}
		// 点击接受
		find := findAndClickButton("2.png")
		if find {
			logger("接受游戏对局")
		}
		time.Sleep(time.Second * 3)
	}
}

// 判断进入游戏
func startedGame() {
	for true {
		// 获取进程列表
		processes, _ := process.Processes()
		for _, p := range processes {
			name, _ := p.Name()
			// 通过游戏进程名称来判断是否进入游戏
			if name == "League of Legends.exe" {
				if changeGameStartState(true) {
					logger("已开启游戏进程,进入游戏")
				}
				return
			}
		}
		time.Sleep(time.Second * 5)
	}
}

// 找图片在屏幕中的位置
func findImageInScreenPosition(filename string) (int, int, bool) {
	// 获取指定素材的路径
	path := fmt.Sprintf("%s%s", folder, filename)
	// 打开素材位图
	bitmap := robotgo.OpenBitmap(path)
	if bitmap == nil {
		return 0, 0, false
	}
	// 释放位图
	defer robotgo.FreeBitmap(bitmap)
	// 找位置
	x, y := robotgo.FindBitmap(bitmap)
	if x < 0 || y < 0 {
		return 0, 0, false
	}
	return x, y, true
}

// 寻找按钮并点击
func findAndClickButton(filename string) bool {
	x, y, find := findImageInScreenPosition(filename)
	if find {
		moveAndClick(x, y)
		return true
	}
	return false
}

// 鼠标移动到指定位置点击指定次数
func moveAndClick(x, y int, args ...int) {
	mouseLock.Lock()
	// 移动鼠标后
	robotgo.MoveMouseSmooth(x, y, .7, .7)
	// 多击
	if len(args) > 1 {
		num := args[0]
		for i := 0; i < num; i++ {
			robotgo.MouseClick("left", false)
		}
	} else {
		// 单击
		robotgo.MouseClick("left", false)
	}
	mouseLock.Unlock()
}

// 标记游戏开始状态
func changeGameStartState(state bool) bool {
	gameStartLock.Lock()
	if state && *gameStartFlag {
		return false
	}
	*gameStartFlag = state
	gameStartLock.Unlock()
	return true
}

// 获取游戏开始状态
func getGameStartState() (state bool) {
	gameStartLock.Lock()
	state = *gameStartFlag
	gameStartLock.Unlock()
	return
}

// 标记游戏结束状态
func changeGameEndState(state bool) bool {
	gameEndLock.Lock()
	if state && *gameEndFlag {
		return false
	}
	*gameEndFlag = state
	gameEndLock.Unlock()
	return true
}

// 获取游戏结束状态
func getGameEndState() (state bool) {
	gameEndLock.Lock()
	state = *gameEndFlag
	gameEndLock.Unlock()
	return
}

// 打印日志
func logger(msg string) {
	log.Printf("游戏对局-%d: %s\n", count, msg)
}

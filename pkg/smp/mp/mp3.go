package mp

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) {

	fmt.Println("Playing MP3 music", source)
	path := filepath.FromSlash(source)
	fmt.Println("Playing MP3 music_2", path)
	// 打开MP3文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 解码MP3文件
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// 初始化扬声器
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个控制器，用于暂停和恢复播放
	ctrl := &beep.Ctrl{Streamer: streamer, Paused: false}

	// 播放控制器
	speaker.Play(ctrl)

	fmt.Println("正在播放MP3文件...")
	fmt.Println("输入p来暂停或恢复播放，输入q来退出程序")

	var input string
	for {
		fmt.Scanln(&input)
		switch input {
		case "p":
			ctrl.Paused = !ctrl.Paused // 切换暂停状态
			if ctrl.Paused {
				fmt.Println("已暂停播放")
			} else {
				fmt.Println("已恢复播放")
			}
		case "q":
			fmt.Println("已退出程序")
			return
		default:
			fmt.Println("无效的输入，请重新输入")
		}
	}

	fmt.Println("\nFinished playing", source)
}

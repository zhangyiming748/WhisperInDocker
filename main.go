package main

import (
	"github.com/zhangyiming748/WhisperInDocker/constant"
	"github.com/zhangyiming748/WhisperInDocker/util"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	setLog()
}
func main() {
	util.ExitAfterRun()
	t := new(util.ProcessDuration)
	t.SetStart(time.Now())
	defer func() {
		log.Printf("程序总用时:%v\n", t.GetDuration().Minutes())
	}()
	if root := os.Getenv("root"); root != "" {
		log.Printf("读取到环境变量，默认$root修改为%s\n", root)
		constant.SetRoot(root)
	} else {
		log.Printf("未读取到环境变量，默认$root修改为%s\n", constant.GetRoot())
	}
	if language := os.Getenv("language"); language != "" {
		log.Printf("读取到环境变量，默认$language修改为%s\n", language)
		constant.SetLanguage(language)
	} else {
		log.Printf("未读取到环境变量，默认$language修改为%s\n", constant.GetLanguage())
	}
	if pattern := os.Getenv("pattern"); pattern != "" {
		log.Printf("读取到环境变量，默认$pattern修改为%s\n", pattern)
		constant.SetPattern(pattern)
	} else {
		log.Printf("未读取到环境变量，默认$pattern修改为%s\n", constant.GetPattern())
	}
	if model := os.Getenv("model"); model != "" {
		log.Printf("读取到环境变量，默认$model修改为%s\n", model)
		constant.SetPattern(model)
	} else {
		log.Printf("未读取到环境变量，默认$model修改为%s\n", constant.GetModel())
	}
	files, _ := util.GetAllFileInfoFast(constant.GetRoot(), constant.GetPattern())
	for _, file := range files {
		log.Printf("文件名:%v\n", file)
		//whisper true.mp4 --model base --language English --model_dir /Users/zen/Whisper --output_format srt
		//cmd := exec.Command("whisper", file.FullPath, "--model", level, "--model_dir", location, "--language", language, "--output_dir", root, "--verbose", "True")
		cmd := exec.Command("whisper", file, "--model", constant.GetModel(), "--model_dir", constant.GetRoot(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", constant.GetLanguage(), "--output_dir", constant.GetRoot(), "--verbose", "True")
		err := util.ExecCommand(cmd)
		if err != nil {
			log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		}
	}
	t.SetEnd(time.Now())
}

func setLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{constant.GetRoot(), "mylog.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}

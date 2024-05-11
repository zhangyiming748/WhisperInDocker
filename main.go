package main

import (
	"fmt"
	"github.com/zhangyiming748/WhisperInDocker/constant"
	"github.com/zhangyiming748/WhisperInDocker/sql"
	"github.com/zhangyiming748/WhisperInDocker/util"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func init() {
	setLog()
	sql.SetEngine()
}
func main() {
	util.ExitAfterRun()
	if root := os.Getenv("root"); root != "" {
		fmt.Printf("读取到环境变量，默认$root修改为%s\n", root)
		constant.SetRoot(root)
	} else {
		fmt.Printf("未读取到环境变量，默认$root修改为%s\n", constant.GetRoot())
	}
	if language := os.Getenv("language"); language != "" {
		fmt.Printf("读取到环境变量，默认$language修改为%s\n", language)
		constant.SetLanguage(language)
	} else {
		fmt.Printf("未读取到环境变量，默认$language修改为%s\n", constant.GetLanguage())
	}
	if pattern := os.Getenv("pattern"); pattern != "" {
		fmt.Printf("读取到环境变量，默认$pattern修改为%s\n", pattern)
		constant.SetPattern(pattern)
	} else {
		fmt.Printf("未读取到环境变量，默认$pattern修改为%s\n", constant.GetPattern())
	}
	if model := os.Getenv("model"); model != "" {
		fmt.Printf("读取到环境变量，默认$model修改为%s\n", model)
		constant.SetPattern(model)
	} else {
		fmt.Printf("未读取到环境变量，默认$model修改为%s\n", constant.GetModel())
	}
	files, _ := util.GetAllFileInfoFast(constant.GetRoot(), constant.GetPattern())
	for _, file := range files {
		slog.Info("文件", slog.String("文件名", file))
		//whisper true.mp4 --model base --language English --model_dir /Users/zen/Whisper --output_format srt
		//cmd := exec.Command("whisper", file.FullPath, "--model", level, "--model_dir", location, "--language", language, "--output_dir", root, "--verbose", "True")
		cmd := exec.Command("whisper", file, "--threads", "0", "--model", constant.GetModel(), "--model_dir", constant.GetRoot(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", constant.GetLanguage(), "--output_dir", constant.GetRoot(), "--verbose", "True")
		err := util.ExecCommand(cmd)
		if err != nil {
			slog.Error("当前字幕生成错误", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("错误原文", err.Error()))
		}
	}
}

func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	file := strings.Join([]string{constant.GetRoot(), "whisper.log"}, string(os.PathSeparator))
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

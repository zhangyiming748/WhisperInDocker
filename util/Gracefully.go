package util

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

var ExitAfterDone = false

func GetExitStatus() bool {
	return ExitAfterDone
}

func SetExitStatus(b bool) {
	slog.Debug("改变退出状态")
	ExitAfterDone = b
}
func ExitAfterRun() {
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			input, _ := reader.ReadString('\n')
			fmt.Printf("You entered is %T\t%v", input, input)
			if input == "q\n" {
				slog.Debug("接收到q")
				//ExitAfterDone = true
				SetExitStatus(true)
				slog.Info("退出状态改变", slog.Bool("新值", ExitAfterDone))
			}
			if input == "q\r\n" {
				slog.Debug("接收到q")
				//ExitAfterDone = true
				SetExitStatus(true)
				slog.Info("windows 退出状态改变", slog.Bool("新值", ExitAfterDone))
			}

		}
	}()
}

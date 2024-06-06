package util

import (
	"bufio"
	"log"
	"os"
)

var ExitAfterDone = false

func GetExitStatus() bool {
	return ExitAfterDone
}

func SetExitStatus(b bool) {
	log.Println("改变退出状态")
	ExitAfterDone = b
}
func ExitAfterRun() {
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			input, _ := reader.ReadString('\n')
			log.Printf("You entered is %T\t%v", input, input)
			if input == "q\n" {
				log.Printf("接收到q\t退出状态改变:%v\n", ExitAfterDone)
				SetExitStatus(true)
			}
			if input == "q\r\n" {
				log.Printf("接收到q\t退出状态改变:%v\n", ExitAfterDone)
				SetExitStatus(true)
			}
		}
	}()
}

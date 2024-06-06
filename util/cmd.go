package util

import (
	"log"
	"os"
	"os/exec"
)

func ExecCommand(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	if output, err := c.CombinedOutput(); err != nil {
		log.Printf("命令执行中产生错误\t命令原文%v\t错误原文%v\n", c.String, err)
		return err
	} else {
		// 这是一段永远不可能被运行的代码
		log.Printf("命令执行完毕\t输出:%v\n", string(output))
	}

	if exit := GetExitStatus(); exit {
		log.Printf("命令端获取到退出状态,命令结束后退出\t最后一条命令", c.String())
		os.Exit(0)
	}
	return nil
}

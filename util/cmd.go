package util

import (
	"log"
	"os/exec"
)

func ExecCommand(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	if output, err := c.CombinedOutput(); err != nil {
		log.Printf("命令:%v执行中产生错误:%v\n", c.String(), err)
		return err
	} else {
		// 这是一段永远不可能被运行的代码
		log.Printf("命令执行完毕\t输出:%v\n", string(output))
	}

	if exit := GetExitStatus(); exit {
		log.Fatalf("命令端获取到退出状态,命令:%v结束后退出", c.String())
	}
	return nil
}

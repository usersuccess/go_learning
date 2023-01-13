package main

import (
	"context"
	"os/exec"
)

func main() {

	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result

	)

	//创建一个结果队列

	ctx,cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"/bin/bash","-c","sleep 2;echo hellp;")

		//执行任务，捕获输出
		output,err = cmd.CombinedOutput()

	}()
}


/*
 * Copyright (c) 2018, 奶爸<1@5.nu>
 * All rights reserved.
 */

package gocd

import "io"

//Running 运行中的任务
type Running struct {
	Finish     chan bool
	Reader     *io.PipeReader
	Log        *PipeLog
	RunningLog []string
}

//RunningLogs 系统中在运行的任务
var RunningLogs map[uint]*Running

func init() {
	RunningLogs = make(map[uint]*Running)
}

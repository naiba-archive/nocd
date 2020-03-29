/*
 * Copyright (c) 2017 - 2020, 奶爸<hi@nai.ba>
 * All rights reserved.
 */

package nocd

//Running 运行中的任务
type Running struct {
	Finish     chan bool
	Closed     bool
	Log        *PipeLog
	RunningLog []string
}

//RunningLogs 系统中在运行的任务
var RunningLogs map[uint]*Running

func init() {
	RunningLogs = make(map[uint]*Running)
}

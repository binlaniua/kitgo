package kitgo

import (
	"os/signal"
	"syscall"
	"os"
	"os/exec"
	"path/filepath"
)

//-------------------------------------
//
//
//
//-------------------------------------
func WaitExitSignal() os.Signal {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	return <-signalChan
}

//-------------------------------------
//
//
//
//-------------------------------------
func RuntimePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

//-------------------------------------
//
//
//
//-------------------------------------
func ExceptionCatch() interface{} {
	err := recover()
	if err != nil {
		Log("尝试恢复, 出错原因 => ", err)
	}
	return err
}
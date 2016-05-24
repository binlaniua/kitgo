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

//-------------------------------------
//
// 使用守护进程启动
//
//-------------------------------------
func RunDaemon() (error) {
	if os.Getppid() != 1 {
		return Restart()
	}
	return nil
}

//-------------------------------------
//
// 重启
//
//-------------------------------------
func Restart() error {
	fp, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(fp, os.Args[:1]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	return nil
}
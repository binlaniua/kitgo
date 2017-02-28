package file

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"sync"
	"time"
)

//-------------------------------------
//
//
//
//-------------------------------------
type FileRotateOption struct {
	DocPath string //存放路径
	Suffix  string //后缀
	Prefix  string //前缀
	Format  string //格式
	Rotate  string //按什么时间分割
}

//-------------------------------------
//
//
//
//-------------------------------------
func (f *FileRotateOption) merge() {
	if f.Format == "" {
		f.Format = "2006_01_02"
	}
	if f.Rotate == "" {
		f.Rotate = "0 1 0 * * *"
	}
	if f.Suffix == "" {
		f.Suffix = ".log"
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
type FileRotate struct {
	file   *os.File
	lock   sync.Mutex
	job    *cron.Cron
	option *FileRotateOption
}

//-------------------------------------
//
// 存放路径, 文件后缀
//
//-------------------------------------
func NewFileRotate(option *FileRotateOption) (*FileRotate, error) {
	//
	option.merge()

	//
	lr := &FileRotate{
		lock:   sync.Mutex{},
		option: option,
		job:    cron.New(),
	}

	//
	if err := lr.reopen(); err != nil {
		return nil, err
	}

	//按天
	err := lr.job.AddFunc(option.Rotate, func() {
		log.Print("已触发....新建文件做log")
		err := lr.reopen()
		if err != nil {
			log.Fatal("重新打开文件失败 => ", err)
		}
	})
	if err != nil {
		return nil, err
	}
	lr.job.Start()
	return lr, nil

}

//-------------------------------------
//
// 新建文件
//
//-------------------------------------
func (lr *FileRotate) reopen() (err error) {
	lr.lock.Lock()
	defer lr.lock.Unlock()
	//关闭文件
	lr.file.Close()

	//新建文件
	o := lr.option
	fileName := time.Now().Format(o.Format)
	filePath := fmt.Sprintf("%s/%s%s%s", o.DocPath, o.Prefix, fileName, o.Suffix)
	lr.file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return
}

//-------------------------------------
//
// 日志写入
//
//-------------------------------------
func (lr *FileRotate) Write(b []byte) (int, error) {
	lr.lock.Lock()
	defer lr.lock.Unlock()
	return lr.file.Write(b)
}

//-------------------------------------
//
// 日志关闭
//
//-------------------------------------
func (lr *FileRotate) Close() error {
	lr.lock.Lock()
	defer lr.lock.Unlock()
	defer lr.job.Stop()
	return lr.file.Close()
}

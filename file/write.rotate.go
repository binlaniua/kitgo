package file

import (
	"sync"
	"os"
	"github.com/robfig/cron"
	"github.com/labstack/gommon/log"
	"fmt"
	"time"
)

//-------------------------------------
//
//
//
//-------------------------------------
type FileRotateOption struct {
	docPath string //存放路径
	suffix  string //后缀
	prefix  string //前缀
	format  string //格式
	rotate  string //按什么时间分割
}

//-------------------------------------
//
//
//
//-------------------------------------
func (f *FileRotateOption) merge() {
	if f.format == "" {
		f.format = "2006_01_02"
	}
	if f.rotate == "" {
		f.rotate = "0 1 0 * * *"
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
	option.merge();

	//
	lr := &FileRotate{
		lock:     sync.Mutex{},
		option: option,
		job: cron.New(),
	}

	//
	if err := lr.reopen(); err != nil {
		return nil, err
	}

	//按天
	err := lr.job.AddFunc(option.rotate, func() {
		log.Print("已触发....新建文件做log")
		err := lr.reopen();
		if err != nil {
			log.Error("重新打开文件失败 => ", err)
		}
	});
	if err != nil {
		return nil, err
	}
	lr.job.Start();
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
	fileName := time.Now().Format(o.format)
	filePath := fmt.Sprintf("%s/%s%s%s", o.docPath, o.prefix, fileName, o.suffix)
	lr.file, err = os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
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
	defer lr.job.Stop();
	return lr.file.Close()
}
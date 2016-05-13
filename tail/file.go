package tail

import (
	"github.com/binlaniua/kitgo"
	"github.com/howeyc/fsnotify"
)

//-------------------------------------
//
//
//
//-------------------------------------
type TailListener interface {
	OnChange(filePath string)
}

//-------------------------------------
//
//
//
//-------------------------------------
func TailFile(filePath string, lis TailListener) *fsnotify.Watcher {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		kitgo.Log(filePath, " 监听失败 => ", err)
		return nil
	}
	err = watch.Watch(filePath)
	if err != nil {
		kitgo.Log(filePath, " 监听添加文件失败 => ", err)
		return nil
	}
	go func() {
		for {
			select {
			case evt := <-watch.Event:
				if evt.IsModify() {
					lis.OnChange(evt.Name)
				}
			case err := <-watch.Error:
				kitgo.Log(err)
			}
		}
	}()
	return watch
}

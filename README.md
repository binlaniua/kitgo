# 依赖
```
go get github.com/tealeg/xlsx
go get github.com/extrame/xls
```
## convert 转换
```
FromUnicode(src)    转换unicode到中文
ToInt(src)          转换到数字
```

## config 配置
```
Log(...)            获取开发环境打日志
```

## excel xls | xlsx 操作
```
ReadCell(filePath, sheetIndex, rowIndex, cellIndex)     读取cell内容
```

## runtime 启动操作
```
WaitExitSignal()    等待退出信号
```

## file 文件
```
ReadString(arg)     读取整个文件
ReadBytes(arg)      读取整个文件
ReadLines(arg)      按行读取文件
WriteString(arg)    写入整个文件
```

## http http请求
```
HttpGet(url)            
HttpGetReply(url, time)
HttpPostFrom(url, map)
HttpPostJson(url, data)
HttpPostFile(url, data)
```

## mail 发邮件
```
```

## security 加解密
```
Guid()
Md5(s)
MD5WithSalt(s, salt)
```

## string 字符串操作
```
FirstLetterLower(src)       首字母小写
FirstLetterUpper(src)       首字母大写
Trim(src)                   去首尾空格
Between(src, start, end)    取start, end中间的
StartWith(src, s)           是否s开头
After(src, s)               从s开始截取
Before(src, s)              截取到s的位置
LeftPad(src, length, pad)   填充
```

# 依赖
```
go get github.com/axgle/mahonia
```

## 
```
kitgo.ConvertGBToUTF(string) 			Gb格式转换为UTF-8
kitgo.ConvertUnicode(string)    		转换unicode到中文
kitgo.ConvertToInt(string)
kitgo.ConvertToFloat(string)			
kitgo.ToFixed(float32, int)				保留几位小数


kitgo.MapGetString(map[string]interface{}, string)
kitgo.MapGetInt64(map[string]interface{}, string)


kitgo.WaitExitSignal()					等待退出信号
kitgo.RuntimePath()						运行目录
kitgo.ExceptionCatch()					捕获异常
kitgo.RunDaemon()						守护进程启动
kitgo.Restart()							重启

```

## string 字符串操作
```
kitgo.StringFirstLetterLower(src)       首字母小写
kitgo.StringFirstLetterUpper(src)       首字母大写
kitgo.StringTrim(src)                   去首尾空格
kitgo.StringBetween(src, start, end)    取start, end中间的
kitgo.StringStartWith(src, s)           是否s开头
kitgo.StringAfter(src, s)               从s开始截取
kitgo.StringBefore(src, s)              截取到s的位置
kitgo.StringMatch(src, reg, group)		正则捕获组
kitgo.StringReplace(src, reg, s)		正则替换
kitgo.StringLeftPad(src, length, pad)   填充
StringSplitByRegexp(src, reg)			按正则分隔字符串

```

## zip
```
zip.ZipDir(src, filePath)				压缩目录
```

## lib  第三方扩展
```
lib.RedisSub(redisConn, channel, handle) 监听
lib.RedisLRangeStrings(redisConn, key, start, end) LRange
```

## security
```
security.AESEncrypt
security.AESDecrypt

security.ToBase64
security.FromBase64

security.DesEncrypt
security.DesDecrypt

security.MD5
security.MD5WithSalt
security.MD5Map

security.SHA1
security.SHA1Map
```

## file
```
file.ReadBytes(filePath)
file.ReadString(filePath)
file.ReadLines(filePath)
file.LoadJsonFile(filePath, interface{})
file.ListFilePaths(dirPath)
file.WriteString(filePath, string)
file.WriteBytes(filePath, []byte)
```

## db (只有mysql)
```
db.Connect()
db.ConnectAsAlias()
db.GetDB()
db.GetDBByAlias()
db.Insert()
db.InsertByAlias()
db.HasData()
db.HasDataByAlias()
db.QueryMaps()
db.QueryMapsByAlias()
db.QueryMap()
db.QueryMapByAlias()
db.Query()
db.QueryByAlias()
db.Update()
db.UpdateByAlias()
```

## http (支持 http/https/socks5 代理)
```
http.HttpGet()
http.HttpPostFrom()
http.HttpPost()
http.HttpPostJson()
http.HttpPostFile()
```

## net
```
kitgo.TcpConnection
```
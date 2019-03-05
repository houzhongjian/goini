# goini

> goini是用golang开发的一个超级轻量级的ini库

### 安装方式
> go get github.com/houzhongjian/goini

### 使用方式
```go
package main

import (
    "github.com/houzhongjian/goini"
    "log"
)

func main() {
    //初始化.
    if err := goini.Load("app.conf"); err != nil {
        log.Printf("%+v\n", err)
        return
    }

    goini.GetString("db_host")
    goini.GetInt("db_port")
    goini.GetBool("isUpload")
    ...
}
```

### 配置文件示例
```
# mysql 配置
db_name = blog
db_host = 127.0.0.1
db_pwd
db_port = 3306

#wechat
appid =sdfsffqr545242


#common
isUpload=true
isShare

#
page_size = 50

```

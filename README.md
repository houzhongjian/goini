# goini

> 当前项目仅支持golang来读取一定格式的ini配置文件

### 安装方式
> go get github.com/houzhongjian/goini

### 使用方式
```
package main

import (
    "github.com/houzhongjian/goini"
    "log"
)

func main() {
    //初始化.
    if err := goini.Init("app.conf"); err != nil {
        log.Printf("%+v\n", err)
        return
    }

    goini.Get("db_host")
    goini.Get("db_user")
    goini.Get("db_passwpord")
    ...
}
```

### 配置文件实例
```
# mysql 配置
db_name = "blog"
db_host = "127.0.0.1"
db_port = 3306

#wechat
appid ="sdfsffqr545242"


#common
isUpload=true

#
page_size = 50

```

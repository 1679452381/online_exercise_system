# 简易在线判题系统
技术栈 Gin Gorm

## 安装mysql
```shell
docker pull mysql:5.7

#启动mysql
docker run --name mysql  -e MYSQL_ROOT_PASSWORD=123456  -p 3306:3306 -d mysql:5.7 
```


## 依赖
```shell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/gin-contrib/cors
go get github.com/redis/go-redis/v9
go get -u github.com/golang-jwt/jwt/v5
go get github.com/google/uuid
go get github.com/jordan-wright/email 

# gin直接连接mysql
go get -u github.com/go-sql-driver/mysql
```
## swagger 
https://gitcode.net/mirrors/swaggo/gin-swagger 
```shell
swag init 

# Fetch errorInternal Server Error doc.json 的时候。是因为没有导入依赖包doc的问题
import _ "项目名/docs"
```


## 功能
* 用户管理
  * 注册
    * * 发送邮箱验证码
  * 登录
  * 用户详情
  * 用户提交排行
  * 用户提交列表
* 问题管理
  * 添加问题
  * 修改问题
  * 获取问题列表
  * 获取问题详情
* 分类管理
  * 添加分类
  * 删除分类
  * 修改分类
  * 获取分类列表
* 代码判断


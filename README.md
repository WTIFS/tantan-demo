# tantan-demo

## Project Description 项目描述
- Tantan Back-End Developer Test 
- The task is to implement a basic RESTful HTTP service in Go for a simplified Tantan backend:
                                   Adding users and swiping other people in order to find a match. 

## Project Structure 项目结构
- config 配置
  - pg PostgreSQL配置
- constants 常量
  - relationship
- controller 路由层
  - common 路由层的公共方法
  - userController 
- dao 数据库交互层
  - baseDao dao层的公共方法
  - relationDao
  - userDao
- model 模型类
  - relationship
  - user
- service 逻辑层
  - relationshipService
  - userService
- test 测试
- vendor 项目依赖库
- main 程序入口

## Environment Setup 运行环境设置
1. Make sure you install go, PostgreSQL and govendor
2. Edit config/pg.go to connect your PostgetSQL
3. Copy this project to $GOPATH/src/
4. Run the following command to fetch dependencies: 
    ```bash
        cd vendor
        govendor sync
    ```
5. Run the following comman to start the project
    ```bash
        go run main.go
    ```
7. Then you can test the Apis
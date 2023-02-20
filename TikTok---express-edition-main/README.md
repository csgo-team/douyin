# TikTok---express-edition
简易版抖音 123

项目采取MVC架构

main函数启动后，会注册router

然后请求路径对应的router先会经过middlerware（如果有的话）的函数，再进入controller的函数

controller主要调用service层的服务来获取response，然后返回response即可

service调用model文件夹下的DAO文件的增删改查的功能，然后在本层作逻辑处理之后，返回相应的response给controller

model下的_dao文件，直接对数据库操作，（如果是userdao）获取对应的user，并返回数据




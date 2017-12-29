# cqut-cli
终端获取重庆理工大学学生个人信息

## 安装

`` go get github.com/Happy-Friday/cqut-cli``

Or

[下载release版](https://github.com/Happy-Friday/cqut-cli/releases)

## 使用

```bash
>> cqut-cli
Usage of ./cqut-cli:
  -f string
        文件名
  -l    是否打印log
  -p string
        登陆密码
  -qu string
        要查询的用户名
  -t string
        学期
  -tp string

        查询类型:
                grades 成绩
                ctable 课表
                userinfo 学生信息
                gpoint 绩点和学分总和
                photo 照片
         (default "help")
  -u string
        登陆账号
  -y string
        学年
```

```bash
./cqut-cli -u 115xxxxxxxx -p xxxxxxx -tp ctable
./cqut-cli -u 115xxxxxxxx -p xxxxxxx -tp photo -qu 115xxxxxxx -f picture.gif
```


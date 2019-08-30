# git部分

## markdown文件
文件后缀名为.md的文件，就是markdown格式文件，在线markdown编辑器就是 [https://maxiang.io](https://maxiang.io)

## 终端介绍

Terminal终端：git bash，进入某个目录，鼠标右键，出现git bash菜单，用来代替cmd

### 三大操作系统：
1. windows微软
2. mac osx苹果，底层是unix
3. linux开源，借鉴unix产生的，服务器上常用（后端都是用linux系统开发，所以终端也是linux风格的）
4. unix闭源

mac系统和linux系统，都称为`类unix`系统

windows和类unix最大的不同是在文件系统方面，格式化的原理就是在写入文件系统到硬盘（磁盘）中
windows的文件系统是分为CD...盘，类unix只有一个根目录/，倒着的树状结构

### 当前用户目录

windows的当前用户目录是在：C:\Users\xxx（用户目录下）

### 终端常用命令

cd命令，在终端中，进入一个文件夹时用到的命令

`cd xxx`：进入某个目录中
`cd或者cd ~`：进入当前用户目录，别名：~

windows下切换盘目录：cmd下，直接输入D:，回车，git-bash下，直接输入`cd /d`

ls命令是查看当前目录下存在的所有文件及目录，单纯的ls命令是不会显示隐藏文件（文件名以.点开头的文件）的，可以通过加参数的形式显示：ls -al

clear命令，终端清屏

pwd命令：查看当前所在目录的路径

touch命令：创建一个空的文件，例如：touch abc.txt

mkdir命令：创建文件夹：例如：mkdir go-training，创建多级目录：mkdir -p go-training/test

rm命令：删除文件，例如：rm abc.txt，删除文件夹：rm -rf go-training（慎用）

cat命令：查看文本文件的内容

echo命令：打印输出内容的，也可以写内容到文件中，例如：echo "# go-training" >> README.md（文件不存在，则创建文件，文件存在，则追加内容）

在终端中，很多情况都是一不小心按进了另一个操作环境中，退出：ctrl+c，或者q

终端中，默认（常用）编辑器，vim（另一个linux环境中常用的编辑器是emacs），退出esc + :q（:q!）

### git的基本使用
git
1. 首先在github创建一个仓库，repositories下面，New新建
2. 在自己电脑上，创建一个文件夹，用来对应github上的代码仓库
3. 进入电脑上的文件夹，用`git init`命令初始化文件夹，此时文件夹为受git管理的文件夹，文件夹下多出`.git`隐藏目录
4. 在当前文件夹内，配置使用git工具的用户名和email，`git config user.name "xxx"`，`git config user.email "xxx"`，
   此时`.git/config`文件就会多出两行内容（--global参数指定全局信息）
5. 在目录中创建README.md文件（其他文件），写入一点内容：`echo "# go-training" >> README.md`
6. `git status`，查看当前git的状态

```bash
$ git status                 # git status命令，查看当前被git管理的文件夹目录下的，文件变动情况
On branch master             # 默认处于master分支

No commits yet               

Untracked files:             # 未被追踪的文件
  (use "git add <file>..." to include in what will be committed)

        README.md

nothing added to commit but untracked files present (use "git add" to track)
```

7. 通过`git add`命令，将文件加入到git管理中，例如：git add README.md（指定具体的文件），git add .，git add -A（两者都表示添加所有修改的文件）
8. 通过`git commit -m "XXX"`命令，将本次对文件夹中所做的修改，进行提交，-m参数指定提交的备注
9. 绑定远程的仓库地址：`git remote add addr_name addr_url`，remote其实类似于一张表，下面可以填写很多个远程链接（addr_url），一个远程链接，对应一个名字（addr_name）
10. `git remote -v`查看已经添加的远程链接
11. 使用`git push -u addr_name branch`，将本地文件推送（push）到addr_name地址下的branch分支，例如：`git push -u origin master`

推送成功提示：
https://github.com/Manson36/project01.git

```bash
$ git push -u origin master
Username for 'https://github.com': Manson36
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Writing objects: 100% (3/3), 228 bytes | 228.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To https://github.com/Manson36/go-training.git
 * [new branch]      master -> master
Branch 'master' set up to track remote branch 'master' from 'origin'.
```

简化使用git推送代码：

1. 新建文件夹，git init初始化目录
2. git config，配置用户名和邮箱
3. 写代码，或者编辑文件
4. git add -A，添加所有内容到git中
5. git commit -m "xxx"，提交并备注
6. git push -u origin master，推送文件变动信息到远程github仓库中

复制一个github远程项目到本地：git clone addr_url，例如：git clone https://github.com/Manson36/go-training.git

mkdir go-training
cd go-training
git init
git config user.name "Manson36"
git config user.email "xxx"
echo "# go-training" >>README.md
git status
git add -A
git commit -m "[change] change some"
git remote add ;'origin https://github.com/Manson36/go-training.git
git remote -v
git push -u origin master

## go包管理器

### govendor

在go 1.11版本之前，govendor是最好用的包管理器，没有之一

go get是go语言自带的包管理器，默认会把包下载到$GOPATH/src

go get -u：-u参数表示下载并更新包

go1.11版本之前，除了govendor以外，其他的包管理工具的缺点：

1. go语言的包很多都被墙了，下载速度慢，当把项目移交给其他人时，其他人需要能够翻墙下载
2. 版本问题，比如说：本次使用的redis包的版本是1.0，测试无问题，开发其他项目时，又下载更新了一次redis包，这次包可能存在未检测出的bug。

govendor借助了go语言vendor机制，消除了上述缺点：
1. govendor会把项目依赖的包下载到当前目录下
2. govendor会通过vendor.json文件维护依赖包的版本

govendor使用：
1. 如果当前项目所在的目录并没有在GOPATH下，可在git bash中临时的设置一下GOPATH：`export GOPATH=C:\\Users\\Manson\\Desktop\\go`
1. 下载govendor：`go get -u github.com/kardianos/govendor`
2. 进入到项目目录，执行`govendor.exe init`
3. 下载包，执行`govendor.exe fetch github.com/go-redis/redis`

搜索开源的第三方包网站：https://godoc.org/

### go module

go语言1.11版本之后支持的包管理工具，已经被设定为标准的包管理工具。

- 第一步，检查go语言的版本信息，是否是1.11以上，git bash中输入`go version`：`go version go1.11.5 windows/amd64`
- 第二步，开启GO111MODULE环境变量，激活go module功能：`export GO111MODULE=on`


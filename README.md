## trpc_demo
https://github.com/trpc-group/trpc-go

## 环境安装
安装 官方命令行工具 文档
通过
go install trpc.group/trpc-go/trpc-cmdline/trpc@latest 
 安装 trpc 命令行工具
使用 trpc setup 一键安装所有依赖 可能会因为一些网络问题导致设置失败，可以通过如下方式手动配置
- 下载官方提供的二进制文件，如 macOS 系统的 https://github.com/trpc-group/trpc-cmdline/releases/tag/v0.0.1-darwin
- 添加可执行权限 
chmod +x 工具二进制文件
- 将这些工具拷贝到 
go/bin  
which 工具二进制文件 

## 自己项目：
https://github.com/DoubleZ0405/trpc_demo
go mod init github.com/xx/xx

touch tdm.proto

生成项目框架代码：
trpc create -p tdm.proto -o . 

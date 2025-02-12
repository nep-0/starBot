# Star Bot
Star Bot 是基于 OneBot 11 API 开发的回复机器人。

开发时间仓促，而且本来没有想开源，因此代码看起来乱糟糟的。欢迎想优化的同志们提 PR。

### 编译部署
0. (可选) 通过我的邀请码 `LyOdlqTM` 注册 [硅基流动](https://cloud.siliconflow.cn/i/LyOdlqTM) 领取 2000万 tokens，还能体验满血版 DeepSeek-R1 [点击注册](https://cloud.siliconflow.cn/i/LyOdlqTM)
1. 安装并启动支持 OneBot 11 API 的 bot 框架（如 [NapCat](https://napneko.github.io/) ）
2. (可选，向量搜索使用) 在 [Zilliz](https://cloud.zilliz.com/) 数据库插入要检索的数据
3. 安装 Go (>=1.24.0) [官网](https://golang.google.cn/doc/install)
4. (可选，推荐在中国大陆使用) 设置 Go module proxy 如 [Goproxy.cn](https://goproxy.cn/)
5. 克隆仓库，编译项目
```sh
git clone https://github.com/nep-0/starBot.git
cd starBot
go build
```
6. 修改 `config.yaml`，与 1, 2 步的配置对应
7. 启动！
```sh
./starBot
```

### 加个功能？
这很简单。看看 `llm/r1.go` 是怎么写的，依葫芦画瓢写一个出来，再在 `main.go` 里调用就行了。

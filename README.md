这是一个 Web 框架性能评测实验项目，
项目中实现了一个性能评测工具，
和三个非常简单的用于评测的服务。

服务分别使用了:
- go 的原生 net/http 包
- rust 的 actix-web 框架（基于tokio）
- nginx 1.23

由于压力评测指标与服务部署的基础设施关联性非常大，
这里也不再给出评测结果和对比结论，
研究者可以根据自身情况自行评测。

关于客户端
----

由于压测工具需要使用本地 socket 资源发送请求，
需要调整本地 TCP 链接回收策略，
以支持大量并下情况下依然有足够的本地 socket 可用：

MacOS：

```
sudo sysctl net.inet.tcp.msl=1000
```

Linux

```
sudo sysctl net.ipv4.tcp_tw_reuse=1
sudo sysctl net.ipv4.tcp_tw_recycle=1
```

关于服务端
----

为了方便研究者重现实验过程，
所有服务都实现了通过 Docker 部署的脚本。

下载项目后直接执行 make 命令即可看到结果。

对于高级研究者来讲，也可以根据自身诉求调整各个服务参数，
已获得想要的评测结果。

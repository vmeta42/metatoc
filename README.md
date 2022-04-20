# MetaTOC
**注意**：开发过程中，master分支可能处于不稳定的甚至中断的状态。 请使用releases分支，来获得稳定的项目文件。

**MetaTOC** 是基于边缘计算的数据采集和模型下的开源项目，是以区块链Violas做为数据去中心化鉴权和追踪的最佳实践。
### 特点
- 采用CRD自定义资源，简化对象传输的定义
- 利用Kubeedge原生的高可用能力，保证服务的可靠性
- 数据收集和AI模型下发使用同样的服务，简化边缘计算的实现
- 支持边缘侧数据采集，链原生确认数据所有权
- 支持数据所有权转让上链
- 支持数据转让追溯
- 支持边缘智能模型下发
### 架构
1、边缘端的数据采集
![采集端](vx_images/223624814226753.jpg =538x)

2、Violas合约实现流程
![（原图）智能合约流程图](vx_images/565104814246919.png =538x)
### 安装部署
####   CEFCO 
1、[部署文档](https://github.com/vmeta42/metatoc/tree/main/cefco#readme)
#### NATS-Kafka 
1、[部署文档](https://github.com/vmeta42/metatoc/blob/main/nats-kafka/docs/buildandrun.md)
2、[配置文档](https://github.com/vmeta42/metatoc/blob/main/nats-kafka/docs/config.md)
3、[监控文档](https://github.com/vmeta42/metatoc/blob/main/nats-kafka/docs/monitoring.md)
#### move-contracts
1、[智能合约源码](https://github.com/vmeta42/metatoc/tree/main/move-contracts)
2、编译
Ubuntu 20.04 环境，已安装 Move 语言编译器 move, 执行如下命令
```
move compile meta42.move scripts.move --mode diem
```
会在 build 目录下生成编译后的 meta42.mv 和多个 meta42*.mv脚本
### 文档列表
1、功能介绍
2、接口文档

### Roadmap
|      时间      |         节点事件         |
| :-----------: | :---------------------: |
| 2019年6月18日  |       Violas开源        |
| 2020年10月11日 | ceco/cefco写下第一行代码 |
| 2022年4月21日  |       发布MetaTOC       |
|    2022年Q3    |     支持链上数据确权      |
|    2022年Q4    |     支持链上数据交易      |
### 贡献者
 感谢小伙伴们的积极贡献，点击[贡献者](https://github.com/vmeta42/metatoc/graphs/contributors)查看详情。
我们非常欢迎更多的贡献者参与共建Meta42项目, 不论是代码、文档，或是其他能够帮助到社区的贡献形式。
### 交流
如果您在使用过程中遇到任何问题，或是有建设性意见，欢迎给我们提[issues](https://github.com/vmeta42/metatoc/issues)。
社群交流建设中，敬请期待。
### License
MetaTOC is under the  MIT license. See the License file for [details](https://github.com/vmeta42/metadb/blob/3.9.39.x/LICENSE.txt)。



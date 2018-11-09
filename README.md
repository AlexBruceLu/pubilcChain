# 比特币公链

>比特币网络是一个分布式的点对点网络， 网络中的矿工通过“挖矿”来完成对交易记录的记账过
>程， 维护网络的正常运行。比特币通过区块链网络提供一个公共可见的记账本， 用来记录发
>生过的交易的历史信息。每个交易包括一些输入和一些输出， 未经使用的交易的输出（ Unspent 
>Transaction Outputs， UTXO） 可以被新的交易引用作为合法的输入。

比特币账户采用了非对称的加密算法， 用户自己保留私钥， 对他发出的交易进行签名确认，
并公开公钥。比特币的账户地址其实就是用户公钥经过一系列 hash（ HASH160， 或先进行 
SHA256， 然后进行 RIPEMD160） 及编码运算后生成的 160 位（ 20 字节） 的字符串。也
常常对账户地址串进行 Base58Check 编码， 并添加前导字节（ 表明支持哪种脚本）和 4 
字节校验字节， 以提高可读性和准确性。

一个区块将包括如下内容：
- 4 字节的区块大小信息；
- 80 字节的区块头信息：
    + 版本号： 4 字节；
    + 上一个区块头的 SHA256 hash 值： 链接到一个合法的块上， 32 字节；
    + 包含的所有验证过的交易的 Merkle 树根的哈希值， 32 字节；
    + 时间戳： 4 字节；
    + 难度指标： 4 字节；
    + Nonce： 4 字节， PoW 问题的答案；
- 交易个数计数器；
- 所有交易的内容。

## 定义区块的结构

- 时间戳
- 前区块的哈希
- 交易数据
- 当前区块的哈希
- nonce难度值

## 创建新的区块链

主要是创建一条带有创世区块的区块链。

## 工作量证明的引入


>比特币的这种基于算力的共识机制被称为 Proof of Work（ PoW） 。 目前， 要让 hash 结果满
>足一定条件并无已知的启发式算法， 只能进行暴力尝试。 尝试的次数越多， 算出来的概率越
>大。 通过调节对 hash 结果的限制， 比特币网络控制约 10 分钟平均算出来一个合法区块。 算
>出来的节点将得到区块中所有交易的管理费和协议固定发放的奖励费（ 目前是 12.5 比特币，
>每四年减半） 。 也即俗称的挖矿。

## 添加Bolt数据库
- 创建数据库、仓库
- 插入数据到数据库
- 查询数据

## 添加命令行参数
- 挖矿
- 查询区块链信息
- help 命令行参数列表
- 


|定位|功能|智能合约|一致性|权限|类型|性能|代表|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|公信的数字货币|记账功能|不带有或功能较弱|PoW|无|公有链|较低|比特币|
|公信的交易处理|智能合约|图灵完备|PoW、PoS|无|公有链|受限|以太坊|
|带权限的交易处理|商业处理|多种语言，图灵完备|多种，可插拔|支持|联盟链|可扩展|Hyperledger|


## 基本原理
- 区块链的基本原理理解起来并不难。 基本概念包括：
    + 交易： 对账本状态的改变， 如添加一条记录；
    + 区块： 记录一段时间内发生的交易和状态， 是对当前账本状态的一次共识；
    + 链： 由一个个区块按照发生顺序串联而成， 是状态变化的日志记录。
>如果把区块链作为一个状态机， 则每次交易就是试图改变一次状态， 每次生成区块就是参与
>者对于其中包括的所有交易改变状态的结果确认。


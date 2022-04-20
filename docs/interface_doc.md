# 接口文档
### 1、智能合约接口说明
使用 Violas SDK 和编译出的 meta42*.mv 脚本字节码，即可以调用如下的接口，详细的链上交易提交方法请参考官方 SDK 文档。
| 接口名称    |  功能   | 参数说明  | 权限说明 |
| :-----: | :---------: | :---------------------------------------------------------------------------------------------------------------------------------------: | :------------: |
|  initialize   |  初始化 Meta42 合约   |无 | 调用权限：Meta42 管理员账户   |
|  accept   | 注册用户信息    | 1. 一个账户在必须调用 accept 接口之后，才能接收别人的分享 Token, 只需调用一次。2. mint_token 会在内部调用 accept, 如果用户没有调用 accept。|任何 Meta42 普通账户 |
|  mint_token   |   铸造一个 token 到自己的账户 |    hdfs_path – 链外的 HDFS 文件的路径，将会保存在 Token 中。 |任何 Meta42 用户 |
|  share_token_by_index  |  分享 token 给另一个账户  |  receiver – 接收者账户地址index – token 在当前账户下的索引，根据客户端能够获取所有的 token, 索引从 0 开始到 n-1 ；Message – 附加的信息； | 任何 Meta42 账户|
| share_token_by_token |分享 token 给另一个账户 |receiver – 接收者账户地址；token_id – token 的唯一标识；Message – 附加的信息；| 任何 Meta42 账户|
|share_token_by_token|分享 token 给另一个账户|receiver – 接收者账户地址；token_id – token 的唯一标识；Message – 附加的信息|任何 Meta42 账户|
|create_child_vasp_account|创建子账户|参见 Violas SDK 接口说明|
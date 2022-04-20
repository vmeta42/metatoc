### 需求概述
Meta42 项目需要将链外数据资产的访问权限记录到区块链上。利用 Violas 区块链的不可篡改、链上交易可追溯性，能够对权限数据(Token)进行创建、分享、溯源，同时根据链上权限数据（Token）的所有权在链外对访问数字资产的用户进行鉴权后授权访问。
### 需求分析
 在会议讨论的原始图例中， HDFS 中有三个数据块，VNET 在链上创建三个对应的

Token1、Token2 和 Token3, 用户 A 创建另外一个 HDFS4 同时在链上创建了 Token4.

期待实现的链上功能是

1. VNET 将 Token1 分享给用户 A

2. 用户 A 将 Token1 分享给用户 B

3. 用户 A 将 Token4 分享给用户 B

在当前状态下，需要通过 Violas 区块链判断某个 Token 的拥有权，还需要对 Token 的分享历史能够溯源。

下面将会议的原始需求做了一个初步的需求分析。
###  术语说明
 Token

链上的一个 Token 代表链外的一个数字资产的详细描述，可有很多子项例如 http url, 
HDSF 路

径等...

Token {

hdfs_path : string,

xxx

}， Token 中可以保存多个 xxx 子项，目前智能合约中 Token 只有 hdfs_path，保存链外的

HDFS 文件路径。

• Token 所有权

Token 存储在哪一个账户下面，就代表这个账户对 Token 的所有权。同时，这个账户根据 
Token

中子项的描述信息有权限访问链外的数字资产。

• Token Id

Token Id 代表的 Token 的唯一标识, 具有唯一性。

Token Id 的计算方式：

1. BCS 算法(Violas SDK 提供）序列化 Token 中的所有数据项，得到字节流

2. 对步骤 1 生成的字节流使用 sha3-256 哈希算法计算 hash ，产生 32 个字节。

• Meta42 账户

由 Meta42 智能合约管理员调用 create_child_vasp_account 创建的子账户，才能够调用 Meta42

合约的接口。

Violas 链上的其它账户无权限调用 Meta42 智能合约接口。
###  功能需求列表

|功能需求  |说明  |
| --- | --- |
|创建Token  |在 Violas 链上创建一个 Token（链外资产的描述结构）, 此 Token 的所有权为当前创建账户。  |
|分享 Token  | 将某个 Token 的分享(拷贝)给另外一个账户，Token 的能够被多个账户拥有。 |
|溯源Token  | 根据一个 Token Id, 追溯对应的 Token 被分享的所有历史信息 |
|链外鉴权  | 当一个用户去访问链外的数据资产时，链外后台服务需要能够判断当前的用户的链上身份，以及是否有权限访问链外的数据资产。账户身份和链外资产对应 Token 的拥有权在 Violas 链上都可以查询。 |
|创建子账户  | Violas 链原生提供创建子账户的功能 |
###  用于验证的测试用例
####  1、测试步骤
1. Violas 管理员创建 VNET 账户， VNET 属于 Parent VASP 角色。
2. VNET 账户创建子账户 Alice 和 Bob， Alice 和 Bob 属于 Child VASP 角色。
3. VNET 账户铸造(mint) Token 1, Token 2, Token 3
4. VNET 账户分享(share) Token 1 给账户 Alice
5. Alice 铸造(mint) Token 4
6. Alice 分享(share) Token1 和 Token 4 给账户 Bob
7. 测试步骤的状态图如下
![9de3a744cfb1037144a3a0b418f32134](权限数据链上存储、溯源的Violas 区块链解决方案.resources/CD5D8551-83CD-4F76-A404-C0518F4A5E7C.png)
#### 2. 溯源 Token 的期待结果
 期待的结果如下

· 对 Alice 的 Token 1 溯源： VNET→ Alice

· 对 Bob 的 Token 1 溯源： VNET → Alice → Bob

· 对 Bob 的 Token 4 溯源 ： Alice → Bob















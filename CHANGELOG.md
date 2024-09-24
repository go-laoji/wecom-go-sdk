# 变更记录

## v2.1.1 2024-09-24

* fix pull [P14](https://github.com/go-laoji/wecom-go-sdk/pull/14) 以前缀删除缓存的key，避免服务商模式删除suite_ticket
* 感谢 [yiGmMk](https://github.com/yiGmMk)

## v2.1.0 2024-08-06

* MediaGetResponse 返回http.Header,满足文件名、文件长度、文件类型等其它需求
* feat [Issues12](https://github.com/go-laoji/wecom-go-sdk/issues/12)

因修改了返回类型，不兼容之前的`Content-Type`,使用临时素材获取接口需要注意sdk版本

## v2.0.5 2024-07-08
* fix pull [P10](https://github.com/go-laoji/wecom-go-sdk/pull/10)
* fix pull [P11](https://github.com/go-laoji/wecom-go-sdk/pull/11)
* 感谢 [yiGmMk](https://github.com/yiGmMk)

## v2.0.4 2024-06-14
* 当请求返回42001时(token过期)则清空缓存触发重试

## v2.0.3 2024-01-17
* 官方2023/11/21变更获取客户群详情接口支持返回群成员版本号
* 官方2023/12/21变更接口微信客服发送消息支持发送“获客链接消息”

## v2.0.2 2023-07-13
* feature: GetBillList 获取对外收款记录
* fix: 修改 KfSyncMsgResponse 的 HasMore 类型, 由 bool -> int
* refactor: 修改 MsgMenu 结构, 将 menu 中的 Click, View, Miniprogram 独立为 struct
* fix: 删掉 MsgLink struct 的 PicUrl, 使用 ThumbMediaId 代替
* refactor: extract item of MsgList as struct KfMessage

## v2.0.1 2023-07-03
* 增加获取临时素材功能,感谢[xzvwang](https://github.com/xzvwang)贡献

## v2.0.0 2023-06-10
使用go-resty代替原生http.Client
* 增加SetProxy(proxyUrl string) 对绕过IP限制有用
* 增加SetDebug(debug bool) 可以监控http请求
* 升级了些基础库版本
* TODO:
  * 低版本测试,开发环境为 go 1.20
  * 单测补充

## v1.5.7 2023-04-03
增加jssdk调用时的config签名函数
* GetConfigSignature
* GetAgentConfigSignature

## v1.5.6 2023-03-31
bug fix: User 序列化时某个字段空值的bug

## v1.5.5 2022-12-20
对应官方 2022/12/01 接口变更

- [x] 新增接口 支持终止尚未发送的朋友圈发表任务
- [x] 变更接口 增加allow_select参数，允许成员在待发送客户列表中重新选择发送客户
- [x] 新增接口 重新触发群发通知，提醒成员完成群发任务
- [x] 新增接口 支持取消尚未发送给客户的群发任务
- [x] 变更接口 获取客服帐号列表接口，支持查看账号管理权限
- [x] 变更接口 读取消息接口，支持指定拉取某个客服账号的消息，支持视频号小店订单、商品类型消息
- [x] 变更接口 进入会话事件，客户从视频号小店和视频号进入时，返回更详细的来源信息
- [x] 变更接口 获取客户基础信息，支持客户从视频号小店和视频号进入时，查看更详细的来源信息

## v1.5.4 2022-11-06

- [x] 接口调用许可
  - [x] 自动激活设置
      - [x] 设置企业的许可自动激活状态
      - [x] 查询企业的许可自动激活状态
- [x] 应用授权
  - [x] 获取应用的管理员列表

## v1.5.3 2022-11-02

* 客户标签ID的转换接口
* ID转换接口
  * corpid的转换
  * userid的转换
  * 转换客户external_userid
  * 转换客户群成员external_userid
  * 微信客服ID的转换


## v1.5.2 2022-10-11

* 移除成员、部门、学生、家长、家校部门修改时非必填项的验证

## v1.5.1 2022-09-13
bug fix:
* GetUserInfo3rd(code string) (resp GetUserInfo3rdResponse)
* GetUserInfoDetail3rd(userTicket string) (resp GetUserInfoDetail3rdResponse)
* GetUserInfo(corpId uint, code string) (resp GetUserInfoResponse)
* GetUserDetail(corpId uint, userTicket string) (resp GetUserDetailResponse)

接口URL修改

## v1.5.0 2022-08-26

* 新增加单元测试代码，文档见test下readme

## v1.4.12 2022-08-11
bug fix:
 * 成员id列表序列化
 * 应用授权逻辑优化

增加:
- [x] 机器人管理
  - [x] 知识库分组管理
    - [x] 添加分组
    - [x] 删除分组
    - [x] 修改分组
    - [x] 获取分组列表

## v1.4.12 2022-08-11

bug fix: MessageUpdateTemplateCard 添加 agentid [cwww3](https://github.com/cwww3)提交[PR](https://github.com/go-laoji/wecom-go-sdk/commit/1204aa2e6eddbce876d2249944a287846f3c5dbf)

增加：
- [x] 自建应用获取用户隐私信息
- [x] 获取成员ID列表

## v1.4.11 2022-07-21

bug fix: DepartmentSimpleListResponse 序列化错误

## v1.4.10 2022-07-05

增加：(仅会话存档应用的secret获取的accesstoken可调用)
- [x] 会话内容存档
  - [x] 获取会话内容存档开启成员列表
  - [x] 获取会话同意情况
  - [x] 获取会话内容存档内部群信息

## v1.4.9 2022-06-30

fix:升级badger库为v3版本,并且不再将缓存写到内存；改写文件 

使用过程中有发现v2版本有占用内存高的情况

## v1.4.8 2022-06-10

增加:
- [x] 通讯录管理
  - [x] 成员管理
      - [x] 创建成员
      - [x] 更新成员
      - [x] 删除成员
  - [x] 部门管理
      - [x] 创建部门
      - [x] 更新部门
      - [x] 删除部门

## v1.4.7 2022-05-27

fix: GetGroupMsgSendResultResponse中external_userid序列化错误

增加：新增ExecuteCorpApi方法，用来执行未实现的授权企业应用接口，返回值需要自行做序列化处理

## v1.4.6 2022-05-27

增加服务商接口调用许可相关接口
- [x] 接口调用许可
    - [x] 订单管理
        - [x] 下单购买帐号
        - [x] 下单续期帐号
        - [x] 获取订单列表
        - [x] 获取订单详情
        - [x] 获取订单中的帐号列表
    - [x] 帐号管理
        - [x] 激活帐号
        - [x] 获取激活码详情
        - [x] 获取企业的帐号列表
        - [x] 获取成员的激活详情
        - [x] 帐号继承

## v1.4.5 2022-05-24

- [x] MessageSend 增加`agentid`参数

## v1.4.4 2022-05-24

* 增加：SetAgentIdFunc、GetAgentId　定义用于获取应用的agentid
* 弃用:　CorpPermanentCode表将永久授权码信息合并存入　agent　表
* defaultAppSecretFunc默认从agent表中读取配置

## v1.4.3 2022-05-23

bugfix:

* KfAccountList 接口定义错误,感谢[Drogenwei](https://github.com/Drogenwei)反馈

## v1.4.2 2022-05-03
官方2022/04/29更新

* 客户联系
  * 变更接口 获取客户详情接口：若客户来源于视频号，则返回视频号添加场景（主页或直播间）
  * 变更接口 批量获取客户详情接口：若客户来源于视频号，则返回视频号添加场景（主页或直播间）
* 微信客服
  * 变更接口 获取客服帐号列表接口支持分页拉取
  * 变更接口 添加接待人员接口支持按部门配置接待人员
  * 变更接口 删除接待人员接口支持按部门删除接待人员
  * 变更接口 获取接待人员列表接口返回接待人员部门的id

## v1.4.1 2022-04-21
增加：
- [x] 微信客服
  - [x] 客服帐号管理
    - [x] 添加客服帐号
    - [x] 删除客服帐号
    - [x] 修改客服帐号
    - [x] 获取客服帐号列表
    - [x] 获取客服帐号连接
  - [x] 接待人员管理
    - [x] 添加接待人员
    - [x] 删除接待人员
    - [x] 获取接待人员列表
  - [x] 会话分配与消息收发
    - [x] 分配客服会话
    - [ ] 接收消息和事件
      - [x] 读取消息
    - [x] 发送消息
    - [x] 发送欢迎语等事件响应消息
  - [x] [升级服务]配置
  - [x] 其它基础信息获取
    - [x] 获取客户基础信息
    - [x] 获取企业状态信息
  - [x] 统计管理
    - [x] 获取「客户数据统计」企业汇总数据
    - [x] 获取「客户数据统计」接待人员明细数据

## v1.4.0 2022-04-11
* 兼容代开发、自建应用、三方应用
* 增加 `SetAppSecretFunc` 接口以处理自定义的应用secret配置
* 增加jsapi的两个ticket获取接口

## v1.3.0 2022-02-18
* 增加代开发应用的回调处理示例
* 修改项目名称为`wecom-go-sdk`

## v1.2.0 2022-02-07

* 增加自建代开发支持(可用于三方应用或者是自建代开发应用)

## v1.1.3 2022-02-07

* logic增加Migrate,修复不使用默认配置文件名时数据库创建失败的bug
* 服务端路由注入格式统一;修改样例中代码


## v1.1.2 2022-01-12

针对企业微信4.0的接口变更
* 新增接口 第三方应用新增组织架构信息权限，可以获取部门组织架构以及上级身份
* 新增接口 可获取指定部门的全部子部门ID列表
* 新增接口 可获取单个部门详情，包括部门负责人
* 变更接口 创建成员接口，可以指定企业邮箱biz_mail
* 变更接口 更新成员接口，可以指定企业邮箱biz_mail
* 变更接口 读取成员接口，新增返回企业邮箱biz_mail
* 变更接口 获取部门成员详情，新增返回企业邮箱biz_mail
* 新增接口 获取带参的应用二维码
* 变更接口 企业用户通过带参的应用二维码安装应用之后，获取企业永久授权码时返回state值
* 变更事件 企业用户通过带参的应用二维码安装应用之后，授权通知事件返回State字段


## v1.1.1 2022-01-01

增加：

- [x] 学校应用
    - [x] 上课直播
        - [x] 获取老师直播ID列表
        - [x] 获取直播详情
        - [x] 获取观看直播统计
        - [x] 获取未观看直播统计
        - [x] 删除直播回放
    - [x] 班级收款
        - [x] 获取学生付款结果
        - [x] 获取订单详情

## v1.1.0 2021-12-31

增加：

- [x] 学校沟通
    - [x] 基础接口
        - [x] 获取「学校通知」二维码
        - [x] 管理「学校通知」的关注模式
            - [x] 设置关注「学校通知」的模式
            - [x] 获取关注「学校通知」的模式
        - [ ] 发送「学校通知」
        - [x] 手机号转外部联系人ID
        - [x] 管理「老师可查看班级」模式
            - [x] 设置「老师可查看班级」的模式
            - [x] 获取「老师可查看班级」的模式
        - [x] 获取可使用的家长范围
    - [x] 学生与家长管理
        - [x] 创建学生
        - [x] 删除学生
        - [x] 更新学生
        - [x] 批量创建学生
        - [x] 批量删除学生
        - [x] 批量更新学生
        - [x] 创建家长
        - [x] 删除家长
        - [x] 更新家长
        - [x] 批量创建家长
        - [x] 批量删除家长
        - [x] 批量更新家长
        - [x] 读取学生或家长
        - [x] 获取部门成员详情
        - [x] 设置家校通讯录自动同步模式
        - [x] 获取部门家长详情
    - [x] 部门管理
        - [x] 创建部门
        - [x] 更新部门
        - [x] 删除部门
        - [x] 获取部门列表
        - [x] 修改自动升年级的配置

## v1.0.2 2021-12-25

增加：
- [x] 应用管理
    - [x] 获取应用
        - [x] 获取指定的应用详情
        - [x] 获取access_token对应的应用列表
- [x] 电子发票
  - [x] 查询电子发票
  - [x] 更新发票状态
  - [x] 批量更新发票状态
  - [x] 批量查询电子发票

## v1.0.1 2021-12-18

增加：
* 统计管理
* 管理商品图册
* 管理聊天敏感词

## v1.0.0 2021-12-15

api列表参考readme
此版本未做完整测试，后续将补充待完善接口及测试代码

因第三方应用限制将不能像自建应用一样做单元测试；后期思路写一个完整的web接口配合postman做接口测试

本sdk需要依赖Mysql数据库只需要在配置文件写好连接信息，运行时会自动创建数据表。

数据表结构仅当前版本测试使用后期有可能修改结构

当前版本生产环境慎用！慎用！慎用！慎用！

版本变更日志参见 [`CHANGELOG.md`](CHANGELOG.md)

[![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg "logo")](https://jb.gg/OpenSourceSupport)

`更多相关文章将会在公众号分享，欢迎关注！`

![Geek进阶](./qrcode.png)

# 文档

[使用wecom-go-sdk快速开发企业微信自建应用](https://mp.weixin.qq.com/s?__biz=Mzg5NTcyOTk1OA==&mid=2247483782&idx=1&sn=84345385511ca981c47b55f7ae24b81c)

# V2版本升级

http请求改用`go-resty`可方便指定重试、错误处理hooks(当前未处理)

## 使用

    go get github.com/go-laoji/wecom-go-sdk/v2

* 增加SetProxy(proxyUrl string) 对绕过IP限制有用,可参考[OpenResty(nginx)配置正向代理绕过企微自建IP限制](https://mp.weixin.qq.com/s/ZDUZyIaz3HOsrqRUQx8w-w)
* 增加SetDebug(debug bool) 可以监控http请求,看如下的输出


    ~~~ REQUEST ~~~
    GET  /cgi-bin/gettoken?corpid=wp0k1qEQAAiwQMWYzF8JOr99RJRX1_1A&corpsecret=Y2YbFpt95RcGTs9CvriQ2uj23Wl8B3IxtbRM9nKfjVU  HTTP/1.1
    HOST   : qyapi.weixin.qq.com
    HEADERS:
        Accept: application/json;charset=UTF-8
        Content-Type: application/json;charset=UTF-8
        User-Agent: wecom-go-sdk-v2
    BODY   :
    ***** NO CONTENT *****
    ------------------------------------------------------------------------------
    ~~~ RESPONSE ~~~
    STATUS       : 200 OK
    PROTO        : HTTP/2.0
    RECEIVED AT  : 2023-06-10T21:03:57.834584+08:00
    TIME DURATION: 181.642531ms
    HEADERS      :
        Content-Length: 176
        Content-Type: application/json; charset=UTF-8
        Date: Sat, 10 Jun 2023 13:03:58 GMT
        Error-Code: 40001
        Error-Msg: invalid credential, hint: [1686402237354962763423746], from ip: 123.116.182.237, more info at https://open.work.weixin.qq.com/devtool/query?e=40001
        Server: nginx
    BODY         :
    {
       "errcode": 40001,
       "errmsg": "invalid credential, hint: [1686402237354962763423746], from ip: 123.116.182.237, more info at https://open.work.weixin.qq.com/devtool/query?e=40001"
    }
    ==============================================================================


## 配置文件格式

    CorpId: wwxxx
    ProviderSecret: xxxxxxx
    SuiteId: xxxxx
    SuiteSecret: xxxxx-Gl8VA
    SuiteToken: xxxxx
    SuiteEncodingAesKey: xxxx
    Dsn: user:pass@tcp(127.0.0.1:3306)/suite?charset=utf8mb4&parseTime=True&loc=Local
    Port: 8080

## 　目录结构
    
    ├── config  解析yaml配置文件
    ├── example　样例程序
    ├── internal　内部函数，包含error和http请求
    └── pkg 附加包
        └── svr 服务端接口部分
            ├── callback　指令及数据回调
            ├── install 应用安装连接生成、应用安装后的跳转
            ├── logic　各事件通知处理逻辑
            ├── middleware　gin的中间件方便handler里引入sdk
            └── models  应用安装时常用到的表定义,基于gorm

## 第三方包

- [github.com/dgraph-io/badger/v3 v3.2103.5](https://github.com/dgraph-io/badger) 一个KeyValue数据库,类redis
- [github.com/gin-gonic/gin  v1.9.1](https://github.com/gin-gonic/gin) web框架
- [github.com/go-laoji/wxbizmsgcrypt v1.0.0](https://github.com/go-laoji/wxbizmsgcrypt) 微信数据加解密
- [github.com/jinzhu/copier v0.3.5](https://github.com/jinzhu/copier) jinzhu大神的数据拷贝库
- [go.uber.org/zap v1.24.0](https://go.uber.org/zap) 日志库
- [gopkg.in/natefinch/lumberjack.v2 v2.2.1](https://gopkg.in/natefinch/lumberjack.v2) 日志切割
- [gopkg.in/yaml.v2 v2.4.0](https://gopkg.in/yaml.v2) yaml解析
- [gorm.io/driver/mysql v1.5.1](https://gorm.io/driver/mysql) 同属jinzhu大神的mysql驱动
- [gorm.io/gorm v1.25.1](https://gorm.io/gorm) 同属jinzhu大神的orm库

## API列表

- ID转换接口
  - [x] corpid的转换
  - [x] userid的转换
  - [x] external_userid的转换
  - [x] 客户标签ID的转换
  - [x] 微信客服ID的转换
- 身份验证
  - 网页授权登录
    - [x] 获取访问用户身份
    - [x] 获取访问用户敏感信息
  - 扫码授权登录
    - [x] 获取登录用户信息
- 应用授权
  - [x] 调用第三方应用凭证
  - [x] 获取预授权码
  - [ ] 设置授权配置
  - [x] 获取企业永久授权码
  - [x] 获取企业授权信息
  - [x] 获取企业凭证　~~getCorpToken~~(存入数据库以主键代替)
  - [x] 获取应用的管理员列表
  - [x] 获取应用二维码
  - [x] 回调接口~~pkg下svr使用gin框架的web接口~~
    - [x] 推送suite_ticket
    - [x] 授权通知事件 *此部分与业务逻辑相关，仅做样例*
      - [x] 授权成功通知
      - [ ] 变更授权通知
      - [ ] 取消授权通知
    - [ ] 成员通知事件
    - [ ] 部门通知事件
    - [ ] 标签通知事件
    - [ ] 共享应用事件回调
    - [ ] 应用管理员变更回调
- [x] 通讯录管理
  - [x] 成员管理
    - [x] 创建成员
    - [x] 读取成员
    - [x] 更新成员
    - [x] 删除成员
    - [x] 获取部门成员
    - [x] 获取部门成员详情
    - [x] userid与openid互换
    - [ ] 二次验证
    - [ ] 邀请成员
    - [x] 手机号获取userid
    - [x] 获取成员授权列表
    - [x] 查询成员用户是否已授权
    - [x] 获取选人ticket对应的用户
    - [x] 自建应用获取用户隐私信息
    - [x] 获取成员ID列表
  - [x] 部门管理
    - [x] 创建部门
    - [x] 更新部门
    - [x] 删除部门
    - [x] 获取部门列表
    - [x] 获取子部门ID列表
    - [x] 获取单个部门详情
  - [x] 标签管理
    - [x] 创建标签
    - [x] 更新标签名字
    - [x] 删除标签
    - [x] 获取标签成员
    - [x] 增加标签成员
    - [x] 删除标签成员
    - [x] 获取标签列表
  - [ ] 异步批量接口
  - [ ] 通讯录回调通知
  - [ ] 通讯录搜索
  - [ ] 通讯录ID转译
  - [ ] 通讯录userid排序
  - [ ] 异步导出接口
  - [ ] 推广二维码
- [x] 客户联系
  - [x] 获取配置了客户联系功能的成员列表
  - [x] 客户联系「联系我」管理
    - [x] 配置客户联系「联系我」方式
    - [x] 获取企业已配置的「联系我」方式
    - [x] 获取企业已配置的「联系我」列表
    - [x] 更新企业已配置的「联系我」方式
    - [x] 删除企业已配置的「联系我」方式
    - [x] 结束临时会话
  - [x] 客户管理
    - [x] 获取客户列表
    - [x] 获取客户详情
    - [x] 批量获取客户详情
    - [x] 修改客户备注信息
    - [x] 外部联系人unionid转换
    - [x] 代开发应用external_userid转换
  - [x] 客户标签管理
    - [x] 管理企业标签
      - [x] 获取企业标签库
      - [x] 添加企业客户标签
      - [x] 编辑企业客户标签
      - [x] 删除企业客户标签
    - [x] 编辑客户企业标签
  - [x] 在职继承
    - [x] 分配在职成员的客户
    - [x] 查询客户接替状态
  - [x] 离职继承
    - [x] 获取待分配的离职成员列表
    - [x] 分配离职成员的客户
    - [x] 查询客户接替状态
    - [x] 分配离职成员的客户群
  - [x] 客户群管理
    - [x] 获取客户群列表
    - [x] 获取客户群详情
    - [x] 客户群opengid转换
  - [x] 客户朋友圈
    - [x] 企业发表内容到客户的朋友圈
      - [x] 创建发表任务
      - [x] 获取任务创建结果
    - [x] 获取客户朋友圈全部的发表记录
      - [x] 获取企业全部的发表列表
      - [x] 获取客户朋友圈企业发表的列表
      - [x] 获取客户朋友圈发表时选择的可见范围
      - [x] 获取客户朋友圈发表后的可见客户列表
      - [x] 获取客户朋友圈的互动数据
  - [x] 消息推送
    - [x] 创建企业群发
    - [x] 获取企业的全部群发记录
      - [x] 获取群发记录列表
      - [x] 获取群发成员发送任务列表
      - [x] 获取企业群发成员执行结果
    - [x] 发送新客户欢迎语
    - [x] 入群欢迎语素材管理
  - [x] 统计管理
    - [x] 获取「联系客户统计」数据
    - [x] 获取「群聊数据统计」数据
      - [x] 按群主聚合的方式
      - [x] 按自然日聚合的方式
  - [ ] 变更回调
    - [ ] 添加企业客户事件
    - [ ] 编辑企业客户事件
    - [ ] 外部联系人免验证添加成员事件
    - [ ] 删除企业客户事件
    - [ ] 删除跟进成员事件
    - [ ] 客户接替失败事件
    - [ ] 客户群创建事件
    - [ ] 客户群变更事件
    - [ ] 客户群解散事件
    - [ ] 企业客户标签创建事件
    - [ ] 企业客户标签变更事件
    - [ ] 企业客户标签删除事件
    - [ ] 企业客户标签重排事件
  - [x] 管理商品图册
    - [x] 创建商品图册
    - [x] 获取商品图册
    - [x] 获取商品图册列表
    - [x] 编辑商品图册
    - [x] 删除商品图册
  - [x] 管理聊天敏感词 
    - [x] 新建敏感词规则
    - [x] 获取敏感词规则列表
    - [x] 获取敏感词规则详情
    - [x] 修改敏感词规则
    - [x] 删除敏感词规则
  - [x] 上传附件资源
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
  - [x] 机器人管理
    - [x] 知识库分组管理
      - [x] 添加分组
      - [x] 删除分组
      - [x] 修改分组
      - [x] 获取分组列表
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
  - [x] 自动激活设置
    - [x] 设置企业的许可自动激活状态
    - [x] 查询企业的许可自动激活状态
- [x] 应用管理
  - [x] 获取应用
    - [x] 获取指定的应用详情
    - [x] 获取access_token对应的应用列表
  - [ ] 设置工作台自定义展示
- [x] 消息推送
  - [x] 发送应用消息
  - [x] 发送应用模板消息
  - [x] 更新应用模板消息
  - [ ] 接收消息与事件
  - [x] 撤回应用消息
- [x] 素材管理
  - [x] 上传临时素材
  - [x] 上传图片
  - [x] 获取临时素材
  - [ ] 获取高清语音素材
- [x] 会话内容存档
  - [x] 获取会话内容存档开启成员列表
  - [x] 获取会话同意情况
  - [x] 获取会话内容存档内部群信息
- [x] 电子发票
  - [x] 查询电子发票
  - [x] 更新发票状态
  - [x] 批量更新发票状态
  - [x] 批量查询电子发票
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

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

    ├── example　样例程序
    ├── internal　内部函数，包含error和http请求
    └── pkg 附加包
        └── svr 服务端接口部分
            ├── callback　指令及数据回调
            ├── config  解析yaml配置文件
            ├── install 应用安装连接生成、应用安装后的跳转
            ├── logic　各事件通知处理逻辑
            ├── middleware　gin的中间件方便handler里引入sdk
            └── models  应用安装时常用到的表定义,基于gorm

## 第三方包

- [github.com/dgraph-io/badger/v2 v2.2007.4](https://github.com/dgraph-io/badger) 一个内存数据库,类redis
- [github.com/gin-gonic/gin  v1.7.4](https://github.com/gin-gonic/gin) web框架
- [github.com/go-laoji/wxbizmsgcrypt v1.0.0](https://github.com/go-laoji/wxbizmsgcrypt) 微信数据加解密
- [github.com/jinzhu/copier v0.3.2](https://github.com/jinzhu/copier) jinzhu大神的数据拷贝库
- [go.uber.org/zap v1.19.1](https://go.uber.org/zap) 日志库
- [gopkg.in/natefinch/lumberjack.v2 v2.0.0](https://gopkg.in/natefinch/lumberjack.v2) 日志切割
- [gopkg.in/yaml.v2 v2.2.8](https://gopkg.in/yaml.v2) yaml解析
- [gorm.io/driver/mysql v1.2.0](https://gorm.io/driver/mysql) 同属jinzhu大神的mysql驱动
- [gorm.io/gorm v1.22.3](https://gorm.io/gorm) 同属jinzhu大神的orm库

## API列表

- 应用授权
  - [x] 调用第三方应用凭证
  - [x] 获取预授权码
  - [ ] 设置授权配置
  - [x] 获取企业永久授权码
  - [x] 获取企业授权信息
  - [x] 获取企业凭证　~~getCorpToken~~(存入数据库以主键代替)
  - [x] 获取应用的管理员列表
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
    - [x] 读取成员
    - [x] 获取部门成员
    - [x] 获取部门成员详情
    - [x] userid与openid互换
    - [ ] 二次验证
    - [ ] 邀请成员
    - [x] 手机号获取userid
    - [x] 获取成员授权列表
    - [x] 查询成员用户是否已授权
    - [x] 获取选人ticket对应的用户
  - [x] 部门管理
    - [x] 获取部门列表
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
    - [ ] 统计管理
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
    - [ ] 管理商品图册 
    - [ ] 管理聊天敏感词 
    - [x] 上传附件资源 
- [ ] 应用管理
- [x] 消息推送
  - [x] 发送应用消息
  - [x] 发送应用模板消息
  - [x] 更新应用模板消息
  - [ ] 接收消息与事件
  - [x] 撤回应用消息
- [x] 素材管理
  - [x] 上传临时素材
  - [x] 上传图片
  - [ ] 获取临时素材
  - [ ] 获取高清语音素材
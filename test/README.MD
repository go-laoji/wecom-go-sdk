# 测试代码用法

## wework_test.go

### getAppSecretFunc方法

```go
switch corpId {
case AppSelf:
    //纯自建应用 的corpId和secret
    return "企业ID", "自建应用的secret", true
case AppContact: // 通讯录同步应用 ToB Dev
    return "企业ID", "企业通讯录同步助手的secret", true
case AppContactEdit: // 服务商通讯录可编辑应用
    return "授权企业通讯录编辑应用企业ID(加密的corpId)", "授权企业的永久授权码", false
case AppEhr: // 自建代开发ehr类应用
    return "授权企业的代开发应用企业ID(加密的corpId)", "授权企业的永久授权码", true
}
```

### testNewWeWork 方法

```go
TestWeWork = wework.NewWeWork(wework.WeWorkConfig{
    CorpId:         "服务商corpId",
    SuiteId:        "通讯录可编辑应用的suiteId",
    SuiteSecret:    "通讯录可编辑应用的secret",
    ProviderSecret: "服务商providerSecret",
})
TestWeWork.SetAppSecretFunc(getAppSecretFunc)
TestWeWork.UpdateSuiteTicket("回调回来的suite_ticket")
```

`UpdateSuiteTicket` 的参数应该配合web回调日志拿suite_ticket每15分钟更换一次

## IP白名单

### 服务商IP(服务商后台)

进入到服务商后台，把当前跑代码的出口IP添加到服务商IP白名单中

### 代开发应用(服务商后台)

如果给授权企业配置了应用IP白名单，也要把当前机器IP做修改；位置在`代开发应用模板详情`里找到相应企业的`使用配置`里

### 自建应用(企业后台)

新创建的应用在应用详情里有可信域名配置

# 声明

**单元测试代码测试用例并不完整，仅为验证一些基础权限。因为极少写单测代码有任何意见欢迎指正！**
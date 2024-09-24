package wework

import (
	"encoding/json"
	"net/url"
	"os"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IWeWork interface {
	getProviderToken() string

	GetCorpId() string
	GetSuiteId() string
	GetSuiteToken() string
	GetSuiteEncodingAesKey() string
	Logger() *zap.Logger
	SetAppSecretFunc(f func(corpId uint) (corpid string, secret string, customizedApp bool))
	SetAgentIdFunc(f func(corpId uint) (agentId int))

	GetLoginInfo(authCode string) (resp GetLoginInfoResponse)
	GetUserInfo3rd(code string) (resp GetUserInfo3rdResponse)
	GetUserInfoDetail3rd(userTicket string) (resp GetUserInfoDetail3rdResponse)
	GetUserInfo(corpId uint, code string) (resp GetUserInfoResponse)
	GetUserDetail(corpId uint, userTicket string) (resp GetUserDetailResponse)

	AgentGet(corpId uint, agentId int) (resp AgentGetResponse)
	AgentList(corpId uint) (resp AgentListResponse)

	UpdateSuiteTicket(ticket string)
	getSuiteAccessToken() string
	GetPreAuthCode() (resp GetPreAuthCodeResponse)
	GetPermanentCode(authCode string) (resp GetPermanentCodeResponse)
	GetAuthInfo(authCorpId, permanentCode string) (resp GetAuthInfoResponse)
	GetAppQrCode(request GetAppQrCodeRequest) (resp GetAppQrCodeResponse)

	UserCreate(corpId uint, user User) (resp internal.BizResponse)
	UserUpdate(corpId uint, user User) (resp internal.BizResponse)
	UserDelete(corpId uint, userId string) (resp internal.BizResponse)
	UserGet(corpId uint, userId string) (resp UserGetResponse)
	UserSimpleList(corpId uint, departId int32, fetchChild int) (resp UserSimpleListResponse)
	UserList(corpId uint, departId int32, fetchChild int) (resp UserListResponse)
	UserId2OpenId(corpId uint, userId string) (resp UserId2OpenIdResponse)
	OpenId2UserId(corpId uint, openId string) (resp OpenId2UserIdResponse)
	ListMemberAuth(corpId uint, cursor string, limit int) (resp ListMemberAuthResponse)
	CheckMemberAuth(corpId uint, openUserId string) (resp CheckMemberAuthResponse)
	GetUserId(corpId uint, mobile string) (resp GetUserIdResponse)
	ListSelectedTicketUser(corpId uint, ticket string) (resp ListSelectedTicketUserResponse)
	UserListId(corpId uint, cursor string, limit int) (resp UserListIdResponse)

	CorpTagList(corpId uint, tagIds, groupIds []string) (resp CorpTagListResponse)
	CorpTagAdd(corpId uint, tagGroup CorpTagGroup) (resp CorpTagAddResponse)
	CorpTagUpdate(corpId uint, tag CorpTag) (resp internal.BizResponse)
	CorpTagDelete(corpId uint, tagIds, groupIds []string) (resp internal.BizResponse)
	MarkTag(corpId uint, userId string, externalUserId string, addTag []int, removeTag []int) (resp internal.BizResponse)

	DepartmentCreate(corpId uint, department Department) (resp DepartmentCreateResponse)
	DepartmentUpdate(corpId uint, department Department) (resp internal.BizResponse)
	DepartmentDelete(corpId uint, id int32) (resp internal.BizResponse)
	DepartmentList(corpId uint, id uint) (resp DepartmentListResponse)
	DepartmentSimpleList(corpId uint, id int32) (resp DepartmentSimpleListResponse)
	DepartmentGet(corpId uint, id int32) (resp DepartmentGetResponse)

	ExternalContactGetFollowUserList(corpId uint) (resp ExternalContactGetFollowUserListResponse)
	ExternalContactList(corpId uint, userId string) (resp ExternalContactListResponse)
	ExternalContactGet(corpId uint, externalUserId, cursor string) (resp ExternalContactGetResponse)
	ExternalContactBatchGetByUser(corpId uint, userIds []string, cursor string, limit int) (resp ExternalContactBatchGetByUserResponse)
	ExternalContactRemark(corpId uint, remark ExternalContactRemarkRequest) (resp internal.BizResponse)
	UnionId2ExternalUserId(corpId uint, unionid, openid string) (resp UnionId2ExternalUserIdResponse)
	ToServiceExternalUserid(corpId uint, externalUserId string) (resp ToServiceExternalUseridResponse)

	ExternalAddContactWay(corpId uint, me ContactMe) (resp ContactMeAddResponse)
	ExternalUpdateContactWay(corpId uint, me ContactMe) (resp internal.BizResponse)
	ExternalGetContactWay(corpId uint, configId string) (resp ContactMeGetResponse)
	ExternalListContactWay(corpId uint, startTime, endTime int64, cursor string, limit int) (resp ContactMeListResponse)
	ExternalDeleteContactWay(corpId uint, configId string) (resp internal.BizResponse)
	ExternalCloseTempChat(corpId uint, userId, externalUserId string) (resp internal.BizResponse)

	AddMsgTemplate(corpId uint, msg ExternalMsg) (resp AddMsgTemplateResponse)
	GetGroupMsgListV2(corpId uint, filter GroupMsgListFilter) (resp GetGroupMsgListV2Response)
	GetGroupMsgTask(corpId uint, filter GroupMsgTaskFilter) (resp GetGroupMsgTaskResponse)
	GetGroupMsgSendResult(corpId uint, filter GroupMsgSendResultFilter) (resp GetGroupMsgSendResultResponse)
	SendWelcomeMsg(corpId uint, msg ExternalMsg) (resp internal.BizResponse)

	GetUserBehaviorData(corpId uint, filter GetUserBehaviorFilter) (resp GetUserBehaviorDataResponse)
	GroupChatStatistic(corpId uint, filter GroupChatStatisticFilter) (resp GroupChatStatisticResponse)
	GroupChatStatisticGroupByDay(corpId uint, filter GroupChatStatisticGroupByDayFilter) (resp GroupChatStatisticResponse)

	AddProductAlbum(corpId uint, product Product) (resp AddProductAlbumResponse)
	GetProductAlbum(corpId uint, productId string) (resp GetProductAlbumResponse)
	GetProductAlbumList(corpId uint, limit int, cursor string) (resp GetProductAlbumListResponse)
	UpdateProductAlbum(corpId uint, request ProductUpdateRequest) (resp internal.BizResponse)
	DeleteProductAlbum(corpId uint, productId string) (resp internal.BizResponse)

	AddInterceptRule(corpId uint, interceptRule InterceptRule) (resp AddInterceptRuleResponse)
	GetInterceptRuleList(corpId uint) (resp GetInterceptRuleListResponse)
	GetInterceptRule(corpId uint, ruleId string) (resp GetInterceptRuleResponse)
	UpdateInterceptRule(corpId uint, request UpdateInterceptRuleRequest) (resp internal.BizResponse)
	DeleteInterceptRule(corpId uint, ruleId string) (resp internal.BizResponse)

	GroupChatList(corpId uint, filter GroupChatListFilter) (resp GroupChatListResponse)
	GroupChat(corpId uint, request GroupChatRequest) (resp GroupChatResponse)
	GroupOpengId2ChatId(corpId uint, opengid string) (resp GroupOpengId2ChatIdResponse)

	MediaUploadAttachment(corpId uint, attrs Media) (resp MediaUploadResponse)
	MediaUpload(corpId uint, fileType MediaType, filePath string) (resp MediaUploadResponse)
	MediaUploadImg(corpId uint, filePath string) (resp MediaUploadImgResponse)
	MediaGet(corpId uint, mediaId string) (resp MediaGetResponse)

	GetBillList(corpId uint, req GetBillListRequest) (resp GetBillListResponse)

	MessageSend(corpId uint, msg interface{}) (resp MessageSendResponse)
	MessageReCall(corpId uint, msgId string) (resp internal.BizResponse)

	MessageUpdateTemplateCard(corpId uint, msg TemplateCardUpdateMessage) (resp MessageUpdateTemplateCardResponse)

	AddMomentTask(corpId uint, task MomentTask) (resp AddMomentTaskResponse)
	GetMomentTaskResult(corpId uint, jobId string) (resp GetMomentTaskResultResponse)
	GetMomentList(corpId uint, filter MomentListFilter) (resp GetMomentListResponse)
	GetMomentTask(corpId uint, filter MomentTaskFilter) (resp GetMomentTaskResponse)
	GetMomentCustomerList(corpId uint, filter MomentCustomerFilter) (resp GetMomentCustomerListResponse)
	GetMomentSendResult(corpId uint, filter MomentCustomerFilter) (resp GetMomentSendResultResponse)
	GetMomentComments(corpId uint, momentId string, userId string) (resp GetMomentCommentsResponse)

	TagCreate(corpId uint, tag Tag) (resp TagCreateResponse)
	TagUpdate(corpId uint, tag Tag) (resp internal.BizResponse)
	TagDelete(corpId uint, id int) (resp internal.BizResponse)
	TagList(corpId uint) (resp TagListResponse)
	TagUserList(corpId uint, id int) (resp TagUserListResponse)
	TagAddUsers(corpId uint, tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse)
	TagDelUsers(corpId uint, tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse)

	TransferCustomer(corpId uint, request TransferCustomerRequest) (resp TransferCustomerResponse)
	TransferResult(corpId uint, request TransferResultRequest) (resp TransferResultResponse)
	GetUnassignedList(corpId uint, request UnAssignedRequest) (resp UnAssignedResponse)
	TransferCustomerResigned(corpId uint, request TransferCustomerRequest) (resp TransferCustomerResponse)
	TransferResultResigned(corpId uint, request TransferResultRequest) (resp TransferResultResponse)
	TransferGroupChat(corpId uint, request GroupChatTransferRequest) (resp GroupChatTransferResponse)

	GetInvoiceInfo(corpId uint, query InvoiceInfoQuery) (resp GetInvoiceInfoResponse)
	GetInvoiceInfoBatch(corpId uint, query InvoiceInfoQueryBatch) (resp GetInvoiceInfoBatchResponse)
	UpdateInvoiceStatus(corpId uint, request UpdateInvoiceStatusRequest) (resp internal.BizResponse)
	UpdateInvoiceStatusBatch(corpId uint, request UpdateInvoiceStatusBatchRequest) (resp internal.BizResponse)

	CreateStudent(corpId uint, student Student) (resp internal.BizResponse)
	BatchCreateStudent(corpId uint, students []Student) (resp BatchStudentResponse)
	DeleteStudent(corpId uint, userId string) (resp internal.BizResponse)
	BatchDeleteStudent(corpId uint, userIdList []string) (resp BatchStudentResponse)
	UpdateStudent(corpId uint, student Student) (resp internal.BizResponse)
	BatchUpdateStudent(corpId uint, students []Student) (resp BatchStudentResponse)

	CreateParent(corpId uint, parent Parent) (resp internal.BizResponse)
	BatchCreateParent(corpId uint, parents []Parent) (resp BatchParentResponse)
	DeleteParent(corpId uint, userId string) (resp internal.BizResponse)
	BatchDeleteParent(corpId uint, userIdList []string) (resp BatchParentResponse)
	UpdateParent(corpId uint, parent Parent) (resp internal.BizResponse)
	BatchUpdateParent(corpId uint, parents []Parent) (resp BatchParentResponse)
	ListParentWithDepartmentId(corpId uint, departmentId int32) (resp ListParentWithDepartmentIdResponse)

	SchoolUserGet(corpId uint, userId string) (resp SchoolUserGetResponse)
	SchoolUserList(corpId uint, departmentId uint32, fetchChild int) (resp SchoolUserListResponse)
	SetArchSyncMode(corpId uint, mode int) (resp internal.BizResponse)
	GetSubScribeQrCode(corpId uint) (resp GetSubScribeQrCodeResponse)
	SetSubScribeMode(corpId uint, mode int) (resp internal.BizResponse)
	GetSubScribeMode(corpId uint) (resp GetSubScribeModeResponse)
	BatchToExternalUserId(corpId uint, mobiles []string) (resp BatchToExternalUserIdResponse)
	SetTeacherViewMode(corpId uint, mode int) (resp internal.BizResponse)
	GetTeacherViewMode(corpId uint) (resp GetTeacherViewModeResponse)
	GetAllowScope(corpId uint, agentId int) (resp GetAllowScopeResponse)
	SetUpgradeInfo(corpId uint, request UpgradeRequest) (resp UpgradeInfoResponse)

	SchoolDepartmentCreate(corpId uint, department SchoolDepartment) (resp SchoolDepartmentCreateResponse)
	SchoolDepartmentUpdate(corpId uint, department SchoolDepartment) (resp internal.BizResponse)
	SchoolDepartmentDelete(corpId uint, departmentId int32) (resp internal.BizResponse)
	SchoolDepartmentList(corpId uint, departmentId int32) (resp SchoolDepartmentListResponse)

	GetUserAllLivingId(corpId uint, request GetUserAllLivingIdRequest) (resp GetUserAllLivingIdResponse)
	GetLivingInfo(corpId uint, liveId string) (resp GetLivingInfoResponse)
	GetWatchStat(corpId uint, request GetWatchStatRequest) (resp GetWatchStatResponse)
	GetUnWatchStat(corpId uint, request GetWatchStatRequest) (resp GetUnWatchStatResponse)
	DeleteReplayData(corpId uint, livingId string) (resp internal.BizResponse)

	GetPaymentResult(corpId uint, paymentId string) (resp GetPaymentResultResponse)
	GetTrade(corpId uint, request GetTradeRequest) (resp GetTradeResponse)

	GetJsApiTicket(corpId uint) (resp TicketResponse)
	GetConfigSignature(corpId uint, referer string) (resp JsTicketSignatureResponse)
	GetJsApiAgentTicket(corpId uint, agentId int) (resp TicketResponse)
	GetAgentConfigSignature(corpId uint, agentId int, referer string) (resp JsTicketSignatureResponse)

	KfAccountAdd(corpId uint, account KfAccount) (resp KfAccountAddResponse)
	KfAccountDel(corpId uint, kfId string) (resp internal.BizResponse)
	KfAccountUpdate(corpId uint, account KfAccount) (resp internal.BizResponse)
	KfAccountList(corpId uint, request KfAccountListRequest) (resp KfAccountListResponse)
	KfAddContactWay(corpId uint, kfId string, scene string) (resp KfAccContactWayResponse)
	KfServicerAdd(corpId uint, request KfServicerRequest) (resp KfServicerResponse)
	KfServicerDel(corpId uint, request KfServicerRequest) (resp KfServicerResponse)
	KfServicerList(corpId uint, kfId string) (resp KfServicerListResponse)
	KfServiceStateGet(corpId uint, request KfServiceStateGetRequest) (resp KfServiceStateGetResponse)
	KfServiceStateTrans(corpId uint, request KfServiceStateTransRequest) (resp KfServiceStateTransResponse)
	KfSyncMsg(corpId uint, request KfSyncMsgRequest) (resp KfSyncMsgResponse)
	KfSendMsg(corpId uint, request SendMsgRequest) (resp SendMsgResponse)
	KfSendMsgOnEvent(corpId uint, request SendMsgOnEventRequest) (resp SendMsgResponse)
	KfCustomerBatchGet(corpId uint, userList []string, needEnterSessionContext int) (resp KfCustomerBatchGetResponse)
	// KfGetCorpQualification 仅支持第三方应用，且需具有“微信客服->获取基础信息”权限
	KfGetCorpQualification(corpId uint) (resp KfGetCorpQualificationResponse)
	KfGetUpgradeServiceConfig(corpId uint) (resp KfGetUpgradeServiceConfigResponse)
	KfUpgradeService(corpId uint, request UpgradeServiceRequest) (resp internal.BizResponse)
	KfCancelUpgradeService(corpId uint, request CancelUpgradeServiceRequest) (resp internal.BizResponse)
	// KfGetCorpStatistic
	// 查询时间区间[start_time, end_time]为闭区间，最大查询跨度为31天，用户最多可获取最近180天内的数据。
	// 当天的数据需要等到第二天才能获取，建议在第二天早上六点以后再调用此接口获取前一天的数据
	KfGetCorpStatistic(corpId uint, filter KfGetCorpStatisticFilter) (resp KfGetCorpStatisticResponse)
	// KfGetServicerStatistic
	// 查询时间区间[start_time, end_time]为闭区间，最大查询跨度为31天，用户最多可获取最近180天内的数据。
	// 当天的数据需要等到第二天才能获取，建议在第二天早上六点以后再调用此接口获取前一天的数据
	KfGetServicerStatistic(corpId uint, filter KfGetServicerStatisticFilter) (resp KfGetServicerStatisticResponse)
	KfKnowLedgeAddGroup(corpId uint, name string) (resp KfKnowLedgeAddGroupResponse)
	KfKnowLedgeDelGroup(corpId uint, groupId string) (resp internal.BizResponse)
	KfKnowLedgeModGroup(corpId uint, name string, groupId string) (resp internal.BizResponse)
	KfKnowLedgeListGroup(corpId uint, filter KfKnowLedgeListGroupFilter) (resp KfKnowLedgeListGroupResponse)

	CreateNewOrder(request CreateOrderRequest) (resp OrderResponse)
	CreateReNewOrderJob(request CreateReNewOrderJobRequest) (resp CreateReNewOrderJobResponse)
	SubmitOrderJob(request SubmitOrderJobRequest) (resp OrderResponse)
	ListOrder(request ListOrderRequest) (resp ListOrderResponse)
	GetOrder(request GetOrderRequest) (resp GetOrderResponse)
	ListOrderAccount(request ListOrderAccountRequest) (resp ListOrderAccountResponse)
	ActiveAccount(request ActiveAccountRequest) (resp internal.BizResponse)
	BatchActiveAccount(request BatchActiveAccountRequest) (resp BatchActiveAccountResponse)
	GetActiveInfoByCode(request GetActiveInfoByCodeRequest) (resp GetActiveInfoByCodeResponse)
	BatchGetActiveInfoByCode(request BatchGetActiveInfoByCodeRequest) (resp BatchGetActiveInfoByCodeResponse)
	ListActivedAccount(request ListActivedAccountRequest) (resp ListActivedAccountResponse)
	GetActiveInfoByUser(request GetActiveInfoByUserRequest) (resp GetActiveInfoByUserResponse)
	BatchTransferLicense(request BatchTransferLicenseRequest) (resp BatchTransferLicenseResponse)
	GetAdminList(request GetAdminListRequest) (resp GetAdminListResponse)
	SetAutoActiveStatus(request SetAutoActiveStatusRequest) (resp internal.BizResponse)
	GetAutoActiveStatus(corpid string) (resp GetAutoActiveStatusResponse)

	GetPermitUserList(corpId uint, T int) (resp GetPermitUserListResponse, err error)
	CheckSingleAgree(corpId uint, request CheckSingleAgreeRequest) (resp CheckSingleAgreeResponse, err error)
	GetAuditGroupChat(corpId uint, roomId string) (resp GetAuditGroupChatResponse, err error)
	// ExecuteCorpApi 用于执行未实现的接口，返回 []byte,error
	ExecuteCorpApi(corpId uint, apiUrl string, query url.Values, data H) (body []byte, err error)

	IdConvertExternalTagId(corpId uint, tagIdList []string) (resp IdConvertExternalTagIdResponse)
	CorpIdToOpenCorpId(corpId string) (resp CorpIdToOpenCorpIdResponse)
	UserIdToOpenUserId(corpId uint, userIdList []string) (resp UserIdToOpenUserIdResponse)
	GetNewExternalUserId(corpId uint, userIdList []string) (resp GetNewExternalUserIdResponse)
	GroupChatGetNewExternalUserId(corpId uint, request GroupChatGetNewExternalUserIdRequest) (resp GetNewExternalUserIdResponse)

	RemindGroupMsgSend(corpId uint, msgid string) (resp internal.BizResponse)
	CancelMomentTask(corpId uint, momentId string) (resp internal.BizResponse)

	SetProxy(proxyUrl string)
	SetDebug(debug bool)
}

type weWork struct {
	corpId              string
	providerSecret      string
	suiteId             string
	suiteSecret         string
	suiteTicket         string
	suiteToken          string
	suiteEncodingAesKey string
	cache               *badger.DB
	logger              *zap.Logger
	engine              *gorm.DB
	getAppSecretFunc    func(corpId uint) (corpid string, secret string, customizedApp bool)
	getAgentIdFunc      func(corpId uint) (appId int)
	httpClient          *resty.Client
}

type WeWorkConfig struct {
	CorpId              string
	ProviderSecret      string
	SuiteId             string
	SuiteSecret         string
	SuiteToken          string
	SuiteEncodingAesKey string
	Dsn                 string
}

const (
	UserAgent   = "wecom-go-sdk-v2"
	ContentType = "application/json;charset=UTF-8"
	qyApiHost   = "https://qyapi.weixin.qq.com"
)

func NewWeWork(c WeWorkConfig) IWeWork {
	var ww = new(weWork)
	ww.corpId = c.CorpId
	ww.providerSecret = c.ProviderSecret
	ww.suiteId = c.SuiteId
	ww.suiteSecret = c.SuiteSecret
	ww.suiteToken = c.SuiteToken
	ww.suiteEncodingAesKey = c.SuiteEncodingAesKey
	ww.cache, _ = badger.Open(badger.DefaultOptions("./cache.db").WithIndexCacheSize(10 << 20))
	ww.logger = logger
	ww.httpClient = resty.New().
		SetHeader("User-Agent", UserAgent).
		SetHeader("Content-Type", ContentType).
		SetBaseURL(qyApiHost)
	ww.httpClient.AddRetryCondition(func(r *resty.Response, err error) bool {
		var biz internal.BizResponse
		json.Unmarshal(r.Body(), &biz)
		if biz.ErrCode == 42001 {
			ww.cache.DropPrefix([]byte("corpToken"))
			return true
		}
		return false
	})
	if c.Dsn != "" {
		ww.engine, _ = gorm.Open(mysql.Open(c.Dsn), &gorm.Config{})
	}
	// 默认获取企业token函数
	return ww
}

func (ww *weWork) Logger() *zap.Logger {
	return ww.logger
}
func (ww *weWork) SetProxy(proxyUrl string) {
	if _, ok := url.Parse(proxyUrl); ok == nil {
		ww.httpClient.SetProxy(proxyUrl)
	}
}
func (ww *weWork) SetDebug(debug bool) {
	ww.httpClient.SetDebug(debug)
}

func (ww *weWork) getRequest(corpid uint) *resty.Request {
	R := ww.httpClient.R().SetQueryParam("access_token", ww.getCorpToken(corpid))
	if os.Getenv("debug") != "" {
		R.SetQueryParam("debug", "1")
	}
	return R
}

func (ww *weWork) getProviderRequest() *resty.Request {
	R := ww.httpClient.R().SetQueryParam("provider_access_token", ww.getProviderToken())
	if os.Getenv("debug") != "" {
		R.SetQueryParam("debug", "1")
	}
	return R
}

// GetCorpId 返回服务商corpId
func (ww *weWork) GetCorpId() string {
	return ww.corpId
}

// GetSuiteId 返回服务商SuiteId
func (ww *weWork) GetSuiteId() string {
	return ww.suiteId
}

// GetSuiteToken 返回服务商配置的SuiteToken
func (ww *weWork) GetSuiteToken() string {
	return ww.suiteToken
}

// GetSuiteEncodingAesKey 返回服务商配置的EncodingAesKey
func (ww *weWork) GetSuiteEncodingAesKey() string {
	return ww.suiteEncodingAesKey
}

// GetAgentId 获取应用的AgentId;三方或代开发应用会将信息存入数据库中
// 如果修改了表结构，需要配合 SetAgentIdFunc 使用
func (ww *weWork) GetAgentId(corpId uint) (appId int) {
	if ww.getAgentIdFunc != nil {
		return ww.getAgentIdFunc(corpId)
	} else {
		return ww.defaultAgentIdFunc(corpId)
	}
}

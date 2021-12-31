package wework

import (
	badger "github.com/dgraph-io/badger/v2"
	"github.com/go-laoji/wework/internal"
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

	GetLoginInfo(authCode string) (resp GetLoginInfoResponse)
	GetUserInfo3rd(code string) (resp GetUserInfo3rdResponse)
	GetUserInfoDetail3rd(userTicket string) (resp GetUserInfoDetail3rdResponse)

	AgentGet(corpId uint, agentId int) (resp AgentGetResponse)
	AgentList(corpId uint) (resp AgentListResponse)

	UpdateSuiteTicket(ticket string)
	getSuiteAccessToken() string
	GetPreAuthCode() (resp GetPreAuthCodeResponse)
	GetPermanentCode(authCode string) (resp GetPermanentCodeResponse)
	GetAuthInfo(authCorpId, permanentCode string) (resp GetAuthInfoResponse)

	UserGet(corpId uint, userId string) (resp UserGetResponse)
	UserSimpleList(corpId uint, departId int32, fetchChild int) (resp UserSimpleListResponse)
	UserList(corpId uint, departId int32, fetchChild int) (resp UserListResponse)
	UserId2OpenId(corpId uint, userId string) (resp UserId2OpenIdResponse)
	OpenId2UserId(corpId uint, openId string) (resp OpenId2UserIdResponse)
	ListMemberAuth(corpId uint, cursor string, limit int) (resp ListMemberAuthResponse)
	CheckMemberAuth(corpId uint, openUserId string) (resp CheckMemberAuthResponse)
	GetUserId(corpId uint, mobile string) (resp GetUserIdResponse)
	ListSelectedTicketUser(corpId uint, ticket string) (resp ListSelectedTicketUserResponse)

	CorpTagList(corpId uint, tagIds, groupIds []string) (resp CorpTagListResponse)
	CorpTagAdd(corpId uint, tagGroup CorpTagGroup) (resp CorpTagAddResponse)
	CorpTagUpdate(corpId uint, tag CorpTag) (resp internal.BizResponse)
	CorpTagDelete(corpId uint, tagIds, groupIds []string) (resp internal.BizResponse)
	MarkTag(corpId uint, userId string, externalUserId string, addTag []int, removeTag []int) (resp internal.BizResponse)

	DepartmentList(corpId uint, id uint) (resp DepartmentListResponse)

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
	MediaUpload(corpId uint, fileType MediaType, filePath string, fileName string) (resp MediaUploadResponse)
	MediaUploadImg(corpId uint, filePath string, fileName string) (resp MediaUploadImgResponse)

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
}

type weWork struct {
	IWeWork
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

func NewWeWork(c WeWorkConfig) IWeWork {
	var ww = new(weWork)
	ww.corpId = c.CorpId
	ww.providerSecret = c.ProviderSecret
	ww.suiteId = c.SuiteId
	ww.suiteSecret = c.SuiteSecret
	ww.suiteToken = c.SuiteToken
	ww.suiteEncodingAesKey = c.SuiteEncodingAesKey
	ww.cache, _ = badger.Open(badger.DefaultOptions("").WithInMemory(true))
	ww.logger = logger
	ww.engine, _ = gorm.Open(mysql.Open(c.Dsn), &gorm.Config{})
	return ww
}

func (ww weWork) Logger() *zap.Logger {
	return ww.logger
}

// GetCorpId 返回服务商corpId
func (ww weWork) GetCorpId() string {
	return ww.corpId
}

// GetSuiteId 返回服务商SuiteId
func (ww weWork) GetSuiteId() string {
	return ww.suiteId
}

// GetSuiteToken 返回服务商配置的SuiteToken
func (ww weWork) GetSuiteToken() string {
	return ww.suiteToken
}

// GetSuiteEncodingAesKey 返回服务商配置的EncodingAesKey
func (ww weWork) GetSuiteEncodingAesKey() string {
	return ww.suiteEncodingAesKey
}

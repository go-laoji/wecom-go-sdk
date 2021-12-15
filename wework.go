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

func (ww weWork) GetCorpId() string {
	return ww.corpId
}

func (ww weWork) GetSuiteId() string {
	return ww.suiteId
}

func (ww weWork) GetSuiteToken() string {
	return ww.suiteToken
}

func (ww weWork) GetSuiteEncodingAesKey() string {
	return ww.suiteEncodingAesKey
}
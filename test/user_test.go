package test

import (
	"github.com/go-laoji/wecom-go-sdk"
	"net/url"
	"testing"
)

var testUser = wework.User{
	Userid:     "laoji_test",
	Name:       "老纪",
	Department: []int32{1},
	Email:      "test@test.com",
	Enable:     1,
}
var testUserId = ""

func testAppSelfUserCreate(t *testing.T) {
	t.SkipNow()
}

func testAppSelfUserSimpleList(t *testing.T) {
	resp := TestWeWork.UserSimpleList(AppSelf, 1, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}
func testAppSelfUserList(t *testing.T) {
	resp := TestWeWork.UserList(AppSelf, 1, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppSelfUserListId(t *testing.T) {
	resp := TestWeWork.UserListId(AppSelf, "", 1000)
	switch resp.ErrCode {
	case 0, 48009, 48002:
		t.Log(resp.ErrCode, resp.ErrorMsg)
	default:
		t.Error(resp.ErrorMsg)
	}
}

func testAppSelfUserGet(t *testing.T) {
	resp := TestWeWork.UserGet(AppSelf, "laoji_test")
	// 自建应用没有写入权限，所以此处不会有测试帐号的id输出
	if resp.ErrCode != 60111 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp.ErrorMsg)
	}
}

func testAppContactUserSimpleList(t *testing.T) {
	resp := TestWeWork.UserSimpleList(AppContact, 1, 1)
	if resp.ErrCode != 48009 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppContactUserList(t *testing.T) {
	resp := TestWeWork.UserList(AppContact, 1, 1)
	if resp.ErrCode != 48009 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppContactUserCreate(t *testing.T) {
	resp := TestWeWork.UserCreate(AppContact, testUser)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log("success")
	}
}

func testAppContactUserListId(t *testing.T) {
	resp := TestWeWork.UserListId(AppContact, "", 1000)
	switch resp.ErrCode {
	case 0, 48009:
		t.Log(resp)
	default:
		t.Error(resp.ErrorMsg)
	}
}

func testAppContactUserGet(t *testing.T) {
	resp := TestWeWork.UserGet(AppContact, "laoji_test")
	// API接口无权限调用，为保障企业数据安全，不再允许通讯录同步助手从新增IP读取通讯录详情
	if resp.ErrCode != 48009 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp.ErrorMsg)
	}
}

func testAppContactUserUpdate(t *testing.T) {
	testUser.Department = []int32{1, 2}
	resp := TestWeWork.UserUpdate(AppContact, testUser)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	}
}

func testAppContactUserDelete(t *testing.T) {
	resp := TestWeWork.UserDelete(AppContact, testUser.Userid)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	}
}

func testAppContactEditUserSimpleList(t *testing.T) {
	resp := TestWeWork.UserSimpleList(AppContactEdit, 1, 1)
	if resp.ErrCode != 48008 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppContactEditUserCreate(t *testing.T) {
	resp := TestWeWork.UserCreate(AppContactEdit, testUser)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log("success")
	}
}

func testAppContactEditUserListId(t *testing.T) {
	resp := TestWeWork.UserListId(AppContactEdit, "", 1000)
	switch resp.ErrCode {
	case 0, 48009:
		t.Log(resp)
	default:
		t.Error(resp.ErrorMsg)
	}
}

func testAppContactEditUserGet(t *testing.T) {
	resp := TestWeWork.UserGet(AppContactEdit, "laoji_test")
	// API接口无权限调用，为保障企业数据安全，不再允许通讯录同步助手从新增IP读取通讯录详情
	if resp.ErrCode != 48008 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp.ErrorMsg)
	}
}

func testAppContactEditUserUpdate(t *testing.T) {
	testUser.Department = []int32{1, 2}
	resp := TestWeWork.UserUpdate(AppContactEdit, testUser)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	}
}

func testAppContactEditUserDelete(t *testing.T) {
	resp := TestWeWork.UserDelete(AppContactEdit, testUser.Userid)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	}
}

func testAppContactEditUserList(t *testing.T) {
	resp := TestWeWork.UserList(AppContactEdit, 1, 1)
	if resp.ErrCode != 48008 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppEhrUserCreate(t *testing.T) {
	resp := TestWeWork.UserCreate(AppEhr, testUser)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}
func testAppEhrUserSimpleList(t *testing.T) {
	resp := TestWeWork.UserSimpleList(AppEhr, 1, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppEhrUserList(t *testing.T) {
	resp := TestWeWork.UserList(AppEhr, 1, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
		if len(resp.UserList) > 0 {
			testUserId = resp.UserList[0].Userid
		}
	}
}

func testAppEhrUserListId(t *testing.T) {
	resp := TestWeWork.UserListId(AppEhr, "", 1000)
	switch resp.ErrCode {
	case 0, 48002:
		t.Log(resp)
	default:
		t.Error(resp.ErrorMsg)
	}
}

func testAppEhrUserGet(t *testing.T) {
	resp := TestWeWork.UserGet(AppEhr, testUserId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppEhrUserUpdate(t *testing.T) {
	testUser.Department = []int32{1, 2}
	resp := TestWeWork.UserUpdate(AppEhr, testUser)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppEhrUserDelete(t *testing.T) {
	resp := TestWeWork.UserDelete(AppEhr, testUser.Userid)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
	} else {
		t.Log(resp)
	}
}

func testAppEhrBatchUserId2OpenId(t *testing.T) {
	query := url.Values{}
	h := wework.H{"userid_list": "jifengwei"}
	resp, err := TestWeWork.ExecuteCorpApi(AppEhr, "/cgi-bin/batch/userid_to_openuserid", query, h)
	if err != nil {
		t.Error(string(resp))
	} else {
		t.Log(string(resp), testUserId)
	}
}

func TestUserAll(t *testing.T) {
	t.Run("sdk init", testNewWeWork)
	t.Run("AppSelf UserCreate", testAppSelfUserCreate)
	t.Run("AppSelf UserSimpleList", testAppSelfUserSimpleList)
	t.Run("AppSelf UserList", testAppSelfUserList)
	t.Run("AppSelf UserListId", testAppSelfUserListId)
	t.Run("AppSelf UserGet", testAppSelfUserGet)

	t.Run("AppContact UserCreate", testAppContactUserCreate)
	t.Run("AppContact UserSimpleList", testAppContactUserSimpleList)
	t.Run("AppContact UserList", testAppContactUserList)
	t.Run("AppContact UserListId", testAppContactUserListId)
	t.Run("AppContact UserGet", testAppContactUserGet)
	t.Run("AppContact UserUpdate", testAppContactUserUpdate)
	t.Run("AppContact UserDelete", testAppContactUserDelete)

	t.Run("AppContactEdit UserCreate", testAppContactEditUserCreate)
	t.Run("AppContactEdit UserSimpleList", testAppContactEditUserSimpleList)
	t.Run("AppContactEdit UserList", testAppContactEditUserList)
	t.Run("AppContactEdit UserListId", testAppContactEditUserListId)
	t.Run("AppContactEdit UserGet", testAppContactEditUserGet)
	t.Run("AppContactEdit UserUpdate", testAppContactEditUserUpdate)
	t.Run("AppContactEdit UserDelete", testAppContactEditUserDelete)

	t.Run("AppEhr UserCreate", testAppEhrUserCreate)
	t.Run("AppEhr UserSimpleList", testAppEhrUserSimpleList)
	t.Run("AppEhr UserList", testAppEhrUserList)
	t.Run("AppEhr UserListId", testAppEhrUserListId)
	t.Run("AppEhr UserGet", testAppEhrUserGet)
	t.Run("AppEhr UserUpdate", testAppEhrUserUpdate)
	t.Run("AppEhr UserDelete", testAppEhrUserDelete)
	t.Run("AppEhr BatchUserId2OpenId", testAppEhrBatchUserId2OpenId)

}

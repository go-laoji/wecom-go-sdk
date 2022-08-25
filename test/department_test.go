package test

import (
	"github.com/go-laoji/wecom-go-sdk"
	"testing"
)

var testDepart = wework.Department{
	Id:       5,
	Order:    1,
	ParentId: 1,
	Name:     "技术委员会",
}

func testAppSelfDepartmentCreate(t *testing.T) {
	resp := TestWeWork.DepartmentCreate(AppSelf, testDepart)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppSelfDepartmentUpdate(t *testing.T) {
	testDepart.NameEn = "Technical Committee"
	resp := TestWeWork.DepartmentUpdate(AppSelf, testDepart)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppSelfDepartmentList(t *testing.T) {
	resp := TestWeWork.DepartmentList(AppSelf, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppSelfDepartmentSimpleList(t *testing.T) {
	resp := TestWeWork.DepartmentSimpleList(AppSelf, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppSelfDepartmentGet(t *testing.T) {
	resp := TestWeWork.DepartmentGet(AppSelf, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppSelfDepartmentDelete(t *testing.T) {
	resp := TestWeWork.DepartmentDelete(AppSelf, testDepart.Id)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppContactDepartmentCreate(t *testing.T) {
	resp := TestWeWork.DepartmentCreate(AppContact, testDepart)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppContactDepartmentUpdate(t *testing.T) {
	testDepart.NameEn = "Technical Committee"
	resp := TestWeWork.DepartmentUpdate(AppContact, testDepart)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppContactDepartmentList(t *testing.T) {
	resp := TestWeWork.DepartmentList(AppContact, 1)
	if resp.ErrCode != 48009 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppContactDepartmentSimpleList(t *testing.T) {
	resp := TestWeWork.DepartmentSimpleList(AppContact, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppContactDepartmentGet(t *testing.T) {
	resp := TestWeWork.DepartmentGet(AppContact, 1)
	if resp.ErrCode != 48009 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppContactDepartmentDelete(t *testing.T) {
	resp := TestWeWork.DepartmentDelete(AppContact, testDepart.Id)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppEhrDepartmentCreate(t *testing.T) {
	resp := TestWeWork.DepartmentCreate(AppEhr, testDepart)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppEhrDepartmentUpdate(t *testing.T) {
	testDepart.NameEn = "Technical Committee"
	resp := TestWeWork.DepartmentUpdate(AppEhr, testDepart)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppEhrDepartmentList(t *testing.T) {
	resp := TestWeWork.DepartmentList(AppEhr, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppEhrDepartmentSimpleList(t *testing.T) {
	resp := TestWeWork.DepartmentSimpleList(AppEhr, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppEhrDepartmentGet(t *testing.T) {
	resp := TestWeWork.DepartmentGet(AppEhr, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func testAppEhrDepartmentDelete(t *testing.T) {
	resp := TestWeWork.DepartmentDelete(AppEhr, testDepart.Id)
	if resp.ErrCode != 60011 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestDepartmentAll(t *testing.T) {
	t.Run("sdk init", testNewWeWork)
	t.Run("AppSelf DepartmentCreate", testAppSelfDepartmentCreate)
	t.Run("AppSelf DepartmentUpdate", testAppSelfDepartmentUpdate)
	t.Run("AppSelf DepartmentList", testAppSelfDepartmentList)
	t.Run("AppSelf DepartmentSimpleList", testAppSelfDepartmentSimpleList)
	t.Run("AppSelf DepartmentGet", testAppSelfDepartmentGet)
	t.Run("AppSelf DepartmentDelete", testAppSelfDepartmentDelete)

	t.Run("AppContact DepartmentCreate", testAppContactDepartmentCreate)
	t.Run("AppContact DepartmentUpdate", testAppContactDepartmentUpdate)
	t.Run("AppContact DepartmentList", testAppContactDepartmentList)
	t.Run("AppContact DepartmentSimpleList", testAppContactDepartmentSimpleList)
	t.Run("AppContact DepartmentGet", testAppContactDepartmentGet)
	t.Run("AppContact DepartmentDelete", testAppContactDepartmentDelete)

	t.Run("AppEhr DepartmentCreate", testAppEhrDepartmentCreate)
	t.Run("AppEhr DepartmentUpdate", testAppEhrDepartmentUpdate)
	t.Run("AppEhr DepartmentList", testAppEhrDepartmentList)
	t.Run("AppEhr DepartmentSimpleList", testAppEhrDepartmentSimpleList)
	t.Run("AppEhr DepartmentGet", testAppEhrDepartmentGet)
	t.Run("AppEhr DepartmentDelete", testAppEhrDepartmentDelete)

}

package dbops

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("trunncate video_info")
	dbConn.Exec("trunncate comments")
	dbConn.Exec("trunncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// user 的 workflow
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("zhangsan", "123")
	if err != nil {
		t.Errorf("error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("zhangsan")
	if err != nil {
		t.Errorf("error of getUserPwd: %v", err)
	}

	if pwd != "123" {
		t.Errorf("user pwd get error: %v", pwd)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("zhangsan", "123")
	if err != nil {
		t.Errorf("error of DeleteUser: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("zhangsan")
	if err != nil {
		t.Errorf("error of getUserPwd: %v", err)
	}

	if pwd != "" {
		t.Errorf("deleting user faild")
	}
}

// video 的 workflow
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PreparedUser", testAddUser)
	t.Run("AddVideo", testAddVideo)
	t.Run("GetVideo", testGetVideInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testGetVideInfo)
}

var tempVid string
func testAddVideo(t *testing.T)  {
	video, err := AddNewVideo(1, "test-video")
	if err != nil {
		t.Errorf("error of AddNewVideo: %v", err)
	}
	tempVid = video.Id
}

func testGetVideInfo(t *testing.T)  {
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T)  {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("error of DeleteVideoInfo: %v", err)
	}
}
package dbops

import "testing"

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

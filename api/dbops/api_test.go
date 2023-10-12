package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
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
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCreatential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser : %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCreatential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser : %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeletUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCreatential("avenssi")
	if err != nil {
		t.Errorf("Error of Reget:%v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

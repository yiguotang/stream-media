package dbops

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"stream-media/api/defs"
	"stream-media/api/utils"
	"time"
)

// 新增用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	// 这里加defer是为了上面的if，可以进行关闭资源
	defer stmtIns.Close()
	return nil
}

// 根据用户名查询用户密码
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
 		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

// 删除用户
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("delete user error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

// 添加视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID();
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("insert into video_info (id, author_id, name, display_ctime) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	// 判断sql执行的结果是否有错误
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		Id: vid,
		AuthorId: aid,
		Name: name,
		DisplayCtime: ctime,
	}

	defer stmtIns.Close()
	return res, nil
}

// 查询视频
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id, name, display_ctime from video_info where id = ?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{
		Id: vid,
		AuthorId: aid,
		Name: name,
		DisplayCtime: dct,
	}

	return res, nil
}

// 删除视频
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_info where id = ?")
	if err != nil {
		log.Printf("delete video error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

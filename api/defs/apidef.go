package defs

// user model
type UserCredential struct {
	User string `json: "user_name"`
	Pwd  string `json: "pwd"`
}

// video model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

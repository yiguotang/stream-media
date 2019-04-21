package defs

// requets model
type UserCredential struct {
	User string `json: "user_name"`
	Pwd  string `json: "pwd"`
}

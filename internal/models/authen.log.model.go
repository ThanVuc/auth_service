package models

type LogString interface {
	String() string
}

type AuthenLogStruct struct {
	UserId string `json:"user_id" bson:"user_id"`
	Action string `json:"action" bson:"action"`
}

func NewAuthenLog(userId string, action string) *AuthenLogStruct {
	return &AuthenLogStruct{
		UserId: userId,
		Action: action,
	}
}

func (a AuthenLogStruct) String() string {
	return "User ID: " + a.UserId + ", Action: " + a.Action
}

package model

type Version struct {
	VersionId int    `json:"version_id" gorm:"primaryKey"`
	Version   string `json:"version" gorm:"not null"`
}

type Title struct {
	TitleId   int    `json:"promise_id" gorm:"primaryKey"`
	VersionId int    `json:"version_id" gorm:"not null"`
	Title     string `json:"promise" gorm:"not null"`
}

type List struct {
	ListId    int    `json:"list_id" gorm:"primaryKey"`
	VersionId int    `json:"version_id" gorm:"not null"`
	TitleId   int    `json:"promise_id" gorm:"not null"`
	List      string `json:"list" gorm:"not null"`
}

type Signature struct {
	UserId        int `json:"user_id" gorm:"not null"`
	PromiseStatus int `json:"promise_status" gorm:"not null"`
}

type Promises struct {
	Version  string         `json:"version"`
	Promise  []PromisesPart `json:"promise"`
	IsSubmit bool           `json:"is_submit"`
}

type PromisesPart struct {
	Title string   `json:"title"`
	List  []string `json:"list"`
}

type RespPromises struct {
	Status int      `json:"status"`
	Info   string   `json:"info"`
	Data   Promises `json:"data"`
}

type ReqPromises struct {
	VersionId int `json:"version_id" form:"version_id" binding:"required"`
	UserId    int `json:"user_id" form:"user_id" binding:"required"`
}

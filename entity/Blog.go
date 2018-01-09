package entity

type Blog struct {
	Id         int
	Title      string
	Content    string
	TagId      int
	WordNum    int
	ReadNum    int
	CommentNum int
	LikeNum    int
	Version    string
	UserId     int
	CreateDate string
	UpdateDate string
}

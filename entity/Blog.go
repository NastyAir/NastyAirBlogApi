package entity

type Blog struct {
	Id         int
	BlogId     string
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

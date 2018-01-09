package dao

import (
	"fmt"
	"NastyAir/Blog/entity"
)

func BlogInsert(blog entity.Blog) (int64, error) {
	stmt, err := Db.Prepare(`INSERT blog(id,title,content,tag_id,word_num,read_num,comment_num,like_num,version,user_id,create_date,update_date) values (?,?,?,?,?,?,?,?,?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(
		blog.Id,
		blog.Title,
		blog.Content,
		blog.TagId,
		blog.WordNum,
		blog.ReadNum,
		blog.CommentNum,
		blog.LikeNum,
		blog.Version,
		blog.UserId,
		blog.CreateDate,
		blog.UpdateDate,
	)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	return id, err
}

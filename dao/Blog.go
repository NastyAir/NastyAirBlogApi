package dao

import (
	"NastyAir/Blog/entity"
	"os/user"
)

func BlogInsert(blog entity.Blog) (int64, error) {
	stmt, err := Db.Prepare(`INSERT blog(id,blog_id,title,content,tag_id,word_num,read_num,comment_num,like_num,version,user_id,create_date,update_date) values (?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		//checkErr(err)
		return -1, err
	}

	res, err := stmt.Exec(
		blog.Id,
		blog.BlogId,
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
	if err != nil {
		//checkErr(err)
		return -1, err
	}
	id, err := res.LastInsertId()
	checkErr(err)
	//fmt.Println(id)
	return id, err
}

func BlogSelectByPage(limit, size int) (entity.Blog, error) {
	rows, err := Db.Query("SELECT id,	blog_id, title, content, tag_id, word_num, read_num, comment_num,	like_num, version, user_id, create_date, update_date FROM blog LIMIT ?,?",limit,size)
	checkErr(err)
	var blog entity.Blog
	for rows.Next() {
		err = rows.Scan(
			&blog.Id,
			&blog.BlogId,
			&blog.Title,
			&blog.Content,
			&blog.TagId,
			&blog.WordNum,
			&blog.ReadNum,
			&blog.CommentNum,
			&blog.LikeNum,
			&blog.Version,
			&blog.UserId,
			&blog.CreateDate,
			&blog.UpdateDate,)
	}
	defer rows.Close()
	//checkErr(err)
	return blog,err
}

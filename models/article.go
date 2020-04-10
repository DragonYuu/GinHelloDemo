package models

import (
	"gin_blog/dao"
	"gin_blog/gb"
)

const (
	pageSize = 4
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64 `db:"create_time"`
	Status     int   // Status=0为正常，1为删除，2为冻结
}

//---------数据处理-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

//-----------数据库操作---------------

// 插入一篇文章
func insertArticle(article Article) (int64, error) {
	return dao.ModifyDB("insert into article(title,tags,short,content,author,create_time,status) values(?,?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime, article.Status)
}

// 查询所有文章

/**
分页查询数据库
limit分页查询语句，
    语法：limit m,n

    m代表从多少位开始获取，与id值无关
    n代表获取多少条数据


	总共有10条数据，每页显示4条。  --> 总共需要(10-1)/4+1 页。
	问第2页数据是哪些？           --> 5,6,7,8  (2-1)*4,4

*/
func QueryArticleWithPage(pageNum int) (articleList []*Article, err error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from article limit ?,?"
	articleList, err = queryArticleWithCon(pageNum, sqlStr)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func QueryCurrUserArticleWithPage(username string, pageNum int) (articleList []*Article, err error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from article where author=? limit ?,?"
	articleList, err = queryArticleWithCon(pageNum, sqlStr, username)
	if err != nil {
		gb.Logger.Error("queryArticleWithCon, ", err)
		return nil, err
	}
	gb.Logger.Debug("QueryCurrUserArticleWithPage,", articleList)
	return articleList, nil
}

// 根据查询条件查询指定页数有的文章
func queryArticleWithCon(pageNum int, sqlStr string, args ...interface{}) (articleList []*Article, err error) {
	pageNum--
	args = append(args, pageNum*pageSize, pageSize)
	gb.Logger.Debug("queryArticleWithCon", sqlStr, args)
	err = dao.QueryRows(&articleList, sqlStr, args...)
	gb.Logger.Debug("dao.QueryRows,", articleList)
	return
}

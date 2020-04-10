package models

import (
	"fmt"
	"gin_blog/utils"
	"strings"
)

type HomeBlockParam struct {
	//Id         int
	//Title      string
	//Tags       [] TagLink
	//Short      string
	//Content    string
	//Author     string
	//CreateTime string
	Article *Article

	TagLinks      []*TagLink
	CreateTimeStr string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

//将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tagStr string) []*TagLink {
	var tagLinks = make([]*TagLink, 0, strings.Count(tagStr, "&"))
	tagList := strings.Split(tagStr, "&")
	for _, tag := range tagList {
		tagLinks = append(tagLinks, &TagLink{tag, "/?tag=" + tag})
	}
	return tagLinks
}

func GenHomeBlocks(articleList []*Article, isLogin bool) (ret []*HomeBlockParam) {
	ret = make([]*HomeBlockParam, 0, len(articleList)) // 内存申请一次到位
	for _, art := range articleList {
		// 将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{
			Article: art,
			IsLogin: isLogin,
		}
		homeParam.TagLinks = createTagsLinks(art.Tags)
		homeParam.CreateTimeStr = utils.SwitchTimeStampToStr(art.CreateTime)

		homeParam.Link = fmt.Sprintf("/show/%d", art.Id)
		homeParam.UpdateLink = fmt.Sprintf("/article/update?id=%d", art.Id)
		homeParam.DeleteLink = fmt.Sprintf("/article/delete?id=%d", art.Id)
		ret = append(ret, &homeParam) // 不再需要动态扩容
	}
	return
}

//----------首页显示内容---------
//func MakeHomeBlocks(articleList []*Article, isLogin bool) template.HTML {
//	htmlHome := ""
//	for _, art := range articleList {
//		// 将数据库model转换为首页模板所需要的model
//		homeParam := HomeBlockParam{
//			Article:art,
//			IsLogin:isLogin,
//		}
//
//		homeParam.TagLinks = createTagsLinks(art.Tags)
//		homeParam.CreateTimeStr = utils.SwitchTimeStampToData(art.Createtime)
//
//		homeParam.Link = fmt.Sprintf("/show/%d", art.Id)
//		homeParam.UpdateLink = fmt.Sprintf("/article/update?id=%d", art.Id)
//		homeParam.DeleteLink = fmt.Sprintf("/article/delete?id=%d", art.Id)
//
//		//处理变量
//		//ParseFile解析该文件，用于插入变量
//		t, _ := template.ParseFiles("views/home_block.html")
//		buffer := bytes.Buffer{}
//		//就是将html文件里面的比那两替换为穿进去的数据
//		t.Execute(&buffer, homeParam)
//		htmlHome += buffer.String()
//	}
//	fmt.Println("htmlHome-->",htmlHome)
//	return template.HTML(htmlHome)
//}

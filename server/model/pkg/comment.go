package pkg

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/system"
)

type Comment struct {
	global.GVA_MODEL
	PostId      *int   `json:"post_id" form:"post_id" gorm:"foreignKey:TermId;index:post_id;column:post_id;comment:该评论所属的帖子或页面的 ID;"`
	Author      string `json:"author" form:"author" gorm:"type:tinytext;column:author;comment:评论作者的名称或用户名;"`
	AuthorEmail string `json:"author_email" form:"author_email" gorm:"size:100;column:author_email;comment:评论作者的电子邮件地址;"`
	AuthorUrl   string `json:"author_url" form:"author_url" gorm:"column:author_url;comment:评论作者的网站或博客的 URL;"`
	AuthorIP    string `json:"author_IP" form:"author_IP" gorm:"size:100;column:author_IP;comment:评论作者的 IP 地址;"`
	Content     string `json:"content" form:"content" gorm:"type:text;column:content;comment:评论的内容;"`
	Karma       *int   `json:"karma" form:"karma" gorm:"column:karma;comment:评论的评分;"`
	Approved    string `json:"approved" form:"approved" gorm:"size:20;index:approved;column:approved;comment:评论是否已被批准，0 表示未批准，1 表示已批准, 2 表示评论已被标记为垃圾评论;"`
	Agent       string `json:"agent" form:"agent" gorm:"size:255;column:agent;comment:评论的浏览器和操作系统信息;"`
	CommentType string `json:"comment_type" form:"comment_type" gorm:"size:20;column:comment_type;comment:评论的类型;"`
	ParentId    *int   `json:"parent_id" form:"parent_id" gorm:"foreignKey:TermId;index:parent_id;column:parent_id;comment:该评论所属的父级评论 ID，如果该评论是顶级评论，则值为 0;"`
	UserId      *int   `json:"user_id" form:"user_id" gorm:"size:200;foreignKey:ID;index:user_id;column:user_id;comment:如果评论是由注册用户发布的，则显示用户的 ID;"`

	User     system.SysUser `gorm:"foreignKey:UserId;"`
	Post     Post           `gorm:"foreignKey:PostId"`
	Children []Comment      `json:"children" gorm:"-"`

	CommentMeta []CommentMeta
}

type CommentMeta struct {
	UmetaId   *int   `gorm:"primarykey"` // 主键ID
	CommentId *int   `json:"comment_id" form:"comment_id" gorm:"index:comment_id;column:comment_id;comment:;default:0;"`
	MetaKey   string `json:"meta_key" form:"meta_key" gorm:"index:meta_key;column:meta_key;comment:;"`
	MetaValue string `json:"meta_value" form:"meta_value" gorm:"column:meta_value;comment:;"`
}

// TableName TermStruct 表名
func (Comment) TableName() string {
	return "Comments"
}

// TableName TermStruct 表名
func (CommentMeta) TableName() string {
	return "commentmeta"
}

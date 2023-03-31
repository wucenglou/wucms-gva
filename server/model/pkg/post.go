package pkg

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/system"
)

// TermStruct 结构体
type Post struct {
	global.GVA_MODEL
	PostAuthor          uint   `json:"post_author" form:"post_author" gorm:"foreignKey:TermId;index:post_author;column:post_author;comment:文章作者的用户ID;"`
	PostTitle           string `json:"post_title" form:"post_title" gorm:"type:text;column:post_title;comment:文章标题;"`
	PostExcerpt         string `json:"post_excerpt" form:"post_excerpt" gorm:"type:text;column:post_excerpt;comment:文章摘要;"`
	PostContent         string `json:"post_content" form:"post_content" gorm:"type:longtext;column:post_content;comment:文章内容;"`
	PostContentFiltered string `json:"post_content_filtered" form:"post_content_filtered" gorm:"type:longtext;column:post_content_filtered;comment:过滤后的文章内容;"`
	PostParent          uint   `json:"post_parent" form:"post_parent" gorm:"column:post_parent;comment:父文章的ID;default:0;"`
	PostType            string `json:"post_type" form:"post_type" gorm:"size:20;index:post_type;index:post_status;column:post_type;comment:文章类型;default:'post';"`
	PostStatus          string `json:"post_status" form:"post_status" gorm:"size:20;index:post_status;column:post_status;comment:;default:'publish';comment:Publish发布Draft草稿Pending Review待审核Private私有Future未来Trash垃圾箱reject拒绝Auto-Draft自动草稿"`
	PostName            string `json:"post_name" form:"post_name" gorm:"size:200;index:post_name;column:post_name;comment:文章的永久链接（slug）;"`

	PostPassword string `json:"post_password" form:"post_password" gorm:"size:255;column:post_password;comment:;"`
	MenuOrder    int    `json:"menu_order" form:"menu_order" gorm:"column:menu_order;comment:文章在菜单中的排序位置;default:0;"`

	PostImg       string `json:"post_img" form:"post_img" gorm:"column:post_img;comment:文章配图;"`
	PostMimeType  string `json:"post_mime_type" form:"post_mime_type" gorm:"size:100;column:post_mime_type;comment:附件的MIME类型;"`
	CommentStatus string `json:"comment_status" form:"comment_status" gorm:"size:20;column:comment_status;comment:;default:'open';"`
	CommentCount  int    `json:"comment_count" form:"comment_count" gorm:"column:comment_count;comment:;default:0;"`

	User system.SysUser `gorm:"foreignKey:PostAuthor;"`

	TermTaxonomy []TermTaxonomy `json:"termtaxonomy" gorm:"many2many:TermRelationships;foreignKey:ID;joinForeignKey:ObjectId;References:TermTaxonomyId;joinReferences:TermTaxonomyId"`

	PostMeta []PostMeta

	// TermTaxonomyId []TermTaxonomy `gorm:"many2many:term_relationships;foreignKey:PostID;"`
}

type PostMeta struct {
	MetaId    *int   `gorm:"primarykey"` // 主键ID
	PostId    *int   `json:"post_id" form:"post_id" gorm:"index:post_id;column:post_id;comment:;default:0;"`
	MetaKey   string `json:"meta_key" form:"meta_key" gorm:"index:meta_key;column:meta_key;comment:;"`
	MetaValue string `json:"meta_value" form:"meta_value" gorm:"column:meta_value;comment:;"`
}

// TableName TermStruct 表名
func (Post) TableName() string {
	return "posts"
}

// TableName TermStruct 表名
func (PostMeta) TableName() string {
	return "postmeta"
}

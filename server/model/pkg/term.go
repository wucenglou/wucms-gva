package pkg

import (
	"time"

	"gorm.io/gorm"
)

// TermStruct 结构体
type Term struct {
	TermId       uint         `json:"term_id" form:"term_id" gorm:"primarykey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 主键ID
	Name         string       `json:"name" form:"name" gorm:"index:name;column:name;comment:;"`
	Slug         string       `json:"slug" form:"slug" gorm:"index:slug;column:slug;comment:;"`
	TermGroup    int          `json:"term_group" form:"term_group" gorm:"column:term_group;comment:;default:0;"`
	TermTaxonomy TermTaxonomy `json:"term_taxonomy" form:"term_taxonomy" gorm:"foreignKey:TermId;onDelete:CASCADE"`

	TermMetas []TermMeta `json:"termmeta" gorm:"foreignKey:TermId;references:TermId;onDelete:CASCADE"`
}

// TermStruct 结构体
type TermTaxonomy struct {
	TermTaxonomyId uint   `gorm:"primarykey"` // 主键ID
	TermId         int    `json:"term_id" form:"term_id" gorm:"index:term_id_taxonomy,unique;column:term_id;comment:;"`
	Taxonomy       string `json:"taxonomy" form:"taxonomy" gorm:"index:taxonomy;index:term_id_taxonomy,unique;column:taxonomy;comment:;"`
	Description    string `json:"description" form:"description" gorm:"type:longtext;column:description;comment:;"`
	ParentID       int    `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:;default:0;"`
	Count          int    `json:"count" form:"count" gorm:"column:count;comment:;default:0;"`

	Parent   *TermTaxonomy   `gorm:"foreignKey:ParentID"`                 // 定义父级关联关系
	Children []*TermTaxonomy `json:"children" gorm:"foreignKey:ParentID"` // 定义子级关联关系
	Posts    []Post          `json:"posts" gorm:"many2many:TermRelationships;foreignKey:TermTaxonomyId;joinForeignKey:TermTaxonomyId;References:ID;joinReferences:ObjectId"`
	// Posts []Post `gorm:"many2many:term_relationships"`
	Term *Term `gorm:"foreignKey:TermId;references:TermId;"` // 定义 Term 关联关系
}

type TermRelationship struct {
	// PostID uint `gorm:"primarykey"` // 主键ID
	ObjectId       uint `gorm:"primarykey;"` // 主键ID
	TermTaxonomyId uint `gorm:"primarykey"`  // 主键ID
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type TermMeta struct {
	MetaId    *int   `gorm:"primarykey"` // 主键ID
	TermId    *int   `json:"term_id" form:"term_id" gorm:"index:term_id;column:term_id;comment:;default:0;"`
	MetaKey   string `json:"meta_key" form:"meta_key" gorm:"index:meta_key;column:meta_key;comment:;"`
	MetaValue string `json:"meta_value" form:"meta_value" gorm:"column:meta_value;comment:;"`
}

func (TermRelationship) TableName() string {
	return "term_relationships"
}

func (TermTaxonomy) TableName() string {
	return "term_taxonomy"
}

func (Term) TableName() string {
	return "terms"
}

// TableName TermStruct 表名
func (TermMeta) TableName() string {
	return "termmeta"
}

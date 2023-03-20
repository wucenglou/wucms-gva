package pkg

import (
	"time"

	"gorm.io/gorm"
)

// TermStruct 结构体
type Term struct {
	TermId       uint         `gorm:"primarykey"` // 主键ID
	Name         string       `json:"name" form:"name" gorm:"index:name;column:name;comment:;"`
	Slug         string       `json:"slug" form:"slug" gorm:"index:slug;column:slug;comment:;"`
	TermGroup    *int         `json:"term_group" form:"term_group" gorm:"column:term_group;comment:;"`
	TermTaxonomy TermTaxonomy `gorm:"foreignKey:TermId;"`

	TermMeta []TermMeta `json:"termmeta" gorm:"foreignKey:TermId;references:TermId;"`
}

// TermStruct 结构体
type TermTaxonomy struct {
	TermTaxonomyId uint   `gorm:"primarykey"` // 主键ID
	TermId         *int   `json:"term_id" form:"term_id" gorm:"index:term_id_taxonomy,unique;column:term_id;comment:;"`
	Taxonomy       string `json:"taxonomy" form:"taxonomy" gorm:"index:taxonomy;index:term_id_taxonomy,unique;column:taxonomy;comment:;"`
	Description    string `json:"description" form:"description" gorm:"type:longtext;column:description;comment:;"`
	Parent         *int   `json:"parent" form:"parent" gorm:"column:parent;comment:;"`
	Count          *int   `json:"count" form:"count" gorm:"column:count;comment:;"`

	Posts []Post `gorm:"many2many:TermRelationship;foreignKey:TermTaxonomyId;joinForeignKey:TermTaxonomyId;References:ID;joinReferences:ObjectId"`
	// Post []Post `gorm:"many2many:term_relationships"`
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

// TableName TermStruct 表名
func (TermRelationship) TableName() string {
	return "term_relationships"
}

// TableName TermStruct 表名
func (TermTaxonomy) TableName() string {
	return "term_taxonomy"
}

// TableName TermStruct 表名
func (Term) TableName() string {
	return "terms"
}

// TableName TermStruct 表名
func (TermMeta) TableName() string {
	return "termmeta"
}

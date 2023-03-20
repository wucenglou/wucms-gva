package pkg

import (
	"fmt"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
)

type MyApi struct{}

func (m *MyApi) CreateApi(c *gin.Context) {
	var Term pkg.Term
	res, _ := utils.GetClaims(c)
	fmt.Println(res.ID)
	fmt.Println(res.Username)
	fmt.Println(res.NickName)
	err := global.GVA_DB.Model(&pkg.Term{}).Preload("TermMeta").Preload("TermTaxonomy").Find(&Term).Error
	if err != nil {
		return
	}
	fmt.Println(Term)
	for i := 0; i < len(Term.TermMeta); i++ {
		fmt.Println(Term.TermMeta[i].MetaKey)
		fmt.Println(Term.TermMeta[i].MetaValue)
		// for k, v := range Term.TermMeta[i] {
		// 	fmt.Println(k)
		// 	fmt.Println(v)
		// }
	}

	var TermTaxonomy pkg.TermTaxonomy
	err = global.GVA_DB.Model(&pkg.TermTaxonomy{}).Preload("Posts").Find(&TermTaxonomy).Error
	if err != nil {
		return
	}
	fmt.Println(TermTaxonomy)
	// var total int64
	fmt.Println("*********-***********")
	// db := global.GVA_DB.Model(&system.SysUser{})
	// err := db.Count(&total).Error
	// if err != nil {
	// 	return
	// }
	// var userList []system.SysUser
	// err = db.Preload("Authorities").Preload("Authority").Find(&userList).Error
	// if err != nil {
	// 	return
	// }
	// response.OkWithDetailed(response.PageResult{
	// 	List: userList,
	// }, "获取成功", c)
	response.Ok(c)
}

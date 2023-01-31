package system

import (
	"errors"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/system"

	"gorm.io/gorm"
)

type ApiService struct{}

var ApiServiceApp struct{}

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: DeleteApi
// @description: 删除基础api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	success := CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if !success {
		return errors.New(entity.Path + ":" + entity.Method + "casbin同步清理失败")
	}
	e := CasbinServiceApp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetAPIInfoList
// @description: 分页获取数据,
// @param: api model.SysApi, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysApi{})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return:  apis []model.SysApi, err error

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

package system

import (
	"errors"
	"fmt"
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
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "*"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error
	if err != nil {
		fmt.Println("99999999999999999999")
		return apiList, total, err
	} else {
		fmt.Println("88888888888888888888")
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't matched any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return:  apis []model.SysApi, err error

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetApiById
// @description: 根据id获取api
// @param: id float64
// @return: api model.SysApi, err error
func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: UpdateApi
// @description: 根据id更新api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: DeleteApis
// @description: 删除选中API
// @param: apis []model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	var apis []system.SysApi
	err = global.GVA_DB.Find(&apis, "id in ?", ids.Ids).Delete(&apis).Error
	if err != nil {
		return err
	} else {
		for _, sysApi := range apis {
			success := CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
			if !success {
				return errors.New(sysApi.Path + ":" + sysApi.Method + "casbin同步清理失败")
			}
		}
		e := CasbinServiceApp.Casbin()
		err = e.InvalidateCache()
		if err != nil {
			return err
		}
	}
	return err
}

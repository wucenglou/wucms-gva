package system

import (
	"wucms-gva/server/global"
	"wucms-gva/server/model/system"
)

type ApiService struct{}

var ApiServiceApp struct{}

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return:  apis []model.SysApi, err error

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

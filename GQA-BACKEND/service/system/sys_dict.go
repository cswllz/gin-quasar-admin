package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceDict struct {
}

func (s *ServiceDict)GetDictList(pageInfo system.RequestPageWithParentId) (err error, role interface{}, total int64, parentId uint) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Where("parent_id=?", pageInfo.ParentId).Find(&system.SysDict{})
	var dictList []system.SysDict
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&dictList).Error
	return err, dictList, total, pageInfo.ParentId
}

func (s *ServiceDict) EditDict(dict system.SysDict) (err error) {
	err = global.GqaDb.Updates(&dict).Error
	return err
}

func (s *ServiceDict) AddDict(d system.SysDict) (err error) {
	var dict system.SysDict
	if !errors.Is(global.GqaDb.Where("value = ?", d.Value).First(&dict).Error, gorm.ErrRecordNotFound) {
		return errors.New("此字典已存在：" + d.Value)
	}
	err = global.GqaDb.Create(&d).Error
	return err
}

func (s *ServiceDict) DeleteDict(id uint) (err error) {
	var dict system.SysDict
	err = global.GqaDb.Where("id = ?", id).Delete(&dict).Error
	return err
}

func (s *ServiceDict) QueryDictById(id uint) (err error, dictInfo system.SysDict) {
	var dict system.SysDict
	err = global.GqaDb.First(&dict, "id = ?", id).Error
	return err, dict
}

func (s *ServiceDict) QueryDictByParentId(parentId uint) (err error, dictInfo system.SysDict) {
	var dict system.SysDict
	err = global.GqaDb.First(&dict, "parent_id = ?", parentId).Error
	return err, dict
}
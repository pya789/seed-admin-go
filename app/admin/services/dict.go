package services

import (
	"errors"
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/request"
	"seed-admin/common"
)

type DictService struct{}

// 获取所有字典
func (*DictService) GetAllDictType(params *request.DictList) ([]entity.AdminSysDictType, int64, error) {
	dict := new(entity.AdminSysDictType)
	dicts := make([]entity.AdminSysDictType, 0)
	count, err := common.DB.
		Where("name LIKE ? AND status LIKE ?", "%"+params.Name+"%", "%"+params.Status+"%").
		Count(dict)
	if err != nil {
		return nil, 0, err
	}
	if err := common.DB.
		Where("name LIKE ? AND status LIKE ?", "%"+params.Name+"%", "%"+params.Status+"%").
		Limit(*params.PageSize, (*params.PageNum-1)*(*params.PageSize)).
		Find(&dicts); err != nil {
		return nil, 0, err
	}
	return dicts, count, nil
}

// 增加字典
func (*DictService) AddDictType(params *request.DictAdd) error {
	dictType := &entity.AdminSysDictType{
		Name:   *params.Name,
		Type:   *params.Type,
		Status: params.Status,
	}
	// 插入角色表
	if _, err := common.DB.Insert(dictType); err != nil {
		return err
	}
	return nil
}

// 获取字典类型信息
func (*DictService) GetDictTypeInfo(id *int) (*entity.AdminSysDictType, error) {
	dictType := new(entity.AdminSysDictType)
	if ok, err := common.DB.Where("id = ?", id).Get(dictType); !ok {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("获取字典类型信息失败")
	}
	return dictType, nil
}

// 更新字典
func (*DictService) UpdateDictType(params *request.DictUpdate) error {
	dictType := &entity.AdminSysDictType{
		Name:   *params.Name,
		Type:   *params.Type,
		Status: params.Status,
	}
	if _, err := common.DB.Where("id = ?", params.Id).AllCols().Update(dictType); err != nil {
		return err
	}
	return nil
}

// 删除字典
func (*DictService) DelDictType(params *request.DictDel) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	// 删除类型
	dictType := new(entity.AdminSysDictType)
	if _, err := common.DB.In("id", params.Ids).Delete(dictType); err != nil {
		return err
	}
	dictData := new(entity.AdminSysDictData)
	// 删除类型下的数据
	if _, err := session.In("pid", params.Ids).Delete(dictData); err != nil {
		return err
	}
	return session.Commit()
}

// 获取词典的全部数据
func (*DictService) GetAllDictData(params *request.DictDataList) ([]entity.AdminSysDictData, int64, error) {
	dictData := new(entity.AdminSysDictData)
	dictDatas := make([]entity.AdminSysDictData, 0)
	count, err := common.DB.
		Where("label LIKE ? AND status LIKE ?", "%"+params.Label+"%", "%"+params.Status+"%").
		Where("pid = ?", params.Pid).Count(dictData)
	if err != nil {
		return nil, 0, err
	}
	if err := common.DB.
		Where("label LIKE ? AND status LIKE ?", "%"+params.Label+"%", "%"+params.Status+"%").
		Where("pid = ?", params.Pid).
		Limit(*params.PageSize, (*params.PageNum-1)*(*params.PageSize)).
		Find(&dictDatas); err != nil {
		return nil, 0, err
	}
	return dictDatas, count, nil
}

// 增加字典数据
func (*DictService) AddDictData(params *request.DictDataAdd) error {
	dictData := &entity.AdminSysDictData{
		Pid:    *params.Pid,
		Label:  *params.Label,
		Value:  *params.Value,
		Status: params.Status,
	}
	// 插入角色表
	if _, err := common.DB.Insert(dictData); err != nil {
		return err
	}
	return nil
}

// 获取字典数据信息
func (*DictService) GetDictDataInfo(id *int) (*entity.AdminSysDictData, error) {
	dictData := new(entity.AdminSysDictData)
	if ok, err := common.DB.Where("id = ?", id).Get(dictData); !ok {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("获取字典数据信息失败")
	}
	return dictData, nil
}

// 更新字典数据
func (*DictService) UpdateDictData(params *request.DictDataUpdate) error {
	dictData := &entity.AdminSysDictData{
		Pid:    *params.Pid,
		Label:  *params.Label,
		Value:  *params.Value,
		Status: params.Status,
	}
	if _, err := common.DB.Where("id = ?", params.Id).AllCols().Update(dictData); err != nil {
		return err
	}
	return nil
}

// 删除字典数据
func (*DictService) DelDictData(params *request.DictDataDel) error {
	dictData := new(entity.AdminSysDictData)
	// 删除类型下的数据
	if _, err := common.DB.In("id", params.Ids).Delete(dictData); err != nil {
		return err
	}
	return nil
}

// 根据类型获取数据
func (*DictService) GetTypeData(dictType string) ([]entity.AdminSysDictData, error) {
	dictData := make([]entity.AdminSysDictData, 0)
	if err := common.DB.Table("admin_sys_dict_data").
		Alias("data").
		Join("INNER", []string{"admin_sys_dict_type", "type"}, "data.pid = type.id").
		Where("type = ?", dictType).
		Find(&dictData); err != nil {
		return nil, err
	}
	return dictData, nil
}

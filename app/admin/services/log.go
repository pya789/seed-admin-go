package services

import (
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/response"
	"seed-admin/common"
)

type LogService struct{}

// 增加日志
func (*LogService) AddLog(log *entity.AdminSysLog) error {
	if _, err := common.DB.Insert(log); err != nil {
		return err
	}
	return nil
}

// 删除日志
func (*LogService) DelLog(params *request.LogDel) error {
	// 删除用户
	if _, err := common.DB.Table("admin_sys_log").In("id", params.Ids).Delete(); err != nil {
		return err
	}
	return nil
}

// 获取全部日志
func (*LogService) GetAllLog(params *request.LogList) ([]response.LogList, int64, error) {
	logs := make([]response.LogList, 0)
	count, err := common.DB.SQL(`
	SELECT count(l.id) FROM 
		admin_sys_log AS l 
	LEFT JOIN 
		admin_sys_user AS u
	ON 
		l.user_id = u.id 
	WHERE
		l.method LIKE ?
	AND
		l.action LIKE ?
	AND
		l.ip LIKE ?`, "%"+params.Method+"%", "%"+params.Action+"%", "%"+params.Ip+"%").
		Count()
	if err != nil {
		return nil, 0, err
	}
	if err := common.DB.SQL(`
	SELECT * FROM 
		admin_sys_log AS l 
	LEFT JOIN 
		admin_sys_user AS u
	ON 
		l.user_id = u.id 
	WHERE
		l.method LIKE ?
	AND
		l.action LIKE ?
	AND
		l.ip LIKE ?
	ORDER BY l.id DESC
	LIMIT ?,?`, "%"+params.Method+"%", "%"+params.Action+"%", "%"+params.Ip+"%", (*params.PageNum-1)*(*params.PageSize), *params.PageSize).
		Find(&logs); err != nil {
		return nil, 0, err
	}
	return logs, count, nil
}

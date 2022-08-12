package app

import (
	"seed-admin/app/admin/entity"
	admin "seed-admin/app/admin/routers"
	api "seed-admin/app/api/routers"
	"seed-admin/common"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	admin.New(&r.RouterGroup)
	api.New(&r.RouterGroup)
	dataTableSync()
}

// 需要同步到数据库的实体
// 只允许增量同步,只允许在开发模式模式下使用
// 如有需要自行添加所需同步的表...
func dataTableSync() {
	if gin.Mode() == gin.DebugMode {
		err := common.DB.Sync(
			new(entity.AdminSysUser),
			new(entity.AdminSysUserRole),
			new(entity.AdminSysRole),
			new(entity.AdminSysMenu),
			new(entity.AdminSysRoleMenu),
			new(entity.AdminSysDept),
			new(entity.AdminSysRoleDept),
			new(entity.AdminSysLog),
			new(entity.AdminSysDictType),
			new(entity.AdminSysDictData),
			new(entity.AdminUploads),
			new(entity.AdminUploadsType),
		)
		if err != nil {
			panic(err.Error())
		}
	}
}

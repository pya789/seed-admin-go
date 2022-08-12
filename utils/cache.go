package utils

import (
	"seed-admin/common"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/goutil/jsonutil"
)

// 过期时间
const TTL = time.Hour * 24

// 前缀
const PATTERN = "auth"

type AuthCache struct{}

func (*AuthCache) SetRoleIds(userId int, roleIds []int) {
	key := PATTERN + strconv.Itoa(userId)
	value, _ := jsonutil.Encode(roleIds)
	common.Redis.HMSet(key, map[string]any{
		"roleIds": value,
	}).Result()
	common.Redis.Expire(key, TTL).Result()
}
func (*AuthCache) GetRoleIds(userId int) ([]int, error) {
	key := PATTERN + strconv.Itoa(userId)
	// redis里取出权限并解码
	roleIds := make([]int, 0)
	roleIdsStr, err := common.Redis.HGet(key, "roleIds").Result()
	if err != nil {
		return nil, err
	}
	jsonutil.DecodeString(roleIdsStr, &roleIds)
	return roleIds, nil
}

// 设置角色权限
func (*AuthCache) SetRoleLabels(userId int, roleLabels []string) {
	key := PATTERN + strconv.Itoa(userId)
	value, _ := jsonutil.Encode(roleLabels)
	common.Redis.HMSet(key, map[string]any{
		"roleLabels": value,
	}).Result()
	common.Redis.Expire(key, TTL).Result()
}
func (*AuthCache) GetRoleLabels(userId int) ([]string, error) {
	key := PATTERN + strconv.Itoa(userId)
	// redis里取出权限并解码
	roleLabels := make([]string, 0)
	roleIdsStr, err := common.Redis.HGet(key, "roleLabels").Result()
	if err != nil {
		return nil, err
	}
	jsonutil.DecodeString(roleIdsStr, &roleLabels)
	return roleLabels, nil
}

// 设置菜单权限
func (*AuthCache) SetMenuPerms(userId int, menuPerms []string) {
	key := PATTERN + strconv.Itoa(userId)
	value, _ := jsonutil.Encode(menuPerms)
	common.Redis.HMSet(key, map[string]any{
		"menuPerms": value,
	}).Result()
	common.Redis.Expire(key, TTL).Result()
}
func (*AuthCache) GetMenuPerms(userId int) ([]string, error) {
	key := PATTERN + strconv.Itoa(userId)
	// redis里取出权限并解码
	menuPerms := make([]string, 0)
	roleIdsStr, err := common.Redis.HGet(key, "menuPerms").Result()
	if err != nil {
		return nil, err
	}
	jsonutil.DecodeString(roleIdsStr, &menuPerms)
	return menuPerms, nil
}

func (*AuthCache) Del(userIds []int) error {
	keyStrs := make([]string, 0, len(userIds))
	for _, item := range userIds {
		keyStrs = append(keyStrs, PATTERN+strconv.Itoa(item))
	}
	if _, err := common.Redis.Del(keyStrs...).Result(); err != nil {
		return err
	}
	return nil
}

// 获取缓存ID
func getIds() ([]int, error) {
	// 拿到缓存所有ID
	data, err := common.Redis.Keys(PATTERN + "*").Result()
	if err != nil {
		return nil, err
	}
	// 去除PATTERN转换成int切片
	ids := make([]int, 0, len(data))
	for _, idstr := range data {
		idNum, _ := strconv.Atoi(strings.TrimPrefix(idstr, PATTERN))
		ids = append(ids, idNum)
	}
	return ids, nil
}

// 更新所有角色权限
func (authCache *AuthCache) UpdateAllRole() error {
	ids, err := getIds()
	if err != nil {
		return err
	}
	// 查询用户拥有的角色ID和角色label
	roleMap := make([]map[string]any, 0)
	if err = common.DB.Table("admin_sys_user_role").Alias("ur").
		Join("INNER", []string{"admin_sys_role", "r"}, "ur.role_id = r.id").
		In("user_id", ids).Cols("ur.user_id", "r.id", "r.label").
		Find(&roleMap); err != nil {
		return err
	}
	// 通过userId进行分组
	roleIdsMap := map[int][]int{}
	roleLabelsMap := map[int][]string{}
	for _, role := range roleMap {
		key := int(role["user_id"].(int32))
		roleIdsMap[key] = append(roleIdsMap[key], int(role["id"].(int32)))
		roleLabelsMap[key] = append(roleLabelsMap[key], role["label"].(string))
	}
	// 修改缓存的角色id
	delIds := make([]int, 0)
	for userId, roleIds := range roleIdsMap {
		authCache.SetRoleIds(userId, roleIds)
		delIds = SliceDelete(ids, userId)
		ids = ids[:len(ids)-1]
	}

	authCache.Del(delIds)
	// 修改缓存的角色label
	for userId, labels := range roleLabelsMap {
		authCache.SetRoleLabels(userId, labels)
	}
	return nil
}

// 更新所有菜单接口权限
func (authCache *AuthCache) UpdateAllPerm() error {
	ids, err := getIds()
	if err != nil {
		return err
	}
	// 查询用户拥有的菜单权限
	menusMap := make([]map[string]any, 0)
	if err = common.DB.Table("admin_sys_menu").Alias("m").
		Join("INNER", []string{"admin_sys_role_menu", "rm"}, "rm.menu_id = m.id").
		Join("INNER", []string{"admin_sys_user_role", "ur"}, "ur.role_id = rm.role_id").
		Where(`m.perms IS NOT NULL AND m.perms != "" AND m.status = 0`).
		In("ur.user_id", ids).
		Distinct("ur.user_id", "m.perms").
		Find(&menusMap); err != nil {
		return err
	}
	// 通过userId进行分组
	permsMap := map[int][]string{}
	for _, menu := range menusMap {
		key := int(menu["user_id"].(int32))
		permsMap[key] = append(permsMap[key], menu["perms"].(string))
	}

	// 修改缓存的菜单权限
	delIds := make([]int, 0)
	for userId, perms := range permsMap {
		authCache.SetMenuPerms(userId, perms)
		delIds = SliceDelete(ids, userId)
		ids = ids[:len(ids)-1]
	}
	authCache.Del(delIds)
	return nil
}

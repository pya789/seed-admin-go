<p align="center" style="margin:30px 0">
  <a href="https://admin.seed-app.com/" target="blank"><img src="https://code-admin-1258623672.cos.ap-beijing.myqcloud.com/logo.png" width="120" alt="seed logo" /></a>
</p>
<h1 align="center" style="font-weight: 600;font-size:36px;">Seed - Admin</h1>
<h4 align="center" style="font-weight: 200;">基于Go/C#(开发中)/Vue的前后端分离后台权限管理系统 </h4>
<div align="center" style="margin-top:10px;">
	<a href="https://github.com/seed-app/seed-admin-vue"><img src="https://img.shields.io/github/stars/seed-app/seed-admin-vue"></a>
	<a href="https://github.com/seed-app/seed-admin-vue/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue"></a>
  <a href="https://github.com/seed-app/seed-admin-vue"><img src="https://img.shields.io/badge/version-v1.0.0-blue"></a>
</div>

## 项目简介
seed-admin是一个纯粹的、完全开源的高可维护权限管理系统，减少了依赖，设计简单化，减少二次开发难度，提高可定制性。支持docker/常规上传部署。
- 前端采用vue3、naive-ui... 支持template和jsx两种方式开发亦可混合使用。
- 后端提供Go/C#两种解决方案可根据自身对语言的熟悉程度择其一使用
- Go方案采用gin、xorm、mysql、redis、jwt
- C#方案正在开发中...
- 即时性动态权限菜单 权限修改后用户无需刷新即可生效
- 高效率开发 前端封装CURD组件仅需填写json即可生成CURD页面

## 默认功能
- 仪表盘：好像并没有没什么用...有数据展示需求请自由发挥
- 用户管理：提供部门配置和用户权限与信息设置。
- 部门系统(内嵌在用户管理中)：配置组织机构(公司，部门，小组等)。
- 角色管理：控制角色菜单权限、按钮权限、部门数据权限等。
- 菜单管理：配置系统菜单、操作权限标识、按钮权限标识等。
- 操作日志：负责记录系统登录日志、操作日志和查询。
- 系统监控：负责监控系统CPU、内存、硬盘、服务等相关信息。
- 字典管理：负责对数据库字段值做字典解析。
- 个人管理：编辑修改个人信息。

为了维持精简性前端echarts等均无引入，后端工作流等业务组件均无开发计划有需求请自行fork开发

## 技术栈
- 前端：**`vue3` `vuex` `vue-Router` `naive-ui` `...`**
- Go后端：**`gin` `xorm` `mysql` `redis` `...`**
- C#后端：**`开发中...`**
## 项目演示
演示地址：  
[https://admin.seed-app.com](https://admin.seed-app.com)

- 账户：admin 
- 密码：123123

> 提示：演示系统数据库每10分钟会自动还原一次

文档地址：  
[https://www.seed-app.com/admin](https://www.seed-app.com/admin/)

<a href="https://admin.seed-app.com/" target="blank"><img src="https://code-admin-1258623672.cos.ap-beijing.myqcloud.com/demo.png"  alt="seed home" /></a>

## 项目关联仓库
Vue3前端：[https://github.com/seed-app/seed-admin-vue](https://github.com/seed-app/seed-admin-vue)
Go服务端：[https://github.com/seed-app/seed-admin-go](https://github.com/seed-app/seed-admin-go)  
C#服务端：[https://github.com/seed-app/seed-admin-dotnet](https://github.com/seed-app/seed-admin-dotnet) (开发中...)

## 学(kai)习(che)交(mo)流(yu)群
企鹅1群：8455822
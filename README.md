# Ovo Gaming 目录结构

项目已经整理为三个顶层模块：

- `frontend-module/`
  - `admin-web/`：管理后台，`Vue 3 + Vben + Ant Design Vue`

- `backend-module/`
  - `server/`：Go 后端 API 和数据库迁移
  - `docs/`：设计文档和计划

- `miniprogram-module/`
  - 微信小程序源码
  - 这里包含 `app.json`、`pages/`、`components/`、`assets/`

启动建议：

1. 小程序请在微信开发者工具中打开 `miniprogram-module/`
2. 后台请在 `frontend-module/admin-web/` 中运行
3. 后端请在 `backend-module/server/` 中运行

# FengFengZhidaoDoc

这是一个基于 **Go + Gin + GORM + DDD** 架构的动态文档系统
目标是实现类似 [fengfengzhidao](https://docs.fengfengzhidao.com/) 的文档站，支持：
- 书籍（Book）管理
- 章节（Chapter）管理
- 文档（Doc）管理
- 树形结构展示
- Markdown 渲染

## 项目结构 (DDD 风格)

### 核心目录 (`internal/`)

| 层级 | 目录名 | 主要职责 |
| :--- | :--- | :--- |
| 领域层 | `domain/` | 包含实体、值对象、领域服务，是业务核心 |
| 应用层 | `application/` | 包含业务用例服务，协调领域对象完成特定任务 |
| 基础设施层 | `infrastructure/` | 实现数据库、缓存、消息队列等具体技术细节 |
| 接口层 | `interfaces/` | 提供 HTTP、RPC、API 等对外接口 |

### 启动入口 (`cmd/`)

| 目录名 | 说明 |
| :--- | :--- |
| `server/` | 应用程序的启动和服务初始化入口 |
## 快速启动


```bash
1. 克隆项目
git clone https://github.com/Alitadie/fengfengzhidaodoc.git
cd fengfengzhidaodoc

2. 安装依赖

go mod tidy
3. 运行服务

go run cmd/server/main.go
4. api示例


获取整本书结构：
POST /api/books
{
  "title": "Go语言入门"
}
创建书籍：
POST /api/books/:book_id/chapters
{
  "title": "第一章 基础语法"
}
创建章节：
POST /api/books/:book_id/chapters
{
  "title": "第一章 基础语法"
}
```

# GVM - Go Version Manager

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

GVM是一个Go语言版本管理工具，类似于Node.js的nvm，rust的cargo,作者在使用了众多包管理器后集百家之所长开发了本工具。它允许开发者在同一系统上安装、管理和切换多个Go版本，非常适合需要在不同项目中使用不同Go版本的开发者。

## 功能特点

- 安装特定版本的Go
- 在多个已安装的Go版本之间切换
- 卸载不再需要的版本
- 列出本地已安装和远程可用的Go版本
- 使用指定Go版本创建新项目
- 配置管理（镜像源等）

## 安装

### 前提条件

- 操作系统：支持macOS、Linux和Windows
- Go 1.24.0或更高版本
- 基本的命令行知识

### 安装步骤

```bash
# 克隆仓库
git clone https://github.com/code-innovator-zyx/gvm.git

# 进入项目目录
cd gvm

# 编译安装
go install
```

或者直接使用go install命令安装：

```bash
go install github.com/code-innovator-zyx/gvm@latest
```

## 使用方法

### 列出Go版本

```bash
# 列出本地已安装的Go版本
gvm list

# 列出远程可用的Go版本
gvm list -r

# 列出特定类型的Go版本（稳定版、非稳定版或归档版）
gvm list -t stable
gvm list -t unstable
gvm list -t archived
```

### 安装Go版本

```bash
# 安装特定版本的Go
gvm install go1.21
```

### 切换Go版本

```bash
# 切换到特定版本的Go
gvm use go1.21
```

### 卸载Go版本

```bash
# 卸载特定版本的Go
gvm uninstall go1.21
```

### 创建新项目

```bash
# 使用当前活动的Go版本创建新项目
gvm new myproject
```

### 配置管理

```bash
# 列出所有配置
gvm config list

# 获取特定配置
gvm config get mirror

# 设置配置
gvm config set mirror https://golang.google.cn/dl/

# 删除配置
gvm config unset custom_setting
```

## 命令参考

| 命令 | 描述 |
|------|------|
| `gvm list` | 列出Go版本 |
| `gvm install` | 安装Go版本 |
| `gvm use` | 切换到特定Go版本 |
| `gvm uninstall` | 卸载Go版本 |
| `gvm new` | 创建新Go项目 |
| `gvm config` | 管理GVM配置 |

更详细的命令说明请参考[命令文档](docs/cli/gvm.md)。

## 项目结构

```
├── cmd/           # 命令行工具实现
├── docs/          # 文档
│   └── cli/       # 命令行文档
├── internal/      # 内部包
│   ├── consts/    # 常量定义
│   ├── registry/  # 版本注册表
│   ├── version/   # 版本管理
│   └── utils/     # 工具函数
└── pkg/           # 公共包
```

## 贡献

欢迎贡献代码、报告问题或提出改进建议！请遵循以下步骤：

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

本项目采用MIT许可证 - 详情请参阅[LICENSE](LICENSE)文件。

## 联系方式

如有任何问题或建议，请通过以下方式联系我们：

- 项目维护者：[mortal](1003941268@qq.com)
- GitHub Issues：[https://github.com/code-innovator-zyx/gvm/issues](https://github.com/code-innovator-zyx/gvm/issues)
# git-toolkit

> 人类懒惰的本性和不满足的本性是驱使科技发展的源泉......

本工具集包含几个部分，自定义命令，`Hook`脚本，以及配置模板

## 安装

**使用curl**

```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/tonydeng/git-toolkit/golang/installer.sh)"
```

**使用wget**

```bash
bash -c "$(wget https://raw.githubusercontent.com/tonydeng/git-toolkit/golang/installer.sh -O -)"
```

## 介绍

### 自定义命令

| 命令 | 描述 |
| -- | -- |
| ci | 提供交互式`git commit`的命令，用于定制统一的`commit message` |
| feat | 接受一个字符串，并创建一个`feat`分支，分支名称格式为 `feat/xxx` |
| fix | 创建`fix`分支 |
| docs | 创建`docs`支 |
| style | 创建`style`分支 |
| refactor | 创建`refactor`分支 |
| chore | 创建`chore`分支 |
| perf | 创建`perf`分支 |
| hotfix | 创建`hotfix`分支(通常用于对`master`紧急修复) |


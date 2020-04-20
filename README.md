# git-toolkit

> 人类懒惰的本性和不满足的本性是驱使科技发展的源泉......

本工具集包含几个部分，自定义命令，`Hook`脚本，以及配置模板

## 下载 & 安装

### 下载

- [macOS](https://git-open.qianxin-inc.cn/dengtao/git-toolkit/raw/master/dist/git-toolkit_darwin_amd64)
- [Linux 64](https://git-open.qianxin-inc.cn/dengtao/git-toolkit/raw/master/dist/git-toolkit_linux_amd64)
- [Linux 32](https://git-open.qianxin-inc.cn/dengtao/git-toolkit/raw/master/dist/git-toolkit_linux_386)

### 查看版本

```bash
./git-toolkit version
```

当你通过上面命令查看到如下信息时，说明下载成功。

```bash
 ██████╗ ██╗████████╗    ████████╗ ██████╗  ██████╗ ██╗     ██╗  ██╗██╗████████╗
██╔════╝ ██║╚══██╔══╝    ╚══██╔══╝██╔═══██╗██╔═══██╗██║     ██║ ██╔╝██║╚══██╔══╝
██║  ███╗██║   ██║          ██║   ██║   ██║██║   ██║██║     █████╔╝ ██║   ██║   
██║   ██║██║   ██║          ██║   ██║   ██║██║   ██║██║     ██╔═██╗ ██║   ██║   
╚██████╔╝██║   ██║          ██║   ╚██████╔╝╚██████╔╝███████╗██║  ██╗██║   ██║   
 ╚═════╝ ╚═╝   ╚═╝          ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝   ╚═╝


Name: git-toolkit
Version: v2.0.0
Author: Tony Deng <wolf.deng@gmail.com>
Arch: darwin/amd64
BuildTime: 2020-04-07 17:51:51
```

### 安装

```bash
./git-toolkit install
```

当看到如下信息，说明安装成功

```bash
👉 remove /Users/tonydeng/git-toolkit
📥 mkdir /Users/tonydeng/git-toolkit/hooks
📥 copy file /Users/tonydeng/git-toolkit/git-toolkit
📥 install symbolic /usr/local/bin/git-ci
📥 install symbolic /usr/local/bin/git-feat
📥 install symbolic /usr/local/bin/git-fix
📥 install symbolic /usr/local/bin/git-docs
📥 install symbolic /usr/local/bin/git-style
📥 install symbolic /usr/local/bin/git-refactor
📥 install symbolic /usr/local/bin/git-test
📥 install symbolic /usr/local/bin/git-chore
📥 install symbolic /usr/local/bin/git-pref
📥 install symbolic /usr/local/bin/git-hotfix
📥 config set core.hooksPath /Users/tonydeng/git-toolkit/hooks
```

## 使用介绍

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


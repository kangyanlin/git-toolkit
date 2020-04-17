#!/bin/bash

# 初始化
function init() {
    REPO_NAME="git-toolkit"

    if [[ -z "$REPO_HOME" ]]
    then
        REPO_HOME="https://github.com/tonydeng/git-toolkit"
    fi

    if [[ -z "$BRANCH" ]]
    then
        BRANCH="github/golang"
        # BRANCH = "master"
    fi

    PRODUCTION=""

    if [[ "$(uname -s)"=="Drawin" ]] ; then
        PRODUCTION=$REPO_NAME"_darwin_amd64"
    elif [[ "$(uname -s)"=="Linux" ]] ; then
        # TODO要判断是否为i386的系统
        if [[ "$(uname -a | egrep -i i386)" ]] ; then
            PRODUCTION=$REPO_HOME"_linux_386"
        else
            PRODUCTION=$REPO_HOME"_linux_amd64"
        fi
    fi
}

# 下载并重命名
function download() {
    DOWNLOAD_URL="$REPO_HOME/raw/$BRANCH/dist/$PRODUCTION"
    echo "################### download  $DOWNLOAD_URL ###################"

    if [[ ! -z "$PRODUCTION" ]] ; then
        rm "./$PRODUCTION" # 删除当前目录已经下载的安装包    
        if command -v wget >/dev/null 2>&1; then
            echo "wget -c -t 0  $DOWNLOAD_URL"
            wget -c -t 0 "$DOWNLOAD_URL"
        elif command -v curl >/dev/null 2>&1; then
            echo "curl -fsSL -C - -O $DOWNLOAD_URL"
            curl -fsSL -C - -O "$DOWNLOAD_URL"
        fi
    fi
}

# 安装
function install() {
    echo ""
    echo "################### install $PRODUCTION ###################"
    if [[ ! -z "$PRODUCTION" && -f "./$PRODUCTION" ]] ; then
        "./$PRODUCTION install"
    fi
}

init
download
install
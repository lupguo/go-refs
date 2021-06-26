#!/bin/bash

export PATH

PROJECT_HOME='/data/projects/csig_edu_ketang_service'
PROJECT_NAME=$1
PROJECT_PATH="${PROJECT_HOME}/$PROJECT_NAME"
OLD_PWD=$(PWD)

# 项目名称
if [ "$PROJECT_NAME" == "" ]; then
  echo "[error]project name is empty, please input project name"
  exit 10
elif [ ! -d "$PROJECT_PATH" ]; then
  echo "[error]project path not exist, project path:$PROJECT_HOME/$PROJECT_NAME"
  exit 11
fi

# 分支检测
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD 2> /dev/null)
if [ "$BRANCH_NAME" == "" ]; then
    echo "[error]branch name is empty, check if it is in the .git project, \$PWD=$PWD"
    exit 20
fi

# DEV_CLOUD配置
DEV_REMOTE_SSH='luping@9.135.17.5 -p 36000'
DEV_REMOTE_SHELL="/data/projects/csig_edu_ketang_service/${PROJECT_NAME}/edu_cp.sh ${BRANCH_NAME}"

# RUN
cd $PROJECT_PATH || exit 31
git push || exit 32
ssh "$DEV_REMOTE_SSH $DEV_REMOTE_SHELL" || exit 33

# 返回之前目录
cd "$OLD_PWD" || exit 40
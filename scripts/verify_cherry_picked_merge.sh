#!/bin/bash
set -e

# 1. 基于 main 创建 temp_main
git checkout main
git pull
git checkout -b temp_main

# 2. 创建 temp_feature 并提交
git checkout -b temp_feature
echo "test $(date)" > temp_file.txt
git add temp_file.txt
git commit -m "temp_feature commit"

# 3. 回到 temp_main，新建 temp_target
git checkout temp_main
git checkout -b temp_target

# 4. cherry-pick temp_feature 的 commit
FEATURE_COMMIT=$(git rev-parse temp_feature)
git cherry-pick $FEATURE_COMMIT

# 5. 检查 merged 状态
echo "已合并到 temp_target 的分支："
git branch --merged | grep temp_

echo "未合并到 temp_target 的分支："
git branch --no-merged | grep temp_

# 6. 清理
read -p "按回车键清理所有 temp_ 分支，或 Ctrl+C 取消..." dummy
git checkout main
git branch | grep temp_ | xargs git branch -D
rm -f temp_file.txt

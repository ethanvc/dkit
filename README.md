安全的删除本地分支。当满足下面的条件时，删除本地分支：
1. 本地分支已经合并到指定的分支列表中，比如origin/master。


git cherry -v dst src
检查src的提交是否已经存在于dst中。
git cherry -v main featurea # featurea是否已经都合并到main中

# lint
1. 当不在git仓库所在目录执行时，报错不友好。
2. 当不在git仓库的根目录下执行时，会出现找不到差异问题的情况。
安全的删除本地分支。当满足下面的条件时，删除本地分支：
1. 本地分支已经合并到指定的分支列表中，比如origin/master。


git cherry -v dst src
检查src的提交是否已经存在于dst中。
git cherry -v main featurea # featurea是否已经都合并到main中

test

# json字符串加工处理需求整理
1. 值如果为字符串，且字符串为合法的json，将其直接展开，避免对比的时候有太多转义。
2. 删除指定条件的字段。比如根据字段名称、字段全路径。
3. （先不考虑这里做，方式还不确定）文本替换。对于不同请求中可变的文本，统一替换，避免这种差异的干扰。

# JsonArrayToObject
将json的array变换成object，方便对比时，有更好的路径提示。
支持通过全路径指定需要变换的array。
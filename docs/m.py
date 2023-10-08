import os
import re

# 递归遍历目录及其子目录中的所有 .md 文件
for root, dirs, files in os.walk('.'):
    for file in files:
        if file.endswith(".md"):
            file_path = os.path.join(root, file)
            with open(file_path, "r") as f:
                content = f.readlines()

            # 查找并删除第一个一级标题
            for i in range(len(content)):
                if content[i].startswith("# "):
                    content[i] = ""
                    break

            # 将新内容写回文件
            with open(file_path, "w") as f:
                f.writelines(content)

print("转换完成！")
#!/bin/bash
# shell实现 遍历指定文件夹下的扩展名为mp4的文件  并把文件名中 [] 包裹的任何内容包括中括号本身删除 执行这个重命名操作
folder_path="/path/to/your/folder"  # 替换为你要遍历的文件夹路径

for file in "$folder_path"/*.mp4; do
    if [ -f "$file" ]; then
        new_name=$(echo "$file" | sed 's/\[[^]]*\]//g')
        mv "$file" "$new_name"
    fi
done

#!/usr/bin/env python3
import os
import shutil
import chardet


def test(filename):
    try:
        isMakeDir(filename)
        dirName = os.listdir(filename)
        print(dirName)
        # 目录下的单个文件，或者文件夹
        for single in dirName:
            innerFileNameOrDir = filename + "/" + single
            if single == 'classes' or single == 'bin'or single.startswith('.'):
                print('ignore name：' + single)
                continue
            elif single.endswith('.properties') or single.endswith('.jar') or single.endswith('.xml') \
                    or single.endswith('.jsp') or single.endswith('.jps') or single.endswith('.JPG') \
                    or single.endswith('.txt'):
                targetFileNameOrDir = "D:" + innerFileNameOrDir.split(":")[1]
                shutil.copy(innerFileNameOrDir, targetFileNameOrDir)
                continue

            if os.path.isfile(innerFileNameOrDir):
                if not single.endswith(".java") and not single.endswith(".cs") and not single.endswith(".lua"):
                    print("copy:", single)
                    targetFileNameOrDir = "D:" + innerFileNameOrDir.split(":")[1]
                    shutil.copy(innerFileNameOrDir, targetFileNameOrDir)
                    continue
                with open(innerFileNameOrDir, 'rb') as fileObj:  # UTF-8   encoding='UTF-8'
                    print("###" + innerFileNameOrDir)
                    javaContent = fileObj.read()
                    baseFileName = os.path.basename(innerFileNameOrDir)
                    singleName = baseFileName.split('.')
                    newSingleName = filename + "/" + singleName[0]
                    newSingleName = "D:" + newSingleName.split(":")[1]
                    fp = open(newSingleName, 'wb')
                    fp.write(javaContent)

            elif os.path.isdir(innerFileNameOrDir):
                test(innerFileNameOrDir)

    except UnicodeDecodeError as error:
        # print("test ",error)
        raise


def isMakeDir(innerFileNameOrDir):
    innerFileNameOrDir = "D:" + innerFileNameOrDir.split(":")[1]
    innerFileNameOrDir = innerFileNameOrDir.strip()
    innerFileNameOrDir = innerFileNameOrDir.rstrip("/")
    isExists = os.path.exists(innerFileNameOrDir)
    if not isExists:
        os.makedirs(innerFileNameOrDir)
        print('目录创建成功:', innerFileNameOrDir)


if __name__ == '__main__':
    # D:/python_rw/client/dfh/dev/hall/script/task
    # target_path = 'E:/python_rw/client'
    target_path = 'c:/temptest'
    test(target_path)

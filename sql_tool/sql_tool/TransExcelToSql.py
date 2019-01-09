# coding=utf-8
# create by R.Zhu
# time 2018/9/7
import sys

import xlrd
import re
from datetime import datetime
from xlrd import xldate_as_tuple


class transExcelToSql:
    def trans_excel(self, filename, sheetname):
        rbook = xlrd.open_workbook(filename)
        sheet = rbook.sheet_by_name(sheetname)
        sql_str = "REPLACE INTO %s" % sheetname
        sql_path = filename[0:-4]
        print("sql生成路径："+sql_path)
        file = open(sql_path, 'w', encoding='utf8')
        # python读取excel中单元格的内容返回的有5种类型
        # ctype： 0 empty, 1 string, 2 number, 3  date, 4 boolean, 5  error
        for i in range(sheet.nrows):
            row_content = []
            for j in range(sheet.ncols):
                ctype = sheet.cell(i, j).ctype  # 表格的数据类型
                cell = sheet.cell_value(i, j)
                if ctype == 2 and cell % 1 == 0:  # 如果是整形
                    cell = int(cell)
                elif ctype == 3:
                    # 转成datetime对象
                    date = datetime(*xldate_as_tuple(cell, 0))
                    cell = date.strftime('%Y:%m:%d %H:%M:%S')
                elif ctype == 4:
                    cell = True if cell == 1 else False
                row_content.append(cell)
            row_value = '('
            if i == 1:
                row_value = sql_str + row_value
                for element in row_content:
                    # print(element)
                    searchObj = re.search(r'(.*)\((.*)\)', element, re.M | re.I)
                    colum_name = searchObj.group(1)
                    # colum_name_type = searchObj.group(2)
                    # print(colum_name)
                    # print(colum_name_type)
                    row_value += ''.join(colum_name)+","
                row_value = row_value[:-1]
                row_value += ') VALUES'
            else:
                row_value = '(' + ','.join("'" + str(element) + "'" for element in row_content) + '),'

            if i == sheet.nrows - 1:
                row_value = row_value[:-1]
                row_value += ";"
            print(row_value)
            if i > 0:
                file.write(row_value+'\n')


if __name__ == '__main__':
    trans = transExcelToSql()
    tableName = "active"
    dataBaseName = "dfh7_cfg"
    if len(sys.argv) >= 2:
        dataBaseName = sys.argv[1]
    if len(sys.argv) >= 3:
        tableName = sys.argv[2]
    filename = "D:/版本/19/" + dataBaseName + "." + tableName + ".xls"
    trans.trans_excel(filename, sheetname=tableName)
# coding=utf-8
# create by R.Zhu
# time 2018/9/7
import sys

import xlwt
import pymysql


class importDataToExcel():

    def createExcel(self, database_name, table_name, path_name):
        try:
            conn = pymysql.connect(host="127.0.0.1", user="root", passwd="123qWE", db=database_name, charset="utf8")
            cur = conn.cursor()
            sql = 'SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = %s'
            # AND table_schema = '%s'
            cur.execute(sql, (table_name,))
            results = cur.fetchall()
            if len(results) == 0:
                print("该表中没有任何字段")
                return

            book = xlwt.Workbook(encoding='utf-8')

            # 创建一个样式----------------------------
            date_style_index = []
            time_style_index = []
            sheet1 = book.add_sheet(table_name)
            for index in range(0, len(results)):
                print(results[index])
                # print(results[index][0]+("("+results[index][1])+")")
                sheet1.col(index).width = 5000
                sheet1.write(0, index, results[index][2],)
                sheet1.write(1, index, results[index][0]+("("+results[index][1])+")")
                # 时间格式记录下，特殊处理存入Excel中
                if results[index][1] == "timestamp" or results[index][1] == "datetime":
                    date_style_index.append(index)
                    # print(results[index][1])
                elif results[index][1] == "time":
                    time_style_index.append(index)

            # 导入数据
            # data_sql = "select * from "+ table_name
            data_sql ="select * from %s" % table_name

            cur.execute(data_sql)
            data_results = cur.fetchall()
            # print(data_results)
            # print(len(data_results))
            # 2012-11-11 12:23:00
            date_format = xlwt.XFStyle()
            date_format.num_format_str = 'yyyy-mm-dd hh:mm:ss'
            time_format = xlwt.XFStyle()
            time_format.num_format_str = 'hh:mm:ss'
            # 如果有数据，则把数据导入到excel
            if len(data_results) > 0:
                for row_index in range(0, len(data_results)):
                    for col_index in range(0, len(data_results[row_index])):
                        # print(data_results[row_index][col_index])
                        data_value = data_results[row_index][col_index]
                        if col_index in date_style_index:
                            sheet1.write(row_index + 2, col_index, data_value, date_format)
                        elif col_index in time_style_index:
                            data_value = str(data_value)
                            sheet1.write(row_index + 2, col_index, data_value)
                        else:
                            sheet1.write(row_index+2, col_index, data_value)

                    # print("--------------")

            book.save(path_name)
            print("创建excel成功："+path_name)
        except pymysql.Error as e:
            print("Mysql Error %d: %s" % (e.args[0], e.args[1]))


if __name__ == '__main__':
    importExcel = importDataToExcel()
    tableName = "active"
    dataBaseName = "dfh7_cfg"
    if len(sys.argv) >= 2:
        dataBaseName = sys.argv[1]
    if len(sys.argv) >= 3:
        tableName = sys.argv[2]
    path = "D:/版本/19/" + dataBaseName + "." + tableName + ".xls"
    importExcel.createExcel(dataBaseName, tableName, path)



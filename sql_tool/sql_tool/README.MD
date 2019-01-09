python -m pip install pymysql
python -m pip install xlwt

1、py ImportDataToExcel.py dfh7_cfg active

argv[1]:数据库，argv[2]:表名
生成一个dfh7_cfg.avtive.xls Excel表格
eg：
id	    活动查找标识	    开始时间	                结束时间	            活动名称	        活动标题	        活动内容
id(int)	logo(varchar)	startTime(timestamp)	endTime(timestamp)	active(varchar)	title(varchar)	content(varchar)

如果表里有数据，也会自动生成出来

2、py TransExcelToSql.py dfh7_cfg active
argv[1]:数据库，argv[2]:表名
根据上面生成的表格，导成sql语句
REPLACE INTO active(id,logo,startTime,endTime,active,title,content) VALUES.........

3、
一般情况常用的字段应该都没问题，有的特殊字段读取到Excel中格式可能会变掉，如时间类型的，这个已经特殊处理了
    mysql中时间格式：
        timestamp:时间格式为 类似 2012-11-11 12:23:00 ,默认值为当前时间
        datetime时间格式类似2012-11-11 12:23:00

        不支持time:
        time时间格式类似12:23:00 默认值为null（没找到解决方法，写入excel成功，但是转化sql有问题）

如果有其他特殊的字段，没加的，后面发现再加.....
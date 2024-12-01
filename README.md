预计初步实现：

-   可以上传 excel 表格导入信息
-   可以让用户填写表格
-   动态表
-   固定表

实现思路：

-   数据库的表

公有表：存放公有字段，必须是所有老师都有的，并且大概率以后不会改变的(如果大概率会变，应该放入动态表)

-   TEACHER_ID 表示是属于哪个老师的
-   SHEET_ID 表示来自哪一张表
-   ...

动态表：存放老师特有的字段，具有 SHEET_ID 的属性，老师需要共有字段以外的字段时则创建一个动态表。老师可以删除或者修改这个表，老师创建动态表时，会在这里添加记录

TEACHER_TO_SHEET：存放老师与表的关系

-   TEACHER_ID
-   SHEET_ID
-   SHEET_NAME

SHEET_TO_SHEET：存放表与表之间的关系。对于可能多项的数据，新建表并指向原动态表。这张表里记录(如果数据量不大，也可以作为字符串存储，发送给前端解析成数组)

-   FATHER_ID
-   CHILD_ID

共有表可以使用 MySQL, 动态表可以使用开源的动态数据库 MongoDB, 不需要预设列名并且天然支持动态字段的增查删改。如果为了更加的灵活和简单实现，可考虑只使用动态表。

前端页面：

-   登陆/注册/忘记密码
-   上传表格：使用开源的第三方库 xlsx 解析表格，并将数据发送给后端
-   查看表格：从后端接收数据，并允许用户下载

前端可考虑使用 electron(可能比较困难)做成跨平台应用(windows/mac/linux), 或者使用 PWA(配置简单，可能部分较老的浏览器和 ios/mac 不支持)实现类似应用的效果

后端：

-   登录/注册/忘记密码接口
-   表格相关接口：  
    获取所有表：返回所有表的 ID 和表名  
    接收特定的数据结构(由前端处理好的数据)并将其中一部分存入固定表，另一部分存入老师专有的动态表(老师与老师之间的数据名可能不同，如果老师还未创建动态表，则先提醒老师是否创建)。老师可选择是否全部存入动态表  
    增加列：传入 SHEET_ID 与列名  
    增加行：传入 SHEET_ID 与特定的数据结构，与接收 excel 数据类似  
    查看表：传入 SHEET_ID，根据 SHEET_ID 查看  
    删除表：传入 SHEET_ID，根据 SHEET_ID 删除  
    修改表：传入 SHEET_ID 与需要修改的行与列以及新值  

前端：vue3   
后端：go/gin                
               

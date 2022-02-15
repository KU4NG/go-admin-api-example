package initialize

import (
    "fmt"
    "go-admin-api-example/src/common"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "log"
    "strings"
    "time"
)

// Mysql 数据库连接初始化
func Mysql() {
    // 用于打印的连接
    dsnLog := fmt.Sprintf("%s:******@tcp(%s:%d)/%s?%s&charset=%s&collation=%s",
        common.Conf.Mysql.Username,
        common.Conf.Mysql.Host,
        common.Conf.Mysql.Port,
        common.Conf.Mysql.Database,
        common.Conf.Mysql.Query,
        common.Conf.Mysql.Charset,
        common.Conf.Mysql.Collation,
    )

    // 实际用于连接数据库的连接串
    dsn := strings.Replace(dsnLog, "******", common.Conf.Mysql.Password, 1)

    // 打开数据库链接
    db, err := gorm.Open(mysql.New(mysql.Config{
        DSN:               dsn, // 连接信息
        DefaultStringSize: 170, // string 类型字段的默认长度
    }), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            TablePrefix:   fmt.Sprintf(common.Conf.Mysql.TablePrefix + "_"), // 表名前缀
            SingularTable: true,                                             // 表名单数
        },
        DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键
        QueryFields:                              true, // 解决查询全部字段可能不走索引的问题
    })

    // 错误退出
    if err != nil {
        message := fmt.Sprintf("数据库连接失败：%s", dsnLog)
        common.Log.Error(message)
        common.Log.Error(err.Error())
        panic(message)
    }

    // 设置数据库连接池
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(common.Conf.Mysql.MaxIdleConns)                                // 空闲连接池中连接的最大数量
    sqlDB.SetMaxOpenConns(common.Conf.Mysql.MaxOpenConns)                                // 最大连接数量
    sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(common.Conf.Mysql.MaxIdleTime)) // 连接最大可复用时间

    // 获取数据库连接
    common.DB = db

    log.Println(fmt.Sprintf("数据库连接成功：%s", dsnLog))
}

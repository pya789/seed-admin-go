package xorm

import (
	"database/sql"
	"fmt"
	goLog "log"
	"os"
	"seed-admin/common"
	"seed-admin/core/zap/lumberjack"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var wirteSqlTask sync.WaitGroup

func AddXorm() *xorm.Engine {
	if common.CONFIG.String("mysql.database") == "" {
		panic("配置mysql.database为空,请检查config.toml配置文件")
	}
	db, err := newXorm()
	if err != nil {
		panic(err.Error())
	}
	if err = db.Ping(); err != nil {
		if strings.Contains(err.Error(), "Unknown database") {
			return initDb()
		}
		panic(err.Error())
	}
	return db
}
func newXorm() (*xorm.Engine, error) {
	var config = common.CONFIG.StringMap("mysql")
	// 组装连接字符串
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", config["user"], config["password"], config["server"], config["port"], config["database"], config["config"])
	db, err := xorm.NewEngine("mysql", dsn)
	if common.CONFIG.String("xorm_log.outType") == "file" {
		db.SetLogger(log.NewSimpleLogger(&lumberjack.Logger{
			Filename:   "./" + common.CONFIG.String("xorm_log.director") + "/db.log",
			MaxSize:    common.CONFIG.Int("xorm_log.maxSize"),
			MaxBackups: common.CONFIG.Int("xorm_log.maxBackups"),
			MaxAge:     common.CONFIG.Int("xorm_log.maxAge"),
			Compress:   common.CONFIG.Bool("xorm_log.compress"),
		}))
	} else {
		db.SetLogger(log.NewSimpleLogger(os.Stdout))
	}
	db.ShowSQL(common.CONFIG.Bool("xorm_log.showSql"))
	db.SetLogLevel(newLogger())
	db.SetMaxIdleConns(common.CONFIG.Int("mysql.maxIdleConns"))
	db.SetMaxOpenConns(common.CONFIG.Int("mysql.maxOpenConns"))
	return db, err
}

func initDb() *xorm.Engine {
	var config = common.CONFIG.StringMap("mysql")
	goLog.Println("未发现数据库,将自动创建...")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/", config["user"], config["password"], config["server"], config["port"])
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			common.LOG.Error(err.Error())
		}
	}(db)
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	createSql := fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", config["database"])
	_, err = db.Exec(createSql)
	if err != nil {
		panic(err.Error())
	}
	goLog.Printf("数据库 %v 创建完毕...\n", config["database"])
	xormDb, err := newXorm()
	if err != nil {
		panic(err.Error())
	}
	wirteSql(xormDb)
	return xormDb
}

// 加载文件
func loadFiles(path string) []string {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	fileInfo, _ := file.ReadDir(-1)
	files := []string{}
	for _, item := range fileInfo {
		files = append(files, item.Name())
	}
	return files
}

// 导入数据表和数据
func wirteSql(db *xorm.Engine) {
	db.ShowSQL(false)
	goLog.Println("开始导入数据表与数据...")
	path := "sql/"
	files := loadFiles(path)
	for _, file := range files {
		wirteSqlTask.Add(1)
		go importFile(db, path+file)
	}
	wirteSqlTask.Wait()
	goLog.Println("数据表与数据导入完成...")
	db.ShowSQL(common.CONFIG.Bool("xorm_log.showSql"))
}

func importFile(db *xorm.Engine, path string) {
	_, err := db.ImportFile(path)
	if err != nil {
		goLog.Fatalln(path + "导入失败 error:" + err.Error())
	}
	wirteSqlTask.Done()
}

// 日志配置
func newLogger() log.LogLevel {
	var level log.LogLevel
	switch common.CONFIG.String("xorm_log.level") {
	case "debug":
		level = log.LOG_DEBUG
	case "info":
		level = log.LOG_INFO
	case "warn":
		level = log.LOG_WARNING
	case "error":
		level = log.LOG_ERR
	default:
		level = log.LOG_INFO
	}
	return level
}

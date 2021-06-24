package database

import (
	"errors"
	"fmt"
	"sync"
	"time"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var once sync.Once

// DbMaster ...
var DbMaster *xorm.Engine

var DbMasterMulti *xorm.Engine

// DbSlave ...
var DbSlave *xorm.Engine

func GetMysqlUrl(multiStatements bool) string {
	url := fmt.Sprintf("root:123456@tcp(49.235.47.5:3306)/d0_server?charset=utf8&parseTime=true&loc=Local&multiStatements=%v", multiStatements)
	return url
}

// GetDbMaster 获取master的db对象
func GetDbMaster(multiStatements bool) *xorm.Engine {
	if multiStatements {
		if DbMasterMulti != nil && DbMasterMulti.Ping() == nil {
			// GlobalConfig.LogInfoHandler.Println("Mysql Master Reuse!")
			return DbMasterMulti
		}
	} else {
		if DbMaster != nil && DbMaster.Ping() == nil {
			// GlobalConfig.LogInfoHandler.Println("Mysql Master Reuse!")
			return DbMaster
		}
	}

	//url := fmt.Sprintf(GlobalConfig.GlobalDbConfig.DbMaster.User+":%s@tcp(%s:%s)/kingnet_pay_cn?charset=utf8&parseTime=true&loc=Local&multiStatements=%v", GlobalConfig.GlobalDbConfig.DbMaster.Passwd, GlobalConfig.GlobalDbConfig.DbMaster.Ip, GlobalConfig.GlobalDbConfig.DbMaster.Port, multiStatements)
	url := GetMysqlUrl(multiStatements)

	var err error
	//GlobalConfig.LogInfoHandler.Println(fmt.Sprintf("DbMaster is nil,url:%s", url))

	if multiStatements {
		DbMasterMulti, err = xorm.NewEngine("mysql", url)

		if err != nil {
			//GlobalConfig.LogFatalHandler.Println("Mysql master open error:" + err.Error())
			return nil
		}

		DbMasterMulti.ShowSQL(false)
		DbMasterMulti.SetMaxIdleConns(3)
		DbMasterMulti.SetMaxOpenConns(8)
		DbMasterMulti.SetConnMaxLifetime(time.Duration(2 * time.Minute))
		return DbMasterMulti
	}

	DbMaster, err = xorm.NewEngine("mysql", url)

	if err != nil {
		//GlobalConfig.LogFatalHandler.Println("Mysql master open error:" + err.Error())
		return nil
	}

	DbMaster.ShowSQL(false)
	DbMaster.SetMaxIdleConns(3)
	DbMaster.SetMaxOpenConns(8)
	DbMaster.SetConnMaxLifetime(time.Duration(2 * time.Minute))

	return DbMaster
}

// GetDbSlave 获取slave的db对象
func GetDbSlave() *xorm.Engine {
	dbSlaveInit := func() {
		url := "root:123456@tcp(49.235.47.5:3306)/do_server?charset=utf8&parseTime=true&loc=Local"

		var err error
		//GlobalConfig.LogInfoHandler.Println(fmt.Sprintf("DbSlave is nil,url:%s", url))

		DbSlave, err = xorm.NewEngine("mysql", url)
		if err != nil {
			//GlobalConfig.LogFatalHandler.Println("Mysql slave open error:" + err.Error())
			return
		}
		DbSlave.ShowSQL(false)
		DbSlave.SetMaxIdleConns(100)
		DbSlave.SetMaxOpenConns(1000)
		DbSlave.SetConnMaxLifetime(time.Duration(2 * time.Minute))
	}

	if DbSlave != nil && DbSlave.Ping() == nil {
		// GlobalConfig.LogInfoHandler.Println("Mysql Slave Reuse!")
		return DbSlave
	}

	once.Do(dbSlaveInit)

	return DbSlave
}

// FindAllMapBySql ...
func FindAllMapBySql(sql string) (results []map[string]string, err error) {
	db := GetDbSlave()
	if db == nil {
		return nil, errors.New(" GetDbSlave is nil")
	}
	// defer db.Close()
	results, err = db.QueryString(sql)
	if err != nil {
		return nil, errors.New("QueryE：" + err.Error() + "|" + sql)
	}
	if len(results) == 0 {
		return nil, errors.New("norecord")
	}
	return results, nil
}

// FindAllMap ...
func FindAllMap(sqlOrArgs ...interface{}) (results []map[string]string, err error) {
	db := GetDbSlave()
	if db == nil {
		return nil, errors.New(" GetDbSlave is nil")
	}
	// defer db.Close()
	results, err = db.QueryString(sqlOrArgs...)
	if err != nil {
		return nil, fmt.Errorf("QueryE：%s, sqlOrArgs: %v", err.Error(), sqlOrArgs)
	}
	if len(results) == 0 {
		return nil, errors.New("norecord")
	}
	return results, nil
}

// FindOneMapBySql ...
func FindOneMapBySql(sql string) (result map[string]string, err error) {
	db := GetDbSlave()
	if db == nil {
		return nil, errors.New("gethandE：" + err.Error())
	}
	// defer db.Close()
	results, err := db.QueryString(sql)
	if err != nil {
		return nil, errors.New("QueryE：" + err.Error() + "|" + sql)
	}
	if len(results) == 0 {
		return nil, errors.New("norecord")
	}
	return results[0], nil
}

// FindOneMap ...
func FindOneMap(sqlOrArgs ...interface{}) (result map[string]string, err error) {
	db := GetDbSlave()
	if db == nil {
		return nil, errors.New("gethandE：" + err.Error())
	}

	// defer db.Close()
	results, err := db.QueryString(sqlOrArgs...)
	if err != nil {
		return nil, fmt.Errorf("QueryE：%s, sqlOrArgs: %v", err.Error(), sqlOrArgs)
	}
	if len(results) == 0 {
		return nil, errors.New("norecord")
	}
	return results[0], nil
}

// ExeSql ...
func ExeSql(sql string) (cnt int64, err error) {
	db := GetDbMaster(false)
	if db == nil {
		return 0, errors.New("GetHandE")
	}
	// defer db.Close()
	res, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	cnt, err = res.RowsAffected()
	return cnt, err
}

func ExeMultiSql(sql string) (cnt int64, err error) {
	db := GetDbMaster(true)
	if db == nil {
		return 0, errors.New("GetHandE")
	}
	// defer db.Close()
	res, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	cnt, err = res.RowsAffected()
	return cnt, err
}

// ExeSqlReturnLastId ...
func ExeSqlReturnLastId(sql string) (cnt int64, err error) {
	db := GetDbMaster(false)
	if db == nil {
		return 0, errors.New("GetHandE")
	}
	// defer db.Close()
	res, err := db.Exec(sql)
	if err != nil {
		return 0, errors.New("ExeE：" + err.Error())
	}
	cnt, err = res.LastInsertId()
	if err != nil {
		return 0, errors.New("ExeE：" + err.Error())
	}
	return cnt, nil
}

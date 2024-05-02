package sql

import (
	"fmt"
	"github.com/zhangyiming748/WhisperInDocker/constant"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"time"
)

var db *gorm.DB

func init() {
	cstLocation, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = cstLocation
}

func SetEngine() {
	file := strings.Join([]string{constant.GetRoot(), "whisper.db"}, string(os.PathSeparator))
	db, _ = gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	// 迁移 schema
	err := db.AutoMigrate(Whisper{})
	if err != nil {
		return
	}
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	//db.Delete(&product, 1)
	fmt.Println(db)
}
func GetEngine() *gorm.DB {
	return db
}

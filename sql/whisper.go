package sql

import (
	"gorm.io/gorm"
	"time"
)

type Whisper struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement;comment:'主键'"`
	Name      string `gorm:"name	;comment:'文件名'"`
	No        string `gorm:"no;comment:'字幕序号'"`
	Duration  string `gorm:"duration;comment:'时间段'"`
	Srt       string `gorm:"原文内容"`
	CreatedAt time.Time
}

func (w *Whisper) FindOneByName() *gorm.DB {
	return GetEngine().Where("name = ?", w.Name).First(&w)
}

func (w *Whisper) SetOne() *gorm.DB {
	return GetEngine().Create(&w)
}

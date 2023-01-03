package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Store interface {
	Create(value interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	DB() (*sql.DB, error)
	Raw(sql string, values ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
}

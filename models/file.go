package models

type FMeta struct {
	Id   int64  `gorm:"primaryKey"`
	Loc  string `gorm:"column:loc"`
	Size int64  `gorm:"column:size"`
	Name string `gorm:"column:name"`
	Ext  string `gorm:"column:ext"`
}

func (FMeta) TableName() string {
	return "dirs"
}

type Results map[string]FMeta

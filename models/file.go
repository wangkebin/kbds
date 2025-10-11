package models

type FMeta struct {
	Id        int64  `gorm:"primaryKey"`
	Loc       string `gorm:"column:loc"`
	Size      int64  `gorm:"column:size"`
	Name      string `gorm:"column:name"`
	Ext       string `gorm:"column:ext"`
	MachineId string `gorm:"column:machine_id"`
}

func (FMeta) TableName() string {
	return "file_info"
}

type Results map[string]FMeta

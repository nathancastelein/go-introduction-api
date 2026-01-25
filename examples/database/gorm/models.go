package gorm

type Category struct {
	ID   int32 `gorm:"primaryKey"`
	Name string
}

func (Category) TableName() string {
	return "category"
}

type Pet struct {
	ID         int32 `gorm:"primaryKey"`
	Name       string
	CategoryID *int32   `gorm:"column:category"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}

func (Pet) TableName() string {
	return "pet"
}

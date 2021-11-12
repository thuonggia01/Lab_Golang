package model

type User struct {
	ID   int64
	Name string
	// gorm.Model
	// Name  string
	// Todos []*Todo `gorm:"ForeignKey:ID;AssociationForeignKey:UserID"`
}

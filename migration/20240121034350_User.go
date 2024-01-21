package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sample_project/src/infrastructure/database"
	"gorm.io/gorm"
)

type V20240121034350User struct {
	database.BaseModel
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
}

func (V20240121034350User) TableName() string {
	return "users"
}

var V20240121034350 *gormigrate.Migration = &gormigrate.Migration{
	ID: "20240121034350_User",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&V20240121034350User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&V20240121034350User{})
	},
}

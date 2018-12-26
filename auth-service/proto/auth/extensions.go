package auth

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return scope.SetColumn("Id", uuid.String())
}

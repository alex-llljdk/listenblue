package gmodel

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id              uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Email           string    `gorm:"column:email;type:varchar(255);comment:用户邮箱;NOT NULL" json:"email"`
	Name            string    `gorm:"column:name;type:varchar(255);comment:用户姓名;NOT NULL" json:"name"`
	Password        string    `gorm:"column:password;type:varchar(255);comment:用户密码;NOT NULL" json:"password"`
	Signature       string    `gorm:"column:sign;type:varchar(255);comment:个性签名;" json:"dec"`
	Avatar          string    `gorm:"column:avatar;type:varchar(255);comment:头像;NOT NULL" json:"avatar"`
	BackgroundImage string    `gorm:"column:background_url;type:varchar(255);comment:背景图;" json:"background_url"`
	CreateTime      time.Time `gorm:"column:createtime;type:datetime(0);autoUpdateTime" json:"createtime"`
	UpdateTime      time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"updatetime"`
	DeletedTime     gorm.DeletedAt
}

func (m *User) TableName() string {
	return "user"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	return &UserModel{
		db: db,
	}
}

// FindOneByMobile
func (m *UserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	result := m.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// InsertUser
func (m *UserModel) InsertUser(ctx context.Context, user *User) error {
	result := m.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

// GetUserByID
func (m *UserModel) GetUserByID(ctx context.Context, userID uint) (*User, error) {
	var user User
	result := m.db.WithContext(ctx).Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser
func (m *UserModel) UpdateUser(ctx context.Context, user *User) error {
	// Save the updated user information in the database
	existingUser := &User{}
	err := m.db.WithContext(ctx).First(existingUser, user.Id).Error
	if err != nil {
		return err
	}
	existingUser.Name = user.Name
	existingUser.Avatar = user.Avatar
	existingUser.BackgroundImage = user.BackgroundImage
	existingUser.Signature = user.Signature

	// 保存更新后的用户信息到数据库
	err = m.db.WithContext(ctx).Save(existingUser).Error
	if err != nil {
		return err
	}

	return nil
}

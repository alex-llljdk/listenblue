package gmodel

import (
	"context"

	"gorm.io/gorm"
)

type LiveRoom struct {
	gorm.Model
	UserID     int    `gorm:"column:user_id;type:bigint(20);NOT NULL;index" json:"user_id"`
	RoomTitle  string `gorm:"column:room_title;type:varchar(255);NOT NULL" json:"room_title"`
	IsOpen     bool   `gorm:"column:is_open;NOT NULL" json:"is_open"`
	Password   string `gorm:"column:password;type:varchar(255)" json:"password"`
	IsRecord   bool   `gorm:"column:is_record;NOT NULL" json:"is_record"`
	RoomNumBer string `gorm:"column:room_number;type:varchar(255);NOT NULL" json:"room_number"`
	UserName   string `gorm:"column:user_name;type:varchar(255);NOT NULL;" json:"user_name"`
	Url        string `gorm:"column:url;type:varchar(255);NOT NULL;" json:"url"`
	CoverUrl   string `gorm:"column:cover_url;type:varchar(255);NOT NULL;" json:"cover_url"`
	Avatar     string `gorm:"column:avatar;type:varchar(255);NOT NULL;" json:"avatar"`
}

func (m *LiveRoom) TableName() string {
	return "liveroom"
}

type LiveRoomModel struct {
	db *gorm.DB
}

func NewLiveRoomModel(db *gorm.DB) *LiveRoomModel {
	if err := db.AutoMigrate(&LiveRoom{}); err != nil {
		panic(err)
	}
	return &LiveRoomModel{
		db: db,
	}
}

func (m *LiveRoomModel) InsertLiveRoom(ctx context.Context, liveroom *LiveRoom) error {
	result := m.db.WithContext(ctx).Create(liveroom)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *LiveRoomModel) FindOneByRoomNumber(ctx context.Context, room_number string) (*LiveRoom, error) {
	var liveroom LiveRoom
	result := m.db.WithContext(ctx).Where("room_number = ?", room_number).First(&liveroom)
	if result.Error != nil {
		return nil, result.Error
	}
	return &liveroom, nil
}

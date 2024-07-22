package gmodel

//所有文件上传存放到记录表中，其他接口相同
import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	Id       uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Path     string `gorm:"column:path;type:varchar(255);comment:文件路径 音视频存放原地址 文档word转pdf后存放pdf地址;NOT NULL" json:"path"`
	Title    string `gorm:"column:title;type:varchar(255);comment:文件标题;NOT NULL" json:"title"`
	UserId   uint64 `gorm:"column:user_id;type:bigint(20);comment:所属用户id;NOT NULL" json:"user_id"`
	DataType uint   `gorm:"column:data_type;type:tinyint(3);comment:数据类型 0-mp3 1-mp4 2-pdf 3-word;NOT NULL" json:"data_type"`
	Keyword  string `gorm:"column:keyword;type:text;comment:音频关键词和文档关键要点;NOT NULL" json:"keyword"`
	Summary  string `gorm:"column:summary;type:text;comment:音频和文档全文概述;NOT NULL" json:"summary"`
	Scanning string `gorm:"column:scanning;type:text;comment:音频章节速览和文档全文速度;NOT NULL" json:"scanning"`
	//下面是音视频单独字段
	InnerText string `gorm:"column:inner_text;type:text;comment:用户编辑框内容;" json:"inner_text"`
	Review    string `gorm:"column:review;type:text;comment:音频要点回顾;" json:"review"`
	TransText string `gorm:"column:trans_text;type:text;comment:音频转写原文;" json:"trans_text"`

	CreateTime  time.Time `gorm:"column:createtime;type:datetime(0);autoUpdateTime" json:"createtime"`
	UpdateTime  time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"updatetime"`
	DeletedTime gorm.DeletedAt
}

func (m *Record) TableName() string {
	return "Record"
}

type RecordModel struct {
	db *gorm.DB
}

func NewRecordModel(db *gorm.DB) *RecordModel {
	if err := db.AutoMigrate(&Record{}); err != nil {
		panic(err)
	}
	return &RecordModel{
		db: db,
	}
}

// FindOneByPath
func (m *RecordModel) FindOneByPath(ctx context.Context, path string) (*Record, error) {
	var record Record
	result := m.db.WithContext(ctx).Where("path = ?", path).First(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

func (m *RecordModel) InsertRecord(ctx context.Context, record *Record) error {
	result := m.db.WithContext(ctx).Create(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type IdAndTitle struct {
	Id    uint64
	Title string
}

func (m *RecordModel) GetListRecorById(ctx context.Context, user_id int32) ([]Record, error) {
	// var idtitles []IdAndTitle
	var record []Record
	// sql := "select from Record where user_id=?"
	// result := m.db.Raw(sql, user_id).Scan(&record)
	result := m.db.WithContext(ctx).Where("user_id = ?", user_id).Order("updatetime desc").Find(&record)
	// result := m.db.WithContext(ctx).Where("user_id = ?", user_id).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (m *RecordModel) GetRecorById(ctx context.Context, id uint64) (*Record, error) {
	var record Record
	result := m.db.WithContext(ctx).Where("id = ?", id).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

func (m *RecordModel) GetRecentRecorById(ctx context.Context, id uint64) (*[]Record, error) {
	var record []Record
	result := m.db.WithContext(ctx).Where("user_id = ?", id).Order("updatetime desc").Limit(2).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

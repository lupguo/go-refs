package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	GiveWait      uint32 = 0 // 待赠课
	GiveCancelled uint32 = 1 // 无需赠课
	GiveSucceed   uint32 = 2 // 赠课成功
	GiveFailed    uint32 = 3 // 赠课失败
)

// GiveCourse 灯塔企业学员课程赠送记录表
type GiveCourse struct {
	ID          uint32    `gorm:"column:id"`            // 主键ID
	StuUid      uint64    `gorm:"column:stu_uid"`       // 学员UID
	StuPhone    string    `gorm:"column:stu_phone"`     // 学员手机号
	CourseId    uint32    `gorm:"column:course_id"`     // 课程ID
	CourseName  string    `gorm:"column:course_name"`   // 课程名称
	TermId      uint32    `gorm:"column:term_id"`       // 期数ID
	TermName    string    `gorm:"column:term_name"`     // 期数名称
	ClassId     uint32    `gorm:"column:class_id"`      // 班级ID
	ComId       uint32    `gorm:"column:com_id"`        // 企业ID
	GiveStatus  uint32    `gorm:"column:give_status"`   // 赠课状态（1：待赠课，2：无需赠课，3：赠课成功，4：赠课失败）
	GiveTime    time.Time `gorm:"column:give_time"`     // 赠课时间
	EndClassPct uint32    `gorm:"column:end_class_pct"` // 班级结课程率
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (c *GiveCourse) TableName() string {
	return "t_give_course"
}

// CourseGiveLog 赠课日志表
type CourseGiveLog struct {
	ID         uint32    `gorm:"column:id"`          // 主键ID
	GiveId     uint32    `gorm:"column:give_id"`     // 赠送记录ID，对应 t_course_give 的ID主键
	OpType     uint32    `gorm:"column:op_type"`     // 赠课类型（1：自动赠课 2：手动赠课）
	StuUid     uint64    `gorm:"column:stu_uid"`     // 学员UID
	CourseId   uint32    `gorm:"column:course_id"`   // 课程ID
	TermId     uint32    `gorm:"column:term_id"`     // 期数ID
	GiveStatus uint32    `gorm:"column:give_status"` // 赠课状态（1：待赠课，2：无需赠课，3：赠课成功，4：赠课失败）
	Log        string    `gorm:"column:log"`         // 备注信息
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
}

func (c *CourseGiveLog) TableName() string {
	return "t_give_course_log"
}

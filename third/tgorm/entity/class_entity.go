package entity

import "time"

// Class 灯塔班级
type Class struct {
	ID         uint32    `gorm:"column:id"`          // 班级ID
	Name       string    `gorm:"column:name"`        // 班级名称
	ComId      uint32    `gorm:"column:com_id"`      // 企业ID
	ComName    string    `gorm:"column:com_name"`    // 企业名称
	CourseId   uint32    `gorm:"column:course_id"`   // 课程ID
	CourseName string    `gorm:"column:course_name"` // 课程名称
	TermId     uint32    `gorm:"column:term_id"`     // 期数ID
	TermName   string    `gorm:"column:term_name"`   // 期数名称
	Province   string    `gorm:"column:province"`    // 开班所在省
	City       string    `gorm:"column:city"`        // 开班所在城市
	Area       string    `gorm:"column:area"`        // 开班所在区域
	StartTime  time.Time `gorm:"column:start_time"`  // 开班日期
	EndTime    time.Time `gorm:"column:end_time"`    // 结班日期
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
}

func (c *Class) TableName() string {
	return "t_class"
}

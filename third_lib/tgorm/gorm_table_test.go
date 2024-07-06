package tgorm

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"x-learn/third/tgorm/entity"
)

var (
	db  *gorm.DB
	err error
	ctx context.Context
)

func init() {
	dsn := "root:Secret123@tcp(9.134.233.187:3306)/lighthouse?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
}

func TestGetDB(t *testing.T) {
	// db, err := sql.Open("mysql", "root:Secrect123.@127.0.0.1/lighthouse")
	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	var courses []*entity.GiveCourse
	tx := db.WithContext(ctx).
		Table("`t_class` AS cl").
		Select("co.*").
		Joins("LEFT JOIN `t_give_course` AS co ON cl.id = co.class_id").
		Where("cl.end_time < ?", time.Now()).
		Where("co.end_class_pct = -1").
		Limit(10).
		Find(&courses)
	// t.Logf(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return gtx.First(courses)
	// }))
	if err := tx.Error; err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	updCourse := &entity.GiveCourse{ID: 11, EndClassPct: 1050}
	// tx := db.WithContext(ctx).
	// 	Model(updCourse).
	// 	Select("end_class_pct").
	// 	// Where("id = ?", updCourse.ID).
	// 	Updates(&updCourse)
	str := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.WithContext(ctx).
			Model(&updCourse).
			Select("end_class_pct").
			Updates(&updCourse)
	})
	t.Logf("sql: %s", str)
	if err != nil {
		t.Error(err)
	}
}

// 批量
func TestUpdates(t *testing.T) {
	updCourses := []*entity.GiveCourse{
		{StuUid: 200, CourseId: 3, TermId: 4, EndClassPct: 130},
	}
	for _, course := range updCourses {
		sqlStr := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.WithContext(ctx).
				Model(&entity.GiveCourse{}).
				Select("end_class_pct", course.EndClassPct).
				Where("stu_uid=?", course.StuUid).
				Where("course_id=?", course.CourseId).
				Where("term_id=?", course.TermId).
				Updates(&course)
		})
		t.Logf("sql: %s", sqlStr)
	}
}

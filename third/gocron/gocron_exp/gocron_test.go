package gocron_exp

import (
	"math/rand"
	"testing"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	// 常规任务
	task = func(t *testing.T) {
		t.Logf("[%s], exec task", time.Now())
	}
	// 序号任务
	taskNumb = func(t *testing.T, jobID int) {
		t.Logf("[%s], job#%d exec task", time.Now(), jobID)
	}
	// 随机耗时任务
	taskLong = func(t *testing.T) {
		if rn := rand.Int31n(2); rn%2 == 0 {
			t.Logf("[%s], run long task...", time.Now())
			time.Sleep(1*time.Minute + 30*time.Second)
		} else {
			t.Logf("[%s], run fast task", time.Now())
		}
	}
)

func TestRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		rn := rand.Int31n(2)
		t.Logf("rn=%d", rn)
	}
}

func TestGoCronExp1(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	// job, _ := s.Every(1).Second().Do(task)
	job, _ := s.Every("1s").Do(func() { task(t) })
	job.LimitRunsTo(2)
	// s.StartBlocking()
	// s.Stop()
	s.StartAsync()
	time.Sleep(4 * time.Second)
}

func TestGoCronExp2(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	t.Logf("scheduler location, %s", s.Location())
	_, _ = s.Every("1s").Do(func() { taskNumb(t, 1) })
	_, _ = s.Every("2s").Do(func() { taskNumb(t, 2) })
	_, _ = s.Every("3s").Do(func() { taskNumb(t, 3) })

	s.StartAsync()
	time.Sleep(5 * time.Second)
}

func TestGoCronAtExp3(t *testing.T) {
	// 配置调度器的时间，用于At执行
	loc, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(loc)

	startTime := "14:47"
	s.Every(1).Days().At(startTime).Do(func() { taskNumb(t, 1) })
	s.Every(1).Days().At(startTime).Do(func() { taskNumb(t, 2) })

	s.StartAsync()
	// 在同一时间调度，间隔开来，需要配置在jobs后面
	s.RunAllWithDelay(2 * time.Second)
	time.Sleep(2 * time.Minute)
}

func TestGoCronString(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)

	// 单例模式，该job如果之前有运行且未完成，则调度器不会继续调度新的job任务
	s.SingletonModeAll()
	s.StartAsync()

	// 分(0-59) 时(0-23) 日(1-31) 月(1-12) 天(0-6)
	job, _ := s.Cron("*/1 * * * *").Do(func() { taskLong(t) }) // every minute
	// job.SingletonMode()
	job.Tag("cron_1_min")

	// 更新cron
	go func() {
		configIsUpd := true
		newCron := "*/2 * * * *"
		if configIsUpd {
			// 如果job已启动，则无法停止
			s.Remove(job)
			s.Cron(newCron).Do(func() { t.Log(time.Now()); task(t) })
		}
	}()

	// s.StartBlocking()
	time.Sleep(10 * time.Minute)
}

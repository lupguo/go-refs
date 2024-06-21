package twire

import (
	"context"
	"log"
)

// App 应用
type App struct {
	srv   *Service
	repos IRepos
}

// NewApp 创建应用实例
func NewApp(srv *Service, repos IRepos) (*App, error) {
	return &App{
		srv:   srv,
		repos: repos,
	}, nil
}

func (app *App) DoThings() error {
	return app.srv.repos.ExecSomething(context.Background(), 88)
}

// IRepos 仓储接口
type IRepos interface {
	ExecSomething(ctx context.Context, id uint32) error
}

// BussInfra 实现
type BussInfra struct {
	dsn string
}

func (b *BussInfra) ExecSomething(ctx context.Context, id uint32) error {
	log.Printf("infra dsn [%s], id#%v exec something...", b.dsn, id)
	return nil
}

// NewBussInfra 创建基础设施实例
// 尽量不要在同一个基础设施初始时候，加入多个不同的dsn，可以通过抽出一个Options或者Config，以及对应的方法进一步方便wire做初始化
func NewBussInfra(dsn string) (*BussInfra, error) {
	return &BussInfra{dsn: dsn}, nil
}

// Service 服务
type Service struct {
	repos IRepos
}

// NewService 创建服务
func NewService(repos IRepos) *Service {
	return &Service{repos: repos}
}

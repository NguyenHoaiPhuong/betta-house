package app

import (
	"github.com/NguyenHoaiPhuong/betta-house/api"
	"github.com/NguyenHoaiPhuong/betta-house/config"
	"github.com/NguyenHoaiPhuong/betta-house/repo"
)

// App struct includes router and mongodb session
type App struct {
	cfg *config.Config
	api *api.API
	db  *repo.MongoDB
}

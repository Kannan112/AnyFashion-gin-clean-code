package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admin req.CreateAdmin, createrId int) (res.AdminData, error)
	AdminLogin(admin req.LoginReq) (string, error)
	ListUsers(ctx context.Context) ([]domain.UsersData, error)
	FindUserByEmail(ctx context.Context, name string) (domain.UsersData, error)
	BlockUser(body req.BlockData, adminId int) error
	UnblockUser(id int) error
	GetDashBord(ctx context.Context) (res.AdminDashboard, error)
	ViewSalesReport(ctx context.Context) ([]res.SalesReport, error)
}

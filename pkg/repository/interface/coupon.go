package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	"golang.org/x/net/context"
)

type CouponRepository interface {
	AddCoupon(ctx context.Context, coupon req.Coupon) error
	UpdateCoupon(ctx context.Context, coupon req.Coupon, CouponId int) error
	DeleteCoupon(ctx context.Context, Couponid int) error
	ViewCoupon(ctx context.Context) ([]domain.Coupon, error)
}

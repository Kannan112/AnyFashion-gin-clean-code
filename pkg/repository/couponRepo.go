package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type CouponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponDatabase{DB}
}

func (c *CouponDatabase) FindCouponByName(ctx context.Context, couponCode string) (bool, error) {
	var check bool
	Querycheck := `SELECT EXISTS(SELECT code FROM coupons WHERE code=$1)`
	err := c.DB.Raw(Querycheck, couponCode).Scan(&check).Error
	if err != nil {
		return false, err
	}
	return check, err
}

func (c *CouponDatabase) AddCoupon(ctx context.Context, coupon req.Coupons) error {
	// var check bool
	// tx := c.DB.Begin()
	// Querycheck := `SELECT EXISTS(SELECT code FROM coupons WHERE code=$1)`
	// err := c.DB.Raw(Querycheck, coupon.Code).Scan(&check).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// if check {
	// 	tx.Rollback()
	// 	return fmt.Errorf("coupon code already exists")
	// }
	query := `INSERT INTO coupons (code,discount_percent,discount_maximum_amount,minimum_purchase_amount, expiration_date)VALUES($1,$2,$3,$4,$5)`
	err := c.DB.Exec(query, coupon.Code, coupon.DiscountPercent, coupon.DiscountMaximumAmount, coupon.MinimumPurchaseAmount, coupon.ExpirationDate).Error
	return err
}
func (c *CouponDatabase) UpdateCoupon(ctx context.Context, coupon req.UpdateCoupon, CouponId int) error {
	var check bool
	tx := c.DB.Begin()
	Querycheck := `SELECT EXISTS(SELECT code FROM coupons WHERE id=$1)`
	err := c.DB.Raw(Querycheck, CouponId).Scan(&check).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if !check {
		tx.Rollback()
		return fmt.Errorf("no coupon code found")
	}

	query := `UPDATE coupons 
	SET discount_percent=$1,
		discount_maximum_amount=$2,
		minimum_purchase_amount=$3,
		expiration_date=$4
	WHERE id=$5;
	`
	err = c.DB.Exec(query, coupon.DiscountPercent, coupon.DiscountMaximumAmount, coupon.MinimumPurchaseAmount, coupon.ExpirationDate, CouponId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (c *CouponDatabase) DeleteCoupon(ctx context.Context, couponId int) error {
	fmt.Println("couponId,", couponId)
	var check bool
	Querycheck := `SELECT EXISTS(SELECT 1 FROM coupons WHERE id=$1)`
	err := c.DB.Raw(Querycheck, couponId).Scan(&check).Error
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("coupon with id not exists")
	}
	query := `DELETE FROM coupons WHERE id=$1`
	err = c.DB.Exec(query, couponId).Error
	return err
}
func (c *CouponDatabase) ViewCoupon(ctx context.Context) ([]domain.Coupon, error) {
	var coupon []domain.Coupon
	query := `SELECT * FROM coupons`
	err := c.DB.Raw(query).Scan(&coupon).Error
	return coupon, err
}
func (c *CouponDatabase) ApplyCoupon(ctx context.Context, userId int, couponCode string) (int, error) {
	tx := c.DB.Begin()
	var check bool
	//Check coupon exists
	CouponExists := `SELECT EXISTS(SELECT * FROM coupons WHERE code=$1)`
	err := tx.Raw(CouponExists, couponCode).Scan(&check).Error
	if err != nil {
		return 0, err
	}
	if !check {
		tx.Rollback()
		return 0, fmt.Errorf("coupon not found")
	}
	//Get the coupon details
	var coupon domain.Coupon
	CouponDetails := `SELECT * FROM coupons WHERE code=$1`
	err = tx.Raw(CouponDetails, couponCode).Scan(&coupon).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//Check coupon is expire

	if coupon.ExpirationDate.Before(time.Now()) {
		tx.Rollback()
		return 0, fmt.Errorf("coupon expired")
	}
	// Check the coupon is used

	var couponUsed bool
	QueryCoupon := `SELECT EXISTS(SELECT 1 FROM orders WHERE coupon_code=$1)`
	err = tx.Raw(QueryCoupon, couponCode).Scan(&couponUsed).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if couponUsed {
		tx.Rollback()
		return 0, fmt.Errorf("coupon has expired")
	}
	// check whether the coupen is already added to the cart
	var cartDetails domain.Carts
	getCartDetails := `SELECT * FROM carts WHERE users_id=$1`
	err = tx.Raw(getCartDetails, userId).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	fmt.Printf("couponTest cartDetails.ID= %v couponID=%v", cartDetails.CouponId, coupon.Id)
	if cartDetails.CouponId == coupon.Id {
		tx.Rollback()
		return 0, fmt.Errorf("coupen is applied to cart")
	}
	//check cart is empty
	if cartDetails.Sub_total == 0 {
		tx.Rollback()
		return 0, fmt.Errorf("cart is empty")
	}
	//check the coupon can apply for minimum purchase amount
	if cartDetails.Sub_total <= int(coupon.MinimumPurchaseAmount) {
		tx.Rollback()
		return 0, fmt.Errorf("minimum purchase amount should be %v", coupon.MinimumPurchaseAmount)
	}
	//check the coupon descount amount is less than the maximum
	discountAmount := (cartDetails.Sub_total / 100) * int(coupon.DiscountPercent)
	if discountAmount > int(coupon.DiscountMaximumAmount) {
		discountAmount = int(coupon.DiscountMaximumAmount)
	}
	//update the cart total with cart.sub_total minuse discount amount

	updateCart := `UPDATE carts SET total=$1,coupon_id=$2 WHERE id=$3`
	err = tx.Exec(updateCart, cartDetails.Sub_total-discountAmount, coupon.Id, cartDetails.Id).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return cartDetails.Total - cartDetails.Sub_total, nil
}

func (c *CouponDatabase) RemoveCoupon(ctx context.Context, userId int) error {
	//check if the coupon is used
	tx := c.DB.Begin()
	var cartDetails domain.Carts
	var CouponDetails domain.Coupon
	getCartDetails := `SELECT * FROM carts WHERE users_id=$1`
	err := tx.Raw(getCartDetails, userId).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if cartDetails.CouponId == 0 {
		return fmt.Errorf("coupon is not used")
	}
	//coupon Details
	details := `select * from coupons where id=$1`
	err = tx.Raw(details, cartDetails.CouponId).Scan(&CouponDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//check if the coupon is expired
	if CouponDetails.ExpirationDate.Before(time.Now()) {
		return fmt.Errorf("coupon is expired")
	}
	//update the cart total with cart.sub_total minuse discount amount
	discountAmount := (cartDetails.Sub_total / 100) * int(CouponDetails.DiscountPercent)
	if discountAmount > int(CouponDetails.DiscountMaximumAmount) {
		discountAmount = int(CouponDetails.DiscountMaximumAmount)
	}
	//update the cart
	update := `UPDATE carts SET coupon_id=$1,total=$2 WHERE id=$3`
	err = tx.Exec(update, nil, discountAmount+cartDetails.Total, cartDetails.Id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

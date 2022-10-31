package CouponRepository

import "sls/internal/entity/couponEntity"

type CouponRepo interface {
	CreateCoupon(coupon *couponEntity.Coupon) (*couponEntity.Coupon, error)
	DeleteCoupon(id string) error
	EditCoupon(id string, coupon *couponEntity.Coupon) (*couponEntity.Coupon, error)
	FindByCode(code string) (*couponEntity.Coupon, error)
	FindAllCoupon() ([]*couponEntity.Coupon, error)
	UseCoupon(id, couponId string) error
}

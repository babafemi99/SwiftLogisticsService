package couponEntity

import (
	"github.com/google/uuid"
)

type Coupon struct {
	CouponId          uuid.UUID `json:"coupon_id"`
	CouponName        string    `json:"coupon_name" validate:"required"`
	CouponDescription string    `json:"coupon_description"`
	CouponCode        string    `json:"coupon_code" validate:"required"`
	CouponDiscount    float32   `json:"coupon_discount" validate:"required"`
	CouponExpiryDate  string    `json:"coupon_expiry_date" validate:"required"`
	CouponMaximum     int       `json:"coupon_maximum"`
}

type UsedCoupons struct {
	UsedCouponsId uuid.UUID `json:"used_coupons_id"`
	UserId        uuid.UUID `json:"user_id" validate:"required"`
	CouponId      uuid.UUID `json:"coupon_id" validate:"required"`
	Count         int       `json:"count"`
}

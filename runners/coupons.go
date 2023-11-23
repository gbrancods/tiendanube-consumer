package runners

import (
	"fmt"

	"github.com/gbrancods/tiendanube/coupons"
)

func CouponsRunner() {
	c := CreateCoupon()
	UpdateCoupon(c.ID)
	GetCouponByID(c.ID)
	GetAllCoupons()
	DeleteCoupon(c.ID)
}

func CreateCoupon() (c coupons.Coupon) {
	c, err := coupons.New(
		coupons.SetCode("COUPONCODE"),
		coupons.SetType("percentage"),
		coupons.SetValue("30.00"),
		coupons.SetStartDate("2023-11-25"),
		coupons.SetEndDate("2023-12-25"),
		coupons.SetMinPrice(10),
		coupons.SetMaxUses(100),
	).Create()

	if err != nil {
		panicErr("Error creating coupon", err)
	}

	prettyPrint("Coupon created!", c)
	return
}

func UpdateCoupon(couponID int) {

	c, err := coupons.New(
		coupons.SetMaxUses(345),
		coupons.SetEndDate("2024-01-01"),
	).Update(couponID)

	if err != nil {
		panicErr("Error updating product", err)
	}

	prettyPrint("Coupon updated!", c)
}

func GetCouponByID(couponID int) {
	c, err := coupons.GetByID(couponID)

	if err != nil {
		panicErr("Error getting coupon", err)
	}

	prettyPrint("Coupon by ID", c)
}

func GetAllCoupons() {
	c, err := coupons.GetAll()

	if err != nil {
		panicErr("Error getting all coupons", err)
	}

	prettyPrint("All Coupons:", c)
}

func DeleteCoupon(couponID int) {
	err := coupons.Delete(couponID)

	if err != nil {
		panicErr("Error getting all coupons", err)
	}

	prettyPrint(fmt.Sprintf("Coupon %d deleted!", couponID), nil)
}

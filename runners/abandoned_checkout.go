package runners

import (
	abandonedcheckout "github.com/gbrancods/tiendanube/abandoned_checkout"
)

func AbandonedCheckoutRunner() {

	r := GetAllAbandonedCheckouts()
	GetAbandonedCheckoutByID(r[0].ID)
	//CreateAbandonedCheckoutCoupon()
}

func GetAllAbandonedCheckouts() (r []abandonedcheckout.Checkout) {
	r, err := abandonedcheckout.GetAll()

	if err != nil {
		panicErr("Error getting abandoned checkouts", err)
	}

	prettyPrint("Abandoned checkouts:", r)
	return
}

func GetAbandonedCheckoutByID(checkoutID int) {
	r, err := abandonedcheckout.GetByID(checkoutID)
	if err != nil {
		panicErr("Error abandoned checkout", err)
	}
	prettyPrint("Abandoned checkouts:", r)
}

func CreateAbandonedCheckoutCoupon(checkoutID, couponID int) {

	r, err := abandonedcheckout.Coupon{
		CouponID: couponID,
	}.Create(checkoutID)

	if err != nil {
		panicErr("Error creating abandoned checkout coupon", err)
	}
	prettyPrint("Coupon added to checkout!", r)
}

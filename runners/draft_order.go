package runners

import (
	"fmt"

	draftorder "github.com/gbrancods/tiendanube/draft_order"
)

func DraftOrderRunner() {

	// Create a product to fill the draft order
	p := CreateProduct()

	df := CreateDraftOrder(p.Variants[0].ID)
	GetDraftOrderByID(df.ID)
	GetAllDraftOrders()

	// Converts it to an order
	ConfirmDraftOrder(df.ID)

	DeleteProduct(p.ID)
}

func CreateDraftOrder(variantID int) (r draftorder.DraftOrder) {
	r, err := draftorder.New(
		draftorder.SetContactName("Contact name"),
		draftorder.SetContactLastName("Last Name"),
		draftorder.SetContactEmail("contact@mail.com"),
		draftorder.SetContactPhone("99 9999-9999"),
		draftorder.SetIdentification("333.333.333-33"),
		draftorder.SetPaymentStatus("unpaid"),
		draftorder.SetNote("Draf Order Note"),
		draftorder.SetProducts([]draftorder.Product{
			{
				VariantID: variantID,
				Quantity:  10,
			},
		}),
		draftorder.SetDiscount("0.00"),
		//draftorder.SetShippingAddress("Shipping address"),
	).Create()
	if err != nil {
		panicErr("Error creating draft order", err)
	}
	prettyPrint("Draft order created!", r)
	return
}

func GetAllDraftOrders() {
	r, err := draftorder.GetAll()
	if err != nil {
		panicErr("Error getting all draft orders", err)
	}
	prettyPrint("All Draft Orders:", r)
}

func GetDraftOrderByID(draftOrderID int) {
	r, err := draftorder.GetByID(draftOrderID)
	fmt.Println(r.Products[0].VariantID)
	if err != nil {
		panicErr("Error getting draft order by ID", err)
	}
	prettyPrint("Draft order:", r)
}

func ConfirmDraftOrder(draftOrderID int) {
	r, err := draftorder.ConfirmCreation(draftOrderID)
	if err != nil {
		panicErr("Error confirming draft order", err)
	}
	prettyPrint("Draft order:", r)
}

func DeleteDraftOrder(draftOrderID int) {
	err := draftorder.Delete(draftOrderID)
	if err != nil {
		panicErr("Error deleting draft order", err)
	}
	prettyPrint("Draft Order deleted!", draftOrderID)
}

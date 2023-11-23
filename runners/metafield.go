package runners

import (
	"github.com/gbrancods/tiendanube/metafields"
)

func MetafieldRunner() {
	CreateMetafield()
	//UpdateMetafield(r.ID)
	//GetMetafieldsByOwner("product")
	//GetMetafieldByID(r.ID)
	//DeleteMetafield(r.ID)
}

// TODO

func CreateMetafield() (r metafields.Metafield) {
	r, err := metafields.New(
		metafields.SetKey("key"),
		metafields.SetValue("Metafield Value"),
		metafields.SetNamespace("namespace"),
		metafields.SetDescription("Metafield Description"),
		metafields.SetOwnerID("188658660"),
		metafields.SetOwnerResource("Product"),
	).Create()
	if err != nil {
		panicErr("Error creating metafield", err)
	}
	prettyPrint("Metafield created!", r)
	return
}

func UpdateMetafield(metafieldID int) {
	r, err := metafields.New().Update(metafieldID)

	if err != nil {
		panicErr("Error updating metafield", err)
	}
	prettyPrint("Metafield updated!", r)
}

func GetMetafieldsByOwner(ownerResourceID string) {
	r, err := metafields.GetByOwnerResource(ownerResourceID)
	if err != nil {
		panicErr("Error getting all metafields", err)
	}
	prettyPrint("All metafields:", r)
}

func GetMetafieldByID(metafieldID int) {
	r, err := metafields.GetByID(metafieldID)
	if err != nil {
		panicErr("Error updating metafield", err)
	}
	prettyPrint("Metafield", r)
}

func DeleteMetafield(categoryID int) {
	err := metafields.Delete(categoryID)
	if err != nil {
		panicErr("Error updating metafield", err)
	}
	prettyPrint("Metafield deleted!", categoryID)
}

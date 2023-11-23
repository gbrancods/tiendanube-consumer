package runners

import (
	productimages "github.com/gbrancods/tiendanube/product_images"
)

func ProductImageRunner() {

}

func CreateProductImage(productID int) (r productimages.ProductImage) {
	r, err := productimages.New(
		productimages.SetFilename("product-image"),
		productimages.SetSrc(""),
		productimages.SetPosition(1),
	).Create(productID)
	if err != nil {
		panicErr("Error creating product image", err)
	}

	prettyPrint("Product image created!", r)

	return
}

func UpdateProductImage(productID, ImageID int) {
	r, err := productimages.New().Update(productID, ImageID)

	if err != nil {
		panicErr("Error updating product image", err)
	}
	prettyPrint("Product image updated!", r)
}

func GetProductImagesByProduct(productID, ImageID int) {
	r, err := productimages.GetByProduct(productID)
	if err != nil {
		panicErr("Error getting product images by product", err)
	}
	prettyPrint("Product images:", r)
}

func GetProductImagesByID(productID, ImageID int) {
	r, err := productimages.GetByID(productID, ImageID)
	if err != nil {
		panicErr("Error updating product image", err)
	}
	prettyPrint("Product image:", r)
}

func DeleteProductImage(productID, imageID int) {
	err := productimages.Delete(productID, imageID)
	if err != nil {
		panicErr("Error deleting product image", err)
	}
	prettyPrint("Product image deleted!", imageID)
}

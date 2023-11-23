package runners

import (
	"github.com/gbrancods/tiendanube/i18n"
	"github.com/gbrancods/tiendanube/products"
)

func ProductRunner() {
	p := CreateProduct()
	UpdateProduct(p)
	GetProductByID(p.ID)
	GetAllProducts()
	DeleteProduct(p.ID)
}

func CreateProduct() (p products.Product) {
	p, err := products.New(
		products.SetName(
			i18n.Langs{
				En: "EN Name",
				Es: "ES Nombre",
				Pt: "PT Nome",
			},
		),
		products.SetImages(
			[]products.Image{
				{Src: "https://fingers-site-production.s3.eu-central-1.amazonaws.com/uploads/images/szLui8773HimqPgfZcnOSt1jcqsUYcJlnaHepZ50.jpg"},
			},
		),
		//products.SetVariants(
		//	[]productvariant.Variant{
		//		productvariant.SetPrice(10),
		//	},
		//),
	).Create()
	if err != nil {
		panicErr("Error creating product", err)
	}
	prettyPrint("Product created!", p)
	return
}

func UpdateProduct(p products.Product) {

	p.Name = i18n.Langs{
		En: "Updated name",
		Es: "Nombre actualizado",
		Pt: "Nome atualizado",
	}

	r, err := p.Update()
	if err != nil {
		panicErr("Error updating product", err)
	}
	prettyPrint("Product updated!", r)
}

func GetAllProducts() {
	p, err := products.GetAll()
	if err != nil {
		panicErr("Error getting all products", err)
	}
	prettyPrint("Products:", p)
}

func GetProductByID(productID int) {
	p, err := products.GetByID(productID)
	if err != nil {
		panicErr("Error getting product", err)
	}
	prettyPrint("Product:", p)
}

func DeleteProduct(productID int) {
	err := products.Delete(productID)
	if err != nil {
		panicErr("Error deletting product", err)
	}
	prettyPrint("Product deleted!", productID)
}

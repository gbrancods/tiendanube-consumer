package runners

import (
	"github.com/gbrancods/tiendanube/categories"
	"github.com/gbrancods/tiendanube/i18n"
)

func CategoryRunner() {
	c := CreateCategory()
	UpdateCategory(c)
	GetCategoryByID(c.ID)
	GetAllCategories()
	DeleteCategory(c.ID)
}

func CreateCategory() (r categories.Category) {
	r, err := categories.New(
		categories.SetName(
			i18n.Langs{
				Pt: "Sou uma nova categoria",
			},
		),
		categories.SetGoogleShoppingCategory("Name to google shopping"),
		categories.SetDescription(
			i18n.Langs{
				Pt: "Sou a descrição da categoria",
			},
		),
		categories.SetHandle(
			i18n.Langs{
				Pt: "handle_example",
			},
		),
	).Create()
	if err != nil {
		panicErr("Error creating category", err)
	}

	prettyPrint("Category created!", r)

	return
}

func UpdateCategory(c categories.Category) {
	c.Name = i18n.Langs{Pt: "Name updated!"}
	r, err := c.Update(c.ID)

	if err != nil {
		panicErr("Error updating category", err)
	}
	prettyPrint("Category updated!", r)
}

func GetAllCategories() {
	r, err := categories.GetAll()
	if err != nil {
		panicErr("Error getting all categories", err)
	}
	prettyPrint("All categories", r)
}

func GetCategoryByID(categoryID int) {
	r, err := categories.GetById(categoryID)
	if err != nil {
		panicErr("Error updating category", err)
	}
	prettyPrint("Category", r)
}

func DeleteCategory(categoryID int) {
	err := categories.Delete(categoryID)
	if err != nil {
		panicErr("Error updating category", err)
	}
	prettyPrint("Category deleted!", categoryID)
}

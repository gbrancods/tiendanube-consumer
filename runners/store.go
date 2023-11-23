package runners

import "github.com/gbrancods/tiendanube/store"

func GetStoreInfo() {
	r, err := store.GetInfo()
	if err != nil {
		panicErr("Error getting store info", err)
	}
	prettyPrint("Store Info", r)
}

package runners

import (
	"github.com/gbrancods/tiendanube/webhooks"
)

func WebhookRunner() {
	r := CreateWebhook()
	UpdateWebhook(r.ID)
	GetWebhookByID(r.ID)
	GetAllWebhooks()
	DeleteWebhook(r.ID)
}

func CreateWebhook() (r webhooks.Webhook) {
	r, err := webhooks.New(
		webhooks.SetEvent("product/created"),
		webhooks.SetURL("https://myapp.com/product_created_hook"),
	).Create()
	if err != nil {
		panicErr("Error creating webhook", err)
	}

	prettyPrint("Webhook Created!", r)

	return
}

func UpdateWebhook(webhookID int) {
	r, err := webhooks.New(
		webhooks.SetEvent("category/created"),
		webhooks.SetURL("https://myapp.com/product_created_hook"),
	).Update(webhookID)
	if err != nil {
		panicErr("Error updating webhook", err)
	}
	prettyPrint("Webhook updated!", r)
}

func GetWebhookByID(webhookID int) {
	r, err := webhooks.GetById(webhookID)
	if err != nil {
		panicErr("Error getting webhook", err)
	}
	prettyPrint("Webhook:", r)
}

func GetAllWebhooks() {
	r, err := webhooks.GetAll()
	if err != nil {
		panicErr("Error getting all webhooks", err)
	}
	prettyPrint("Webhooks", r)
}

func DeleteWebhook(webhookID int) {
	err := webhooks.Delete(webhookID)
	if err != nil {
		panicErr("Error deleting webhook", err)
	}
	prettyPrint("Webhook deleted!", webhookID)
}

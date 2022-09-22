package item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	it "github.com/jinfluenza/item-client"
)

func itemOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceItemCreate,
		ReadContext:   resourceItemRead,
		UpdateContext: resourceItemUpdate,
		DeleteContext: resourceItemDelete,
		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Title of the item",
			},
			"body": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Body of the item",
			},
		},
	}
}

func resourceItemCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*it.Client)

	var diags diag.Diagnostics

	title := d.Get("title").(string)
	body := d.Get("body").(string)

	sendItem := it.Item{
		Title: title,
		Body:  body,
	}

	_, err := c.CreateItem(sendItem)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("1")

	if err := d.Set("title", title); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("body", body); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceItemUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*it.Client)

	var diags diag.Diagnostics

	if d.HasChange("body") {
		title := d.Get("title").(string)
		body := d.Get("body").(string)

		item := it.Item{
			Title: title,
			Body:  body,
		}
		ci, err := c.UpdateItem(item)

		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("title", ci.Title); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("body", ci.Body); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceItemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*it.Client)

	var diags diag.Diagnostics

	title := d.Get("title").(string)

	item, err := c.GetItem(title)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("1")

	if err := d.Set("title", item.Title); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("body", item.Body); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceItemDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*it.Client)

	var diags diag.Diagnostics

	item := it.Item{
		Title: d.Get("title").(string),
		Body:  d.Get("body").(string),
	}

	di, err := c.DeleteItem(item)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("title", di.Title); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("body", di.Body); err != nil {
		return diag.FromErr(err)
	}

	return diags

}

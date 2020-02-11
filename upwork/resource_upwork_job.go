package upwork

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {

	_ = resourceJobRead(d, m)
	return nil
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

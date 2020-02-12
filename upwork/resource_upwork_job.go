package upwork

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/hr/jobs"
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
	client := m.(api.ApiClient)
	jbs := jobs.New(client)
	params := make(map[string]string)
	params["buyer_team__reference"] = "~12345abcdf"
	params["title"] = "Test oAuth API create job"
	params["job_type"] = "hourly"
	params["description"] = "A description"
	params["visibility"] = "public"
	params["category2"] = "Web, Mobile & Software Dev"
	params["subcategory2"] = "Web Development"
	params["skills"] = "python;javascript"
	resp, _ := jbs.PostJob(params)
	println(resp.Status)
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

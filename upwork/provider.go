package upwork

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/upwork/golang-upwork/api"
)

func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"consumer_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPWORK_CONSUMER_KEY", ""),
				Description: "",
			},
			"consumer_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPWORK_CONSUMER_SECRET", ""),
				Description: "",
			},
			"access_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPWORK_ACCESS_TOKEN", ""),
				Description: "",
			},
			"access_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPWORK_ACCESS_SECRET", ""),
				Description: "",
			},
			"debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPWORK_DEBUG", ""),
				Description: "",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"upwork": resourceJob(),
		},
		ConfigureFunc: configureFunc,
	}
}

func configureFunc(rd *schema.ResourceData) (interface{}, error) {

	config := &api.Config{
		ConsumerKey:    rd.Get("consumer_key").(string),
		ConsumerSecret: rd.Get("consumer_secret").(string),
		AccessToken:    rd.Get("access_token").(string),
		AccessSecret:   rd.Get("access_secret").(string),
		Debug:          rd.Get("debug").(bool),
	}

	client := api.Setup(config)
	return client, nil
}

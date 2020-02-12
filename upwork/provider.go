package upwork

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/auth"
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
			"upwork_job": resourceJob(),
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
	if !client.HasAccessToken(){
		aurl := client.GetAuthorizationUrl("")

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Visit the authorization url and provide oauth_verifier for further authorization")
		fmt.Println(aurl)
		verifier, _ := reader.ReadString('\n')

		// get access token
		token := client.GetAccessToken(verifier)
		fmt.Println(token)


	}

	// http.Response and []byte will be return, you can use any
	_, jsonDataFromHttp1 := auth.New(client).GetUserInfo()

	// here you can Unmarshal received json string, or do any other action(s)
	fmt.Println(string(jsonDataFromHttp1))

	return client, nil
}

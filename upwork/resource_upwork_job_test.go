package upwork

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	testUpworkProvider = Provider().(*schema.Provider)
	testUpworkProviders = map[string]terraform.ResourceProvider{
		"upwork": testUpworkProvider,
	}
	testUpworkProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			":": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
	testUpworkProviderFunc = func() *schema.Provider { return testUpworkProvider }
}

func TestDibbyPProject_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testUpworkProviders,
		Steps: []resource.TestStep{
			{
				Config: projectResource("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("upwork_job.test", "name"),
				),
			},
		},
	})
}
func TestResourceProjectJob(t *testing.T) {
	_ = resourceJobCreate(nil, nil)
}

func projectResource(name string) string {
	return fmt.Sprintf(basicConfig(), name)
}

func basicConfig() string {
	return `
resource "upwork_job" "test" {
  name = "%s"
}
`
}

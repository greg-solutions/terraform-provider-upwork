package upwork

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testUpworkProvider *schema.Provider
var testUpworkProviders map[string]terraform.ResourceProvider
var testUpworkProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testUpworkProviderFunc func() *schema.Provider

func init() {
	testUpworkProvider = Provider().(*schema.Provider)
	testUpworkProviders = map[string]terraform.ResourceProvider{
		"upwork": testUpworkProvider,
	}
	testUpworkProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"upwork": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
	testUpworkProviderFunc = func() *schema.Provider { return testUpworkProvider }
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

# Terraform Upwork provider

Terraform allows you to keep your infrastructure as a code (HCL language).
This provider supports controlling upwork over its http api.

## Configuration of provider

```hcl
provider "upwork" {
  consumer_key: "YOUR_CONSUMER_KEY",
  consumer_secret: "YOUR_CONSUMER_SECRET",
  access_token: "",
  access_secret: "",
  debug: "off"
}
```
provider "upwork" {
  consumer_key= "YOUR_CONSUMER_KEY"
  consumer_secret= "YOUR_CONSUMER_SECRET"
  access_token= ""
  access_secret= ""
  debug= false
}

resource "upwork_job" "devops_job" {
  name = "DevOps"
}
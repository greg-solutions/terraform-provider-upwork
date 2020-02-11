NAME=terraform-provider-upwork


build:
	go build -o $(NAME) main.go

check:
	make build
	terraform init
	terraform plan
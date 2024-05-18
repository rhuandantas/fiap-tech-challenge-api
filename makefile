install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/swag
	go get github.com/onsi/ginkgo/v2@latest
	go install github.com/onsi/ginkgo/v2/ginkgo
generate-swagger: install
	swag init
generate-di:
	# TODO if the files doesnt exists ignore the exit(1)
	rm wire_gen.go
	go generate

set-vars:
	export DB_PASS=12345678
	export DB_NAME=tech_challenge
	export DB_USER=root
	export DB_HOST=mysql
	export DB_PORT=3306
	export AUTH_SECRET=Testando
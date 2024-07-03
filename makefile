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

mock:
	mockgen -source=internal/adapters/repository/cliente.go -package=mock_repo -destination=test/mock/repository/cliente.go
	mockgen -source=internal/adapters/repository/produto.go -package=mock_repo -destination=test/mock/repository/produto.go

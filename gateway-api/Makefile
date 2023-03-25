server: 
	go run main.go 
gen: 
	rm -rf pb/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*/*.proto
mock:
	rm -rf user/mock/*
	mockgen -package=mock -destination=mock/user.mock.go aqrus/Microservice/gateway-api/user UserController,UserService,UserRepository
xembd: 
	docker exec -it postgres_container bash; psql -U root & \c ecommerce
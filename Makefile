
update_mocks:
	mockgen -source internal/endpoints/user.go -destination internal/mocks/endpoints/user.go
	mockgen -source internal/service/user.go -destination internal/mocks/service/user.go
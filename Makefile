.PHONY: duild zip deploy


build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/createorganizationsettings cmd/awslambda/organizationsettings/create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/deleteorganizationsettings cmd/awslambda/organizationsettings/delete/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/getorganizationsettings cmd/awslambda/organizationsettings/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/updateorganizationsettings cmd/awslambda/organizationsettings/update/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/restoredeletedorganizationsettings cmd/awslambda/organizationsettings/restoredeleted/main.go


zip:
	zip -r bin/organizationsettings/createorganizationsettings.zip bin/organizationsettings/createorganizationsettings
	zip -r bin/organizationsettings/deleteorganizationsettings.zip bin/organizationsettings/deleteorganizationsettings
	zip -r bin/organizationsettings/getorganizationsettings.zip bin/organizationsettings/getorganizationsettings
	zip -r bin/organizationsettings/updateorganizationsettings.zip bin/organizationsettings/updateorganizationsettings
	zip -r bin/organizationsettings/restoredeletedorganizationsettings.zip bin/organizationsettings/restoredeletedorganizationsettings


deploy: build zip

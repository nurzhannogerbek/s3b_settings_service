.PHONY: duild zip deploy


build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/createorganizationsettings cmd/awslambda/organizationsettings/create/main.go


zip:
	zip -r bin/organizationsettings/createorganizationsettings.zip bin/organizationsettings/createorganizationsettings


deploy: build zip

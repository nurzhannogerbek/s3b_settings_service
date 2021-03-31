.PHONY: duild zip deploy


build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/createorganizationsettings cmd/awslambda/organizationsettings/create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/deleteorganizationsettings cmd/awslambda/organizationsettings/delete/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/getorganizationsettings cmd/awslambda/organizationsettings/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/updateorganizationsettings cmd/awslambda/organizationsettings/update/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/restoredeletedorganizationsettings cmd/awslambda/organizationsettings/restoredeleted/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/facebookmessenger/getfacebookpages cmd/awslambda/facebookmessenger/getfacebookpages/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/createchannel cmd/awslambda/channel/createchannel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/getchannel cmd/awslambda/channel/getchannel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/getchannels cmd/awslambda/channel/getchannels/main.go


zip:
	zip -r bin/organizationsettings/createorganizationsettings.zip bin/organizationsettings/createorganizationsettings
	zip -r bin/organizationsettings/deleteorganizationsettings.zip bin/organizationsettings/deleteorganizationsettings
	zip -r bin/organizationsettings/getorganizationsettings.zip bin/organizationsettings/getorganizationsettings
	zip -r bin/organizationsettings/updateorganizationsettings.zip bin/organizationsettings/updateorganizationsettings
	zip -r bin/organizationsettings/restoredeletedorganizationsettings.zip bin/organizationsettings/restoredeletedorganizationsettings

	zip -r bin/facebookmessenger/getfacebookpages.zip bin/facebookmessenger/getfacebookpages

	zip -r bin/channel/createchannel.zip bin/channel/createchannel
	zip -r bin/channel/getchannel.zip bin/channel/getchannel
	zip -r bin/channel/getchannels.zip bin/channel/getchannels


deploy: build zip

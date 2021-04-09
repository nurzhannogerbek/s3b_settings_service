.PHONY: duild zip deploy


build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/createorganizationsettings cmd/awslambda/organizationsettings/createorganizationsettings/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/deleteorganizationsettings cmd/awslambda/organizationsettings/deleteorganizationsettings/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/getorganizationsettingsbyid cmd/awslambda/organizationsettings/getorganizationsettingsbyid/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/updateorganizationsettings cmd/awslambda/organizationsettings/updateorganizationsettings/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organizationsettings/restoredeletedorganizationsettings cmd/awslambda/organizationsettings/restoredeletedorganizationsettings/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/createorganization cmd/awslambda/organization/createorganization/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/createorganizationdepartment cmd/awslambda/organization/createorganizationdepartment/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/deleteorganizations cmd/awslambda/organization/deleteorganizations/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/getallarchivedorganizationdepartments cmd/awslambda/organization/getallarchivedorganizationdepartments/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/getallorganizationdepartments cmd/awslambda/organization/getallorganizationdepartments/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/getarchivedorganizationdepartmentsbyid cmd/awslambda/organization/getarchivedorganizationdepartmentsbyid/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/getorganizationbyid cmd/awslambda/organization/getorganizationbyid/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/getorganizationdepartmentsbyid cmd/awslambda/organization/getorganizationdepartmentsbyid/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/getorganizationsbyids cmd/awslambda/organization/getorganizationsbyids/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/organization/restoredeletedorganizations cmd/awslambda/organization/restoredeletedorganizations/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/facebookmessenger/getfacebookpages cmd/awslambda/facebookmessenger/getfacebookpages/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/createchannel cmd/awslambda/channel/createchannel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/updatechannel cmd/awslambda/channel/updatechannel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/getchannel cmd/awslambda/channel/getchannel/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/channel/getchannels cmd/awslambda/channel/getchannels/main.go


zip:
	zip -r bin/organizationsettings/createorganizationsettings.zip bin/organizationsettings/createorganizationsettings
	zip -r bin/organizationsettings/deleteorganizationsettings.zip bin/organizationsettings/deleteorganizationsettings
	zip -r bin/organizationsettings/getorganizationsettingsbyid.zip bin/organizationsettings/getorganizationsettingsbyid
	zip -r bin/organizationsettings/updateorganizationsettings.zip bin/organizationsettings/updateorganizationsettings
	zip -r bin/organizationsettings/restoredeletedorganizationsettings.zip bin/organizationsettings/restoredeletedorganizationsettings

	zip -r bin/organization/createorganization.zip bin/organization/createorganization
	zip -r bin/organization/createorganizationdepartment.zip bin/organization/createorganizationdepartment
	zip -r bin/organization/deleteorganizations.zip bin/organization/deleteorganizations
	zip -r bin/organization/getallarchivedorganizationdepartments.zip bin/organization/getallarchivedorganizationdepartments
	zip -r bin/organization/getallorganizationdepartments.zip bin/organization/getallorganizationdepartments
	zip -r bin/organization/getarchivedorganizationdepartmentsbyid.zip bin/organization/getarchivedorganizationdepartmentsbyid
	zip -r bin/organization/getorganizationbyid.zip bin/organization/getorganizationbyid
	zip -r bin/organization/getorganizationdepartmentsbyid.zip bin/organization/getorganizationdepartmentsbyid
	zip -r bin/organization/getorganizationsbyids.zip bin/organization/getorganizationsbyids
	zip -r bin/organization/restoredeletedorganizations.zip bin/organization/restoredeletedorganizations

	zip -r bin/facebookmessenger/getfacebookpages.zip bin/facebookmessenger/getfacebookpages

	zip -r bin/channel/createchannel.zip bin/channel/createchannel
	zip -r bin/channel/updatechannel.zip bin/channel/updatechannel
	zip -r bin/channel/getchannel.zip bin/channel/getchannel
	zip -r bin/channel/getchannels.zip bin/channel/getchannels


deploy: build zip

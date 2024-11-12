build:
	@go build 

install:
	@go build 
	@go install 

clean:
	@rm fileshare-client

run:
	@go build 
	@go install 
	@fileshare-client

FROM golang:1.18-alpine

# copy local code to the container
COPY . /go/src/app/
WORKDIR /go/src/app/

# get the required modules and install the application
RUN go get -d -v ./... && \
	go install -v ./...

# run the application
CMD ["atm"]

FROM golang

# set working directory 
WORKDIR /go/src/app

# copy the source code 
COPY . .

# expose the port 
EXPOSE 8000

# build the go app 
RUN go build -o main cmd/main.go

# run the exe
CMD ["./main"]
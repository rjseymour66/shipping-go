# base image of most recent version of Go
FROM golang:1.17 AS deps

# create a working directory to store source code
WORKDIR /hello-api
# Add dependency files and download the contents
ADD *.mod *.sum ./
RUN go mod download 

# Create a new stage of the build to leverage caching
FROM deps as dev 
ADD . .
EXPOSE 8080
# Go flags for container optimization
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -X main.docker=true" -o api cmd/main.go
CMD ["/hello-api/api"]

# Use base scratch image
FROM scratch as prod 

WORKDIR / 
EXPOSE 8080
# Copy binary over from the dev stage
COPY --from=dev /hello-api/api /
CMD ["/api"]
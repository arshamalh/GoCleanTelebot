FROM golang:1.15-alpine AS build

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/signalvip

# Allow pulling private modules
ARG GITLAB_LOGIN
ARG GITLAB_TOKEN
RUN echo "machine git.coryptex.com login ${GITLAB_LOGIN} password ${GITLAB_TOKEN}" > ~/.netrc
RUN go env -w GOPRIVATE=git.coryptex.com
RUN go env -w GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY . .

# Build the Go app
WORKDIR /app/signalvip/src
RUN go build -o ./out/signalvip .

FROM alpine AS final

# Copy compile app to final light weight image
COPY --from=build /app/signalvip/src/out/signalvip /signalvip

# http port
EXPOSE 8081

## Run the binary
ENTRYPOINT ["/signalvip"]
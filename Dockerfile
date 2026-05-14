FROM golang:1.22-alpine AS build
WORKDIR /src
COPY main.go .
RUN go build -ldflags="-s -w" -o /out/hello main.go

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/hello /hello
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/hello"]

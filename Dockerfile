FROM golang:1.13.7 AS build
WORKDIR /go/src/app
COPY . .
RUN go build -o terraform-eval src/main.go

FROM hashicorp/terraform:light
WORKDIR /root/
COPY --from=build /go/src/app/terraform-eval .
ENTRYPOINT ["./terraform-eval"]
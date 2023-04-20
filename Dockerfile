FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN cd cmd && go build -o /app/operator


FROM scratch
WORKDIR /
COPY --from=build /app/operator /
EXPOSE 8080
ENTRYPOINT ["/operator"]
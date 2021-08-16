FROM alpine:3.14.1

COPY --from=golang:1.16-alpine /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /go/src/github.com/lindsaygelle/nook

COPY . .

RUN go mod download all

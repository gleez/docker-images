FROM golang:1.9.2 as builder

RUN mkdir -p /go/src/github.com/gleez

# Copy source code
COPY studio /go/src/github.com/gleez/studio

RUN rm -rf glide.lock studio

WORKDIR /go/src/github.com/gleez/studio

RUN curl https://glide.sh/get | sh
RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o studio .


FROM scratch
MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>

# Add the binary
COPY --from=builder /go/src/github.com/gleez/studio/studio /bin/

# Add the root CAs
# ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/

#RUN chmod a+x /bin/studio

# No arguments are needed for running the app
ENTRYPOINT ["/bin/studio"]

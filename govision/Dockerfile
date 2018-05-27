FROM openalpr/openalpr

# Install golang

RUN apt-get update && apt-get install -y\
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config

ENV GOLANG_VERSION 1.10.2

RUN curl -O https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz; \
        tar xvf go1.10.2.linux-amd64.tar.gz; \
        rm go1.10.2.linux-amd64.tar.gz;\
        chown -R root:root ./go; \
        mv go /usr/local; \
        export PATH="/usr/local/go/bin:$PATH";

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# WORKDIR /srv/openalpr/src/build/bindings/go
# RUN make install

WORKDIR $GOPATH
RUN go get github.com/openalpr/openalpr/src/bindings/go/openalpr
RUN go get github.com/patrickmn/go-cache

WORKDIR $GOPATH/vision
COPY . .

EXPOSE 8123

ENTRYPOINT ["/bin/sh","-c"]
CMD [ "go run main.go"]
FROM    golang:1.7.3
COPY    . /go/src/github.com/moul/translator
WORKDIR /go/src/github.com/moul/translator
CMD     ["translator"]
EXPOSE  8000 9000
RUN     make install

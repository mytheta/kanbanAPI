FROM golang:1.9

LABEL maintainer = "Yutsuki Miyashita <j148015n@st.u-gakugei.ac.jp>"
LABEL description = "kanbanAPI"

WORKDIR /
ENV GOPATH /go
ENV APIDIR ${GOPATH}/src/github.com/mytheta/kanbanAPI

COPY . ${APIDIR}
RUN go get github.com/alecthomas/template
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/swag
RUN go get -u github.com/swaggo/gin-swagger/swaggerFiles
RUN cd ${APIDIR} && rm -rf ./docs
RUN cd ${APIDIR} && go build -o bin/kanbanAPI ./main.go
RUN cd ${APIDIR} && cp bin/kanbanAPI /usr/local/bin/

CMD ["kanbanAPI"]

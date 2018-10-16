FROM golang:1.9

LABEL maintainer = "Yutsuki Miyashita <j148015n@st.u-gakugei.ac.jp>"
LABEL description = "kanbanAPI"

WORKDIR /
ENV GOPATH /go
ENV APIDIR ${GOPATH}/src/github.com/mytheta/kanbanAPI

COPY . ${APIDIR}
RUN cd ${APIDIR} && go build -o bin/kanbanAPI ./main.go
RUN cd ${APIDIR} && cp bin/kanbanAPI /usr/local/bin/

CMD ["kanbanAPI"]

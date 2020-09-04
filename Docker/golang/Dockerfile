# ベースとなるDockerイメージ指定
FROM golang:1.14

RUN go get -u github.com/gin-gonic/gin \
    && go get github.com/oxequa/realize

EXPOSE 8080

CMD [ "realize", "start", "--run" ]

FROM node:10-alpine as uibuilder
COPY survey-ui /survey-ui
RUN cd /survey-ui && npm install
RUN cd /survey-ui && node_modules/@angular/cli/bin/ng build --prod

FROM golang:1.14 as serverbuilder
WORKDIR /survey
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=uibuilder /survey-ui/dist  /server/survey-ui/dist
COPY surveys1.json  /server
COPY --from=serverbuilder /survey/main  /server

WORKDIR  /server

CMD ["./main"]

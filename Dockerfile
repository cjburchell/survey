FROM scratch

COPY survey-ui/dist/.  /server/survey-ui/dist
COPY main  /server

WORKDIR  /server

CMD ["/server/main"]

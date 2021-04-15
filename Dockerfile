FROM scratch
COPY ./bin/feuer /
ENTRYPOINT [ "./feuer" ]
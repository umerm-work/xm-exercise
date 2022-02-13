############################
# STEP 1 build executable binary
############################
FROM golang:1.16-alpine
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -o ./bin/application .

############################
# STEP 2 build a small image
############################
#FROM alpine:3.9
##FROM scratch
#COPY --from=builder /app/bin/application /application
#
#RUN chmod +x application

EXPOSE 8080

CMD ["./application"]

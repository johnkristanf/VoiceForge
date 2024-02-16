FROM golang:1.21.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ENV DATABASE_URL="postgres://voiceforgedb:plzIIvHHRaBTTMjcEADwBFc8MVuxbZYO@dpg-cn07tfed3nmc7389p0l0-a.oregon-postgres.render.com/voiceforge_emyw"
ENV REDIS_URL="rediss://red-cn07ve6d3nmc7389p7c0:Ush4HtpRUuTdBx59qB5eo2Tdb1rri1e1@oregon-redis.render.com:6379"

ENV SMTP_HOST="smtp.gmail.com"
ENV SMTP_PORT=465
ENV SMTP_PASSWORD="bkzd cgbb wywn qmvi"
ENV SMTP_FROM="johnkristan01@gmail.com"

ENV AUTHORIZATION_API_KEY="5e8988bd9f954d13824c6a51ffdbd634"
ENV USER_API_KEY="cqyGECJHZjQEkr5aLyliSOPTGmN2"

ENV SERVER_PORT=800

ENV ACCESS_TOKEN_JWTSECRET="_Oyxb3GaAeizA8x5CWVr7olRTB4-K2IgSjhBat8Sr1E"
ENV REFRESH_TOKEN_JWTSECRET="MD2D-i7oSSjxjFDQmgVky-AaAcSlxhRT_BaFyzD6kXE"
ENV VERIFICATION_TOKEN_JWTSECRET="b5d1_tfuiV-vZ8-4MwPbE8iGapTBOnm-v9V6FFlZl88"

COPY . .

RUN go build -o server ./cmd/api/main.go

EXPOSE 800

CMD ["./server"]

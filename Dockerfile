FROM golang:1.21 as builder

ENV NODE_VERSION=20.10.0
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN node --version
RUN npm --version

WORKDIR /
COPY frontend /frontend
WORKDIR /frontend
RUN npm install
RUN npm run build
RUN ls /frontend/dist

WORKDIR /app
COPY go.mod go.sum ./
COPY .env /.env
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /server

FROM gcr.io/distroless/base-debian11 as final

COPY --from=builder /server /server
COPY --from=builder /frontend /frontend
COPY --from=builder /.env /.env

ENV PORT 3000
EXPOSE $PORT

ENTRYPOINT ["/server"]

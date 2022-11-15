ARG GO_VERSION

FROM golang:${GO_VERSION}

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app

ENV WAIT_VERSION 2.9.0
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait

RUN chmod +x /wait

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air
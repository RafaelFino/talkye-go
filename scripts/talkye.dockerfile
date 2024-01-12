FROM alpine:latest

RUN apk update && apk add --no-cache \
    zsh \
    curl \
    jq \
    openssh-client \
    && rm -rf /var/cache/apk/*

RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

RUN sed -i 's/robbyrussell/agnoster/g' ~/.zshrc

WORKDIR /srv/talkye

ADD ../build/linux/talkye /srv/talkye/talkye
ADD ../etc/config.json /etc/talkye/config.json

ENTRYPOINT ["/srv/talkye/talkye", "/etc/talkye/config.json"]


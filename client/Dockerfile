FROM node:14.13

ARG HTTP_PROXY
ARG HTTPS_PROXY

WORKDIR /available-velibs/client

ADD package.json .

RUN yarn install

ADD . .

CMD ["yarn", "start"]
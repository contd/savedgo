FROM node:11.6.0

RUN npm i -g npm forever && \
  mkdir -p /app;

WORKDIR /app

COPY ./. .

RUN rm node_modules -rf && npm install

CMD forever index.js


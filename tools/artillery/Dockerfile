FROM node:lts-alpine

WORKDIR /usr/src/app

RUN npm i -g artillery

COPY . .

ENV GHOST_ROOT_KEY empty

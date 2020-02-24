FROM node:13.8.0-alpine

RUN npm install elm@latest-0.19.1

ENV PATH /node_modules/elm/bin:$PATH

WORKDIR /src/pants/view/public
WORKDIR /src/pants/view

CMD ["elm", "make", "src/Main.elm", "--output=public/main.js"]

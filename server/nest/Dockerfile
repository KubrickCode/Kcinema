FROM node:18.18-alpine

WORKDIR /app

COPY package.json ./
COPY yarn.lock ./

RUN yarn install --immutable --immutable-cache --check-cache

COPY . .

EXPOSE 3000

CMD ["yarn", "dev"]

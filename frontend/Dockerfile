# FROM node:14-alpine
FROM node:22-alpine

WORKDIR /app

COPY yarn.lock package.json ./
RUN yarn install --frozen-lockfile

COPY index.html ./

CMD ["yarn", "start"]

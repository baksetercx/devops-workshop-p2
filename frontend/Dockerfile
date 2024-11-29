FROM node:23-alpine

WORKDIR /app

COPY yarn.lock package.json ./
RUN yarn install --frozen-lockfile

COPY index.html ./

CMD ["yarn", "start"]

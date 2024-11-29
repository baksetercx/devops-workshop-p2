FROM node:18-alpine3.19

WORKDIR /app

COPY yarn.lock package.json ./
RUN yarn install --frozen-lockfile

COPY index.html ./

CMD ["yarn", "start"]

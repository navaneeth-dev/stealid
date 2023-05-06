FROM docker.io/node as build

WORKDIR /app

RUN npm i -g pnpm

COPY package.json ./
COPY pnpm-lock.yaml ./
RUN pnpm i

COPY . ./

ENV NODE_ENV=production
RUN pnpm build

FROM docker.io/node

WORKDIR /app
COPY --from=build /app .

ENV HOST=0.0.0.0
EXPOSE 4173
CMD ["npm","run", "preview","--", "--host", "0.0.0.0"]

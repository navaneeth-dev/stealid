FROM docker.io/node as build

ARG POCKET_BASE_URL

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

# run prod from build folder
ENV HOST=0.0.0.0
ENV PORT=4173
EXPOSE 4173
CMD ["node","build"]

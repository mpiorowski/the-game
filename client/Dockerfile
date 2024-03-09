# Development
FROM node:21-alpine as dev

WORKDIR /app
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json
RUN npm ci

COPY . .

CMD ["npm", "run", "dev"]

# Production
FROM node:21-alpine as prod
WORKDIR /app
COPY --from=dev /app /app
ENV NODE_ENV=production
ENV VITE_API_URL="https://the-game-server.fly.dev"
ENV VITE_WS_URL="wss://the-game-server.fly.dev"
RUN npm run build

CMD ["node", "build"]

FROM docker.io/library/node:22-slim AS base
WORKDIR /app

FROM base AS builder
COPY . .
RUN npm ci; \
    npm run build; \
    npm prune --omit dev;

FROM base AS runner
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/files ./files
# Check /docker-entrypoint.sh for more info
CMD ["dist/main.js"]

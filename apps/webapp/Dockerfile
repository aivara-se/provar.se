FROM node:20-alpine as builder

WORKDIR /repo

COPY . .
RUN yarn install --silent --immutable
RUN cd apps/webapp && yarn build

# ---

FROM node:20-alpine as service

WORKDIR /app
COPY --from=builder /repo/apps/webapp/build ./
COPY --from=builder /repo/apps/webapp/node_modules ./node_modules

ENV AUTH_EMAIL_FROM=""
ENV AUTH_EMAIL_SERVER=""
ENV AUTH_GITHUB_ID=""
ENV AUTH_GITHUB_SECRET=""
ENV AUTH_GOOGLE_ID=""
ENV AUTH_GOOGLE_SECRET=""
ENV AUTH_SECRET=""
ENV MONGODB_URL=""
ENV NODE_ENV=production
ENV PORT=3000

EXPOSE 3000

ENTRYPOINT ["node","./index.js"]

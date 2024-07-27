# Build
FROM node:18.17.1 AS build
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn vite:build

# Final
FROM golang:1.22
WORKDIR /app
COPY main.go go.mod go.sum ./
COPY VERSION ./
# COPY server ./server
COPY --from=build /app/dist ./dist
RUN go build -o daylang-server .
EXPOSE 5173
CMD ["./daylang-server"]

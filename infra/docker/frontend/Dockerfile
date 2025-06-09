FROM oven/bun:latest
WORKDIR /app
COPY . .
RUN bun install
EXPOSE 4200
CMD ["bun", "run", "dev", "--host", "0.0.0.0"]

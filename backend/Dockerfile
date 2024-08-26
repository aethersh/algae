ARG RUST_VERSION=1.79.0
ARG ALPINE_VERSION=3.20

FROM rust:$RUST_VERSION-slim-bookworm as builder
WORKDIR /server
COPY . .

RUN  --mount=type=cache,target=/usr/local/cargo/registry/ --mount=type=cache,target=/app/target cargo build --release

RUN cp ./target/release/algae /


FROM debian:bookworm-slim as runner
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    appuser
COPY --from=builder /algae /usr/local/bin
RUN chown appuser /usr/local/bin/algae
# COPY --from=builder /app/config /opt/algae/config
# RUN chown -R appuser /opt/algae
USER appuser
ENV RUST_LOG="algae=debug,info"
# WORKDIR /opt/algae
ENTRYPOINT ["algae"]
EXPOSE 2152/tcp

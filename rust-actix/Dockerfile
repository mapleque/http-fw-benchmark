FROM rust:1.64-slim
USER root
WORKDIR /opt
RUN mkdir -p /opt/rust-actix
COPY Cargo.toml /opt/rust-actix/Cargo.toml
COPY src /opt/rust-actix/src
RUN cd /opt/rust-actix && cargo build -r
ENTRYPOINT /opt/rust-actix/target/release/rust-actix

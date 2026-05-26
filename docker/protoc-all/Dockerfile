FROM rust:1.88.0-bookworm AS rust

FROM namely/protoc-all:1.51_2

COPY --from=rust /usr/local/cargo /usr/local/cargo
COPY --from=rust /usr/local/rustup /usr/local/rustup

ENV CARGO_HOME=/usr/local/cargo
ENV RUSTUP_HOME=/usr/local/rustup
ENV PATH=/usr/local/cargo/bin:${PATH}

RUN rustc --version && cargo --version && protoc --version

FROM rust:1.88.0-bookworm AS rust

FROM namely/protoc-all:1.51_2

COPY --from=rust /usr/local/cargo /usr/local/cargo
COPY --from=rust /usr/local/rustup /usr/local/rustup

ENV CARGO_HOME=/usr/local/cargo
ENV RUSTUP_HOME=/usr/local/rustup
ENV PATH=/usr/local/cargo/bin:${PATH}

RUN rm -f /etc/apt/sources.list.d/nodesource.list \
    && printf 'deb http://archive.debian.org/debian bullseye main\ndeb http://archive.debian.org/debian bullseye-updates main\n' > /etc/apt/sources.list \
    && apt-get update \
    && apt-get install -y --no-install-recommends build-essential \
    && rm -rf /var/lib/apt/lists/* \
    && rustc --version \
    && cargo --version \
    && protoc --version \
    && chmod -R a+rwX /usr/local/cargo /usr/local/rustup

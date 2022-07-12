FROM ghcr.io/ryboe/alpinecodespace:latest

ENV PATH="$HOME/.cargo/bin:$PATH"

RUN set -o pipefail \
    && curl --retry 3 --retry-delay 2 --max-time 30 -sSfL https://sh.rustup.rs | sh -s -- -y

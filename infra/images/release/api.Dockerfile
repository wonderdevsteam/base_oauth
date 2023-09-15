FROM python:3.11 AS builder

COPY setup.py .

RUN mkdir /deps && pip install --upgrade pip && pip install  --prefix=/deps  . \
    # temporary complementary install step
    && pip install --prefix=/deps git+https://github.com/Bogdanp/dramatiq

FROM python:3.11

ARG version=2.0.0
ENV VERSION ${version}

WORKDIR /boy

COPY --from=builder /deps /usr/local
COPY . .
COPY ./secrets/cockroach_certificate /home/.postgresql/root.crt

ENV LOG_LEVEL=INFO
ENV LOG_HANDLER=gcp
ENV SENTRY_ENABLED=True
ENV APP_SANDBOX=True

RUN rm ./secrets/cockroach_certificate

ENTRYPOINT ["uvicorn", "boy:serve"]
CMD ["--host", "0.0.0.0", "--port", "8080", "--reload-exclude", "infra/*"]


HEALTHCHECK --interval=60s --timeout=10s --retries=20 CMD curl --fail http://0.0.0.0:8080/ping || exit 1
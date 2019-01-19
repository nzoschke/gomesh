FROM python:slim

RUN pip install yamllint
RUN echo "{extends: default, rules: {key-ordering: {}, line-length: {max: 140}}}" > /etc/yamllint.yaml

WORKDIR /github/workspace
ENTRYPOINT ["yamllint", "-c", "/etc/yamllint.yaml"]
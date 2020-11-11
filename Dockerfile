# 打包成docker镜像
FROM ubuntu:18.04

WORKDIR /workspace

ARG BINARY_FILE
ENV EXECUTE_CMD /workspace/$BINARY_FILE

COPY ./bin/$BINARY_FILE /workspace/
COPY ./conf/ /workspace/conf/

RUN chmod +x /workspace/$BINARY_FILE

CMD ["sh", "-c", "${EXECUTE_CMD}"]
FROM ubuntu:22.04 as builder 
RUN apt update && apt install -y iputils-ping curl wget gcc make apt-transport-https gnupg debian-keyring debian-archive-keyring
RUN curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg && \
    curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | tee /etc/apt/sources.list.d/go-swagger-go-swagger.list
RUN apt update && apt install swagger -y 
# INSTALLAZIONE GO
RUN curl -o /opt/go1.22.tar.gz -L https://golang.org/dl/go1.22.0.linux-amd64.tar.gz 
RUN tar -C /opt/ -xzvf /opt/go1.22.tar.gz && \
    rm -f /opt/go1.22.tar.gz 
ENV PATH=$PATH:/opt/go/bin/

# INSTALLAZIONE SWAGGER UI
# RUN curl -o /usr/local/bin/swagger -L https://github.com/go-swagger/go-swagger/releases/download/v2.8.1/swagger_linux_amd64 && \
    # chmod +x /usr/local/bin/swagger

RUN curl -L https://github.com/swagger-api/swagger-ui/archive/v4.10.3.tar.gz -o swagger-ui.tar.gz && \
    tar -xzf swagger-ui.tar.gz -C /tmp && \
    mkdir -p /myswagger && \
    mv /tmp/swagger-ui-4.10.3/dist/* /myswagger/ && \
    sed -i 's#url: .*#url: "./swagger.yaml",#' /myswagger/swagger-initializer.js
WORKDIR /app
COPY . .
# BUILD DEL CODICE E GENERAZIONE DEL FILE SWAGGER
RUN make build 
RUN swagger generate spec -o /myswagger/swagger.yaml --scan-models



# ----SECOND STAGE---- 
FROM debian:bookworm-slim
WORKDIR /app
# RUN adduser --system --uid 1000 apprunner
# RUN chown -R apprunner:root .
COPY --from=builder  /app/bin/apistore /bin/apistore
COPY --from=builder /myswagger /myswagger
ENV DB_USER=""
ENV DB_NAME=""
ENV DB_HOST=""
ENV DB_PASSWORD=""
ENV JWT_SECRET=""
# USER apprunner
EXPOSE 8080
ENTRYPOINT ["/bin/apistore"]
# CMD [ "bash" , "-c" , "sleep infinity" ]
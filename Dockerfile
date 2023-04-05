FROM golang:latest as build
COPY src/ /build/
WORKDIR /build
ENV TZ Asia/Jakarta
RUN go mod tidy && go build -o /build/app .

FROM ubuntu:latest
COPY --from=build /build/app /app/app

ADD https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-linux.tar.gz .
RUN tar -xzf openshift-client-linux.tar.gz
WORKDIR /usr/bin
RUN cp /oc /bin/oc && cp /kubectl /bin/kubectl
RUN chmod +x /bin/oc && chmod +x /usr/bin/oc
RUN chmod +x /bin/kubectl && chmod +x /usr/bin/kubectl
RUN rm /oc && rm /kubectl && rm /openshift-client-linux.tar.gz

COPY exec.sh /app/exec.sh
RUN chmod +x /app/exec.sh
WORKDIR /app
#ENV GIN_MODE release
ENV TZ Asia/Jakarta
CMD ["/app/exec.sh"]

FROM envoyproxy/envoy:latest
RUN mv /etc/apt/sources.list /etc/apt/sources.list.bak && \
echo 'deb http://mirrors.163.com/debian/ jessie main non-free contrib' > /etc/apt/sources.list && \
echo 'deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib' >> /etc/apt/sources.list && \
echo 'deb http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib' >> /etc/apt/sources.list
RUN apt-get update
COPY envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml

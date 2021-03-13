FROM ubuntu:18.04

RUN apt update && apt install -y nginx wget gzip

RUN wget https://releases.mattermost.com/focalboard/0.5.0/focalboard-server-linux-amd64.tar.gz \
  && gzip -dc focalboard-server-linux-amd64.tar.gz | tar -xvzf - \
  && mv focalboard /opt \
  && rm focalboard-server-linux-amd64.tar.gz

ADD focalboard /etc/nginx/sites-available/

RUN ln -s /etc/nginx/sites-available/focalboard /etc/nginx/sites-enabled/focalboard \
  && rm /etc/nginx/sites-enabled/default \
  && nginx -t \
  && /etc/init.d/nginx reload \
  && /etc/init.d/nginx restart

WORKDIR /opt/focalboard
ENTRYPOINT ["bin/focalboard-server"]

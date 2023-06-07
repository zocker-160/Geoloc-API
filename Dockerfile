FROM scratch

MAINTAINER zocker_160

ADD Geoloc-API /
ADD https://github.com/zocker-160/Geoloc-API/releases/download/0.1/ip-locations.sqlite /

EXPOSE 9001

CMD ["/Geoloc-API", "ip-locations.sqlite"]

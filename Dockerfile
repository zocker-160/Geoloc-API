FROM scratch

MAINTAINER zocker_160

ADD https://github.com/zocker-160/Geoloc-API/releases/download/0.2/Geoloc-API /
ADD https://github.com/zocker-160/Geoloc-API/releases/download/0.1/ip-locations.txt /

EXPOSE 9001

CMD ["/Geoloc-API", "ip-locations.txt"]

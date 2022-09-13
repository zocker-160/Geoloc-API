FROM openjdk:17

MAINTAINER zocker_160

ADD https://github.com/zocker-160/Geoloc-API/releases/download/0.1/Geoloc-API.jar .
ADD https://github.com/zocker-160/Geoloc-API/releases/download/0.1/ip-locations.txt .

ENTRYPOINT ["java", "-jar", "Geoloc-API.jar", "ip-locations.txt"]

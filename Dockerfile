FROM golang:1.21-alpine AS build

RUN apk add --no-cache make

COPY . /build
WORKDIR /build

RUN make buildStaticFinal
RUN chmod +x Geoloc-API

FROM scratch
MAINTAINER zocker_160

COPY --from=build /build/Geoloc-API .
ADD https://github.com/zocker-160/Geoloc-API/releases/download/0.1/ip-locations.txt .

ENV GEOAPI_RAM_OPT "0"
ENV GEOAPI_PORT 9001

EXPOSE $GEOAPI_PORT

CMD ["/Geoloc-API", "ip-locations.txt"] 

FROM golang:alpine
LABEL maintainer = "Hector Lovo <lovohh@gmail.com>"

# Needed to download Go-packages
RUN apk add --no-cache --virtual git

# Directory is based upon current Go-file's references
ENV app_dir /go/src/github.com/lovohh/FoodAcidityAPI/
RUN mkdir -p ${app_dir}
WORKDIR ${app_dir}

# Download live-reloading
RUN go get github.com/pilu/fresh

# Adding Go files
COPY . ${app_dir}

# Download & install Go-dependencies
RUN go get ./...

EXPOSE 3000

CMD ["fresh"]
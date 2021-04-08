# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.15 as build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
# Fuck GFW
ENV GOPROXY=https://goproxy.cn

WORKDIR /go/src/app
COPY . .

RUN  go build  -o k8s-practice .


FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.13.2 as ship
# Fuck GFW
RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/http:\/\/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN sed -i 's/https:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add \
    ca-certificates

RUN addgroup -S app \
    && adduser -S -g app app

WORKDIR /home/app


COPY --from=build /go/src/app/k8s-practice    .
RUN chown -R app:app ./

USER app

CMD ["./k8s-practice"]

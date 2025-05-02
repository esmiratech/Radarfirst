# step1
# ----------------------------------------------

FROM node:18 AS build

WORKDIR /usr/local/app

COPY ./init/package*.json .

RUN --mount=type=cache,target=/root/.yarn YARN_CACHE_FOLDER=/root/.yarn yarn install
RUN yarn global add @angular/cli # add angular

COPY ./init .

RUN yarn build # build the angular project, the output will be in $WORKDIR/dist/ui/browser


# step2
# ----------------------------------------------

FROM nginx:stable-alpine

RUN rm -rf /usr/share/nginx/html/*

COPY --from=build /usr/local/app/dist/ui/browser /usr/share/nginx/html
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
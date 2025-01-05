#! /bin/sh

set -e


cd "$(dirname "$0")/../sumbur"

go build  -o ../public/sumbur  -v

strip -s ../public/sumbur


cd ../public/static

gzip -fkv --  *.css  *.js  || :


cd ../..

rsync \
    --archive \
    --delete-excluded \
    --exclude "sumbur-debug*" \
    --verbose \
    public/ \
    sumbur.su:/srv/web/sumbur.su/sumbur/

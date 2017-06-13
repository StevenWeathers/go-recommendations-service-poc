FROM redis:alpine

COPY load /tmp/load

RUN unix2dos /tmp/load/*.txt

CMD [ "redis-server" ] #, "--appendonly yes"
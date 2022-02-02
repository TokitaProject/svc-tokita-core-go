FROM asia.gcr.io/uii-cloud-project/finance/backend/os/go:1.16-bullseye

ADD main /

EXPOSE 80

CMD ["/main"]
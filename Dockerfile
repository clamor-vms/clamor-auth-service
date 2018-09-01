FROM drone/ca-certs

ADD src/auth /

CMD ["/auth"]

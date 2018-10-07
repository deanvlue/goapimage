FROM scratch
ENV authCode "1234"
ADD app /
EXPOSE 8080
CMD ["/app"]

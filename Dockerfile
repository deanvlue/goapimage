FROM scratch
ENV authCode "1234"
ADD main /
EXPOSE 8080
CMD ["/main"]

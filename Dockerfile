FROM scratch
ENV authCode "1234"
ENV PORT 80
ADD main /
EXPOSE 80
CMD ["/main"]

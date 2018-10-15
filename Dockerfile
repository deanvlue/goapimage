FROM scratch
ENV authCode "1234"
ENV PORT 5005
ADD main /
EXPOSE 5005
CMD ["/main"]

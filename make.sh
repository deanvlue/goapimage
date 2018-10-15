CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t 20181115 jmunozm/dksbxgoldcardapi .
#docker tag latest jmunozm/dksbxgoldcardapi
#docker push deanvlue/dksbxgoldcardapi
docker push jmunozm/dksbxgoldcardapi
#docker build . -t jmunozm:sbxnamedgoldcard
#docker run --publish 6060:8080 --name sbxgoldcard_jcmunoz --rm jmunozm:sbxnamedgoldcard

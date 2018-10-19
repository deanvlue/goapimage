echo "Compiling..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

echo "Generating image..."
docker build . -t jmunozm/dksbxgoldcardapi:latest
echo "Pushing Image..."
docker push jmunozm/dksbxgoldcardapi:latest

echo "Para correr la imagen local: docker run --publish 6060:5005 --name dksbxgoldcardapi --rm jmunozm/dksbxgoldcardapi"
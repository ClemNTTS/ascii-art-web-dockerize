# Utiliser une image de Go
FROM golang:1.22-alpine


LABEL maintainer="cnuttens clement.nutt@gmail.com"
LABEL description="Docker for ascii-art-web"
LABEL version="1.0"
LABEL release-date="2024-06-18"

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum dans le répertoire de travail
COPY go.mod ./

# Télécharger les dépendances. Les dépendances seront mises en cache si les fichiers go.mod et go.sum n'ont pas changé
RUN go mod download

# Copier le reste des fichiers de l'application dans le répertoire de travail
COPY . .

# Construire l'application Go
RUN go build -o main .

EXPOSE 3000

# Définir la commande à exécuter lorsque le conteneur démarre
CMD ["./main"]
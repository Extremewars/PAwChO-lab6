# PAwChO-lab6

polecenia:

DOCKER_BUILDKIT=1 docker build --ssh default -t lab6_test .
docker build --ssh default=/home/pawel/.ssh/id_ed25519 -t lab6_test .
docker run -p 8080:80 lab6_test
export CR_LAB=<klucz>
echo $CR_LAB | docker login ghcr.io -u Extremewars --password-stdin

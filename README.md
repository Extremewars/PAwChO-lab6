# PAwChO-lab6

polecenia:
```bash
DOCKER_BUILDKIT=1 docker build --ssh default -t lab6_test .
```
```bash
docker run -p 8080:80 lab6_test
```
```bash
export CR_LAB=<klucz>
```
```bash
echo $CR_LAB | docker login ghcr.io -u Extremewars --password-stdin
```

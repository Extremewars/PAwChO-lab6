# PAwChO-lab6

Obraz można pobrać poleceniami:
- przez github (ghcr.io, obraz również jest podpięty jako pakiet do tego repozytorium) https://github.com/Extremewars/PAwChO-lab6/pkgs/container/lab6
```bash
docker pull ghcr.io/extremewars/lab6:latest
```
- przez docker hub https://hub.docker.com/r/extremical/lab6
```bash
docker pull extremical/lab6:latest
```
# Tworzenie obrazu

Do utworzenia obrazu użyto polecenia:
```bash
DOCKER_BUILDKIT=1 docker build --ssh default -t lab6 .
```
Uruchomić kontener można poleceniem:
```bash
docker run -p 8080:80 lab6
```

# Publikowanie obrazu

Aby upublicznić obraz przez githuba (ghcr.io) trzeba najpierw się zalogować. Można do tego wykorzystać token utworzony w ustawieniach konta.
```export``` utworzy zmienną w terminalu, która przechowa wartość klucza.
Drugie polecenie pobierze wartość klucza i wprowadzi go do loginu.
```bash
export CR_LAB=<klucz>
echo $CR_LAB | docker login ghcr.io -u <nazwa_użytkownika> --password-stdin
```
Aby upublicznić obraz jako pakiet na github trzeba go nazwać według schematu:
```bash
docker tag <obraz> ghcr.io/<nazwa_użytkownika_github>/<nazwa_pakietu>
docker push ghcr.io/<nazwa_użytkownika_github>/<nazwa_pakietu>
```
Analogicznie jak do docker hub:
```bash
docker tag <obraz> <nazwa_użytkownika_dockerhub>/<nazwa_obrazu>
docker push <nazwa_użytkownika_dockerhub>/<nazwa_obrazu>
```

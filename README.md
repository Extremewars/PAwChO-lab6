# PAwChO-lab6

Do utworzenia obrazu użyto polecenia:
```bash
DOCKER_BUILDKIT=1 docker build --ssh default -t lab6 .
```
Uruchomić kontener można poleceniem:
```bash
docker run -p 8080:80 lab6
```
Aby upublicznić obraz przez githuba (ghcr.io) trzeba najpierw się zalogować. Można do tego wykorzystać token utworzony w ustawieniach konta.
```bash export``` utworzy zmienną w terminalu, która przechowa wartość klucza.
Drugie polecenie pobierze wartość klucza i wprowadzi go do loginu.
```bash
export CR_LAB=<klucz>
echo $CR_LAB | docker login ghcr.io -u <nazwa_użytkownika> --password-stdin
```
Aby upublicznić obraz jako package trzeba go nazwać według schematu:
```bash
docker tag <obraz> ghcr.io/<nazwa_użytkownika>/<nazwa_pakietu>
docker push ghcr.io/<nazwa_użytkownika>/<nazwa_pakietu>
```

1. Собрать образ:
docker build -t admiralu82/moproxy .

2. Создать сеть в docker (поменять адрес сети, шлюз и родительский интерфейс на свои):
docker network create -d macvlan --subnet=10.31.78.0/24 --gateway=10.31.78.1 -o parent=eno1.33 moproxy_vlan33

3. Запустить контейнер с прокси по умолчанию (поменять ip адрес контейнера)
docker run --name moproxy --network moproxy_vlan33 --ip=10.31.78.250 --cap-add=NET_ADMIN admiralu82/moproxy
 или запустить со своим прокси
docker run --name moproxy --network moproxy_vlan33 --ip=10.31.78.250 --cap-add=NET_ADMIN admiralu82/moproxy /run.sh 10.11.12.13:1415

4. Для дигностики запустить консоль в контейнере (откроется консоль - exit выход):
docker exec -it moproxy bash

5. Остановить контейнер:
docker stop moproxy

6. Запустить повторно:
docker start moproxy

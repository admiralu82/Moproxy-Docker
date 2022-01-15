# Настройка Dockera
Образ расчитан на Linux amd64 машины. Контейнер  использует VLAN 33 для свзяи с внешней сетью. 
### 1. Скачать репозитарий:
git clone https://github.com/admiralu82/Moproxy-Docker

cd ./Moproxy-Docker
### 2. Собрать образ:
docker build -t admiralu82/moproxy .
### 3. Создать сеть в docker (поменять адрес сети, шлюз и родительский интерфейс на свои):
docker network create -d macvlan --subnet=**10.31.78.0/24** --gateway=**10.31.78.1** -o parent=**eno1.33** moproxy_vlan33
### 3. Запустить контейнер с прокси по умолчанию (поменять ip адрес контейнера)
docker run --name moproxy --network moproxy_vlan33 --ip=**10.31.78.250** --cap-add=NET_ADMIN admiralu82/moproxy
#### или запустить со своим прокси
docker run --name moproxy --network moproxy_vlan33 --ip=**10.31.78.250** --cap-add=NET_ADMIN admiralu82/moproxy **/run.sh 10.11.12.13:1415**
### 4. Для дигностики запустить консоль в контейнере (откроется консоль - exit выход):
docker exec -it moproxy bash
### 5. Остановить контейнер:
docker stop moproxy
### 6. Запустить повторно:
docker start moproxy

# Настройка маршрутизатора Mikrotik
Код приведён для примера - требуется доработка под своё оборудования.
### Добавляем мост для сети выхода в интернет 
/interface bridge add name=bridge_proxy
### Добавляем VLAN 33 на интерфейс где подключен Docker сервер
/interface vlan add interface=bridge name=vlan33 vlan-id=33
### Добавляем VLAN 33 в мост для сети выхода в интернет
/interface bridge port add bridge=bridge_proxy interface=vlan33
### Добавляем шлюз по умолчанию
/ip route add distance=1 gateway=10.31.78.1
### Настраиваем выход в интернет через другой шлюз
/ip firewall mangle add action=route chain=prerouting dst-address-list=!LanNet dst-port=80,443 passthrough=yes protocol=tcp route-dst=10.31.78.250 src-address-list=LanNet



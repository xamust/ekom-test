# Cервис, который будет считать клики и собирать их в поминутную статистику (timestamp, bannerID, count)

### Планируемая нагрузка:
Максимум 100 баннеров × 1440 минут в сутки = 144,000 документов в день</br>

### Локальный запуск (docker-compose):
cd ./deploy/local </br>
docker-compose build --no-cache </br>
docker-compose up </br>
</br>

### Swagger endpoint:
http://localhost:8080/swagger/index.html

### Vegeta load testing (by default 200 RPS)
make vegeta_load

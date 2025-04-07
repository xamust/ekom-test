### Cервис, который будет считать клики и собирать их в поминутную статистику (timestamp, bannerID, count)
</br>
Максимум 100 баннеров × 1440 минут в сутки = 144,000 документов в день</br>
</br>
cd ./deploy/local </br>
docker-compose build --no-cache </br>
docker-compose up </br>
</br>
http://localhost:8080/swagger/index.html

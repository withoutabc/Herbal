### 不会docker-compose,也懒得每次都重新敲一遍
1. rm /home/xiaote/herbal/project/main
2. go build .\cmd\main.go(本地)
3. scp ./main root@49.7.114.49:/home/xiaote/herbal/project(本地)
4. chmod 777 /home/xiaote/herbal/project/main
5. cd /home/xiaote/herbal/project/
6. docker build -f ./Dockerfile  -t herbal-go-server:5.0 .
7. docker run -it --link mysql-c1:3306 -p 1010:1010 -v ./output:/output herbal-go-server:5.0
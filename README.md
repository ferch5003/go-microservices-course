# go-microservices-course

## Configure Linode
```
1  adduser <username>
2  usermod -aG sudo <username>
3  ufw allow ssh
4  ufw allow http
5  ufw allow https
6  ufw allow 2377/tcp
7  ufw allow 7946/tcp
8  ufw allow 7946/udp
9  ufw allow 4789/udp
10  ufw allow 8025/tcp
11  ufw enable
13  ufw status
```

## Configure Install Docker and Use on remote server
```
1  sudo ls
2  sudo apt-get update
3  sudo apt-get install ca-certificates curl
4  sudo install -m 0755 -d /etc/apt/keyrings
5  sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
6  sudo chmod a+r /etc/apt/keyrings/docker.asc
7  echo   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
$(. /etc/os-release && echo "$VERSION_CODENAME") stable" |   sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
8  sudo apt-get update
9  sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
10  sudo docker run hello-world
11  sudo apt upgrade
12  sudo reboot
14  sudo hostnamectl set-hostname node-1
15  sudo vi /etc/hosts
16  sudo mkdir swarm
17  cd swarm/
18  mkdir caddy_data
19  mkdir caddy_config
20  vi swarm.yml
21  docker stack deploy -c swarm.yml myapp
22  sudo usermod -aG docker fvp
23  docker pull <name>/micro-caddy-production:1.0.1
24  docker service scale myapp_caddy=2
25  docker service update --image <name>/micro-caddy-production:1.0.1 myapp_caddy
26  docker stack deploy -c swarm.yml myapp
```
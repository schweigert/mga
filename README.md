[![Build Status](https://travis-ci.org/schweigert/mga.svg?branch=master)](https://travis-ci.org/schweigert/mga)
[![Coverage Status](https://coveralls.io/repos/github/schweigert/mga/badge.svg?branch=master)](https://coveralls.io/github/schweigert/mga?branch=master)

# mga

Go Massively Game Architecture.

## Dev deps


  - Dev deps: https://github.com/kardianos/govendor
  - Docker:
```
sudo apt remove docker docker-engine docker.io
sudo apt install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu xenial stable"
# sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update
sudo apt install docker-ce
sudo usermod -aG docker exampleuser
sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

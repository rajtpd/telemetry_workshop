#!/bin/sh

echo "Installing zip and unzip packages...."
sudo apt-get update
sudo apt-get install zip unzip -y
echo "Done! Installing zip and unzip packages...."
echo
echo
echo "Installing ansible package...."
sudo apt update
sudo apt install software-properties-common -y
sudo apt-add-repository --yes --update ppa:ansible/ansible
sudo apt install ansible -y
echo "Done! Installing ansible...."
echo
echo
echo "Encrypting vault"
ansible-vault encrypt ~/telemetry_workshop/environment/ansible/secrets.yml
echo "Done encrypting"
echo
echo
echo "Installing Influxdb....."
echo
echo
wget -qO- https://repos.influxdata.com/influxdb.key | sudo apt-key add -
source /etc/lsb-release
echo "deb https://repos.influxdata.com/${DISTRIB_ID,,} ${DISTRIB_CODENAME} stable" | sudo tee /etc/apt/sources.list.d/influxdb.list

sudo apt-get update && sudo apt-get install -y influxdb
sudo systemctl unmask influxdb.service
sudo systemctl start influxdb
echo "Done installing Influxdb"
echo
echo
echo "Installing grafana...."
echo
echo
sudo apt-get install -y software-properties-common
sudo add-apt-repository "deb https://packages.grafana.com/oss/deb stable main"
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install -y grafana
echo "Done installing Grafana"
echo
echo
echo
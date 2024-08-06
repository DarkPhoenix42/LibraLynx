#!/bin/bash

CUR_PATH=$(pwd)

echo "This script will install the required dependencies for the project."

# Read the .env file and export the variables
while IFS== read -r key value; do
  printf -v "$key" %s "$value" && export "$key"
done <.env

echo "Environment variables exported."

echo "Creating mysql database."
mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" < init.sql
echo "Database created."

make migration_up

echo "Downloading apache2 and setting up a virtual host."
sudo apt update && sudo apt-get install apache2 -y
sudo a2enmod proxy proxy_http
cd /etc/apache2/sites-available
sudo mkdir -p /var/log/apache2/
sudo chmod -R 744 /var/log/apache2/

sudo sh -c 'echo "<VirtualHost *:80>
	ServerName libralynx.org
	ServerAdmin praneethsarode@gmail.com
	ProxyPreserveHost On
	ProxyPass / http://127.0.0.1:'$APP_PORT'/
	ProxyPassReverse / http://127.0.0.1:'$APP_PORT'/
	TransferLog /var/log/apache2/libralynx_access.log
	ErrorLog /var/log/apache2/libralynx_error.log
</VirtualHost>" > libralynx.org.conf'

sudo a2ensite libralynx.org.conf
sudo sh -c 'echo "127.0.0.1     libralynx.org" >> /etc/hosts'
sudo a2dissite 000-default.conf

sudo apache2ctl configtest
sudo systemctl restart apache2

echo "Starting the server..."

cd $CUR_PATH
go mod tidy
make
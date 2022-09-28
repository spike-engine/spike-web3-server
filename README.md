## spike-web3-server quick start

### Prequisition
#### SDK
1. Go.  
https://go.dev/
2. Morails.  
https://moralis.io/
3. bscscan.  
https://bscscan.com/
4. Redis.  
https://redis.io/
5. MySQL.  
https://www.mysql.com/
6. ChainStack.  
   https://chainstack.com/

#### Services
1. Signature Service.  
https://github.com/spike-engine/spike-signature-server

### Build and install spike-web3-server

1. Clone the repository
```shell
git clone https://github.com/spike-engine/spike-web3-server.git
cd spike-web3-server/
```
2. Install all the dependencies
```shell
go mod tidy
```
3. Install swagger
```shell
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
sudo mv $(go env GOPATH)/bin/swag /usr/local/bin
swag -v
```
if you want to update swagger doc, please execute :
```shell
swag init
```
4. Make build
```shell
go build -o spike-web3-server ./main.go
```
5. Update Config
```shell
cp config-example.toml config.toml
```
6. Run
```
./spike-web3-server
```

### Register spike as a system service
1. Link server into system binary path
```shell
sudo ln ./spike-web3-server /usr/local/bin
```
2. Copy config file into spike home
```shell
sudo mkdir -p /etc/spike/
sudo cp ./config.toml /etc/spike/config-web3.toml
```
3. Startup script
```shell
sudo vim /etc/systemd/system/spike-web3-server.service
```
Specify the path to the binary
```markdown
[Service] 
ExecStart=/usr/local/bin/spike-web3-server
Environment=SPIKE_WEB3_CONFIG=/etc/spike/config.toml
Restart=always
RestartSec=5 
```
```shell
sudo systemctl daemon-reload
sudo systemctl start spike-web3-server.service
sudo journalctl -u spike-web3-server.service -f
```

### Config
By default, spike-web3-server reads configuration from config.toml under the current folder. 
If it is not created, you may copy from config-example.config.
If you need to specify the config file, you may set the enviornment as follows:
```
export SPIKE_WEB3_CONFIG=~/spike_home/config-web3.toml
```

### Database
```shell
# create database, replace 'spike_web3_server' with DbName config
mysql -u root -p <sql/create-db.sql

# create tables
mysql -u root -p spike_web3_server < sql/create-tables.sql
```

### Swagger
If you run the project successfully, you can visit http://localhost:3000/swagger/index.html.
And you can see some interface information.

# HUM
## build
```
bigdevs:@hum 🌱  $ make
cd src/hum && GO15VENDOREXPERIMENT=1 GOPATH=/Users/bigdevs/Documents/Workspaces/other/hum:/Users/bigdevs/Documents/Workspaces/other/hum:/User
s/bigdevs/go dep ensure
GO15VENDOREXPERIMENT=1 GOPATH=/Users/bigdevs/Documents/Workspaces/other/hum:/Users/bigdevs/Documents/Workspaces/other/hum:/Users/bigdevs/go g
o build -o ./bin/hum ./src/hum
bigdevs:@hum 🌱  $ 
```

## usage
```
bigdevs:@hum 🌱  $ ./bin/hum 
Usage of ./bin/hum:
  -appVersion string
        specific app's version to deploy (default "latest")
  -f string
        specific deployment template
  -registry string
        registry url (default "github.com")
bigdevs:@hum 🌱  $ 
```
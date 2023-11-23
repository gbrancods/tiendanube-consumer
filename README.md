## Test environment for the Tiendanube package

 This package is intended to test the [tiendanube](https://github.com/gbrancods/tiendanube) package

### How to configure with docker:
#### Create your .env file in the root folder and put yours enviroments variables, example:

~~~~
ACCESS_TOKEN=6a2652eaMyAcceessToken9ac9b38
USER_ID=1234567
~~~~

#### To run:

~~~~
docker build -t tiendanube-consumer .
docker-compose up -d
~~~~

#### *Linux users can run `sudo chmod +x up.sh && ./up.sh` , after first use of chmod, just run `./up.sh`
# TCLIENT

### Env variables.
Applicatioin depends on this env variables:

```
TGA_PHONE='+7123456789'

TGA_APP_ID=12345
TGA_API_HASH=''
TGA_PORT=12131 #optional

TGA_MONGO_CONNECTION="mongodb://localhost:27017"

TGA_ENV=development # or "production"

TGA_RETRY_COUNT = 3 #optional, default 3
TGA_RETRY_DELEY_SECONDS=5

```

Also supports .env file on the same lavel as executed binary file. Just copy `.env.example` file

```bash
cp .env.example .env
```

and fill it with your information.

### how to use

- download latest release
- add filled .env file near the binary
- execute it
- do some query to server 
```bash
curl -X GET 'http://localhost:9001/create?sap=123&title=title&admin=someNick' \
-H 'Content-Type: application/json'
```


### For local build:

use this https://github.com/techknowlogick/xgo/

```bash
~/go/bin/xgo --targets=darwin/amd64 --out=tclient --ldflags="-X config.env=production" --dest="$(pwd)/build" ./cmd/main.go
```

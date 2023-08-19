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
```

Also supports .env file on the same lavel as executed binary file. Just copy `.env.example` file

```bash
cp .env.example .env
```

and fill it with your information.


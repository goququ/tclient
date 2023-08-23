FROM techknowlogick/xgo

COPY . /app
WORKDIR /app

CMD xgo \
  --targets="windows/amd64,darwin/amd64" \
  --out=tclient \
  --ldflags="-X config.env=production" \
  --dest="./build" \
  github.com/goququ/tclient/cmd


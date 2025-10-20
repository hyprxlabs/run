# run

run "KEY=VALUE" npm install


run 
  -c "context"
  -D "env-file"
  -p projectnae 
  -a "...:deploy"
  -f "path/to/run/file"


run ...:build

run ./**/test:build

run deploy @traefik 

 #!/usr/bin/env bash

curl -s http://localhost:1317/auth/accounts/$(nscli keys show alice -a)
curl -s http://localhost:1317/auth/accounts/$(nscli keys show jack -a)

#!/usr/bin/zsh
source .env
curl -H "Authorization: $PRE_SHARED_SECRET" `up url -s production`

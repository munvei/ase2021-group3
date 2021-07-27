#!/bin/zsh
token=`cat ./token.txt`
echo token is '"'$token'"'
now=`date "+%H:%M:%S"`
#<<COMMENTOUT 
curl  -X POST https://api.line.me/v2/bot/message/broadcast \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer {'."$token".'}' \
-d '{
    "messages":[
        {
            "type":"text",
            "text":"'$now"にバイクが横転しました"'"
        }
    ]
}'
#COMMENTOUT

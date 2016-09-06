# lab-reserved
representing items in a lab, reserved by an email holder, until some date


## Use

    go get github.com/tompscanlan/labreserved
    cd $GOPATH/src/github.com/tompscanlan/labreserved
    go get ./...
    ./scripts/make-cert.sh &&  ./scripts/build-run.sh &

    curl --request GET   --url http://127.0.0.1:2080/items   --header 'content-type: application/json'
    curl -k --request GET   --url https://127.0.0.1:20443/items   --header 'content-type: application/json'

    curl --request POST --url http://127.0.0.1:2080/item --header 'content-type: application/json' --data '{ "name": "server-9","description": "from rest"}'
    curl --request POST --url http://127.0.0.1:2080/item --header 'content-type: application/json' --data '{ "name": "server-10","description": "from rest"}'
    curl --request POST --url http://127.0.0.1:2080/item --header 'content-type: application/json' --data '{ "name": "server-11","description": "from rest"}'

    curl --request POST  --url http://127.0.0.1:2080/user  --header 'content-type: application/json'  --data '{ "name": "tom","email": "tom@test.com"}'
    curl --request POST  --url http://127.0.0.1:2080/user  --header 'content-type: application/json'  --data '{ "name": "bob","email": "bob@test.com"}'


    curl --request POST --url http://127.0.0.1:2080/item/server1/reservation  --header 'content-type: application/json'  --data '{ "username": "tom","begin": "2016-09-06T00:09:04.032-04:00","hoursheld": 3}'
    curl --request POST --url http://127.0.0.1:2080/item/server-11/reservation  --header 'content-type: application/json'  --data '{ "username": "bob","begin": "2016-09-05T00:09:04.032-04:00","hoursheld": 3}'
    curl --request POST --url http://127.0.0.1:2080/item/server-11/reservation  --header 'content-type: application/json'  --data '{ "username": "bob","begin": "2016-09-06T00:09:04.032-04:00","hoursheld": 24}'

    curl --request GET   --url http://127.0.0.1:2080/items   --header 'content-type: application/json' | jq .

    -----
    {
      "server-10": {
        "description": "from rest",
        "name": "server-10"
      },
      "server-11": {
        "description": "from rest",
        "name": "server-11",
        "reservations": [
          {
            "begin": "2016-09-05T00:09:04.032-04:00",
            "hoursheld": 3,
            "username": "bob"
          },
          {
            "begin": "2016-09-06T00:09:04.032-04:00",
            "hoursheld": 24,
            "username": "bob"
          }
        ]
      },
      "server-9": {
        "description": "from rest",
        "name": "server-9"
      },
      "server1": {
        "description": "testing",
        "name": "server1",
        "reservations": [
          {
            "begin": "2016-09-06T00:09:04.032-04:00",
            "hoursheld": 3,
            "username": "tom"
          },
          {
            "begin": "2016-09-06T00:09:04.032-04:00",
            "hoursheld": 3,
            "username": "tom"
          }
        ]
      },
      "server9": {
        "description": "from rest",
        "name": "server9"
      }
    }

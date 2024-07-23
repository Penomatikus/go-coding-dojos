# Try it!

Disclaimer: .http files migth replace this readme

- You need to have `curl` 
- If you need another port than 8080: `go run . -p 1212` 

### Add a Character

`curl -i -X POST http://localhost:8080/api/v1/fatecore/character/new -H "Content-Type: application/json" -d '{"Name": "Penomatikus", "Description": "goofy gopher"}'`

Would return:   

_HTTP/1.1 200 OK  
Content-Type: application/json  
Date: Tue, 23 Jul 2024 08:58:05 GMT  
Content-Length: 17_
```json
{
    "CharacterID":1
}
```

### Start Session

`curl -i -X POST http://localhost:8080/api/v1/fatecore/session/new -H "Content-Type: application/json" -d '{"Title": "Gophers Place", "OwnerID": 1}'`

Would return:  

_HTTP/1.1 200 OK  
Content-Type: application/json  
Date: Tue, 23 Jul 2024 09:25:25 GMT  
Content-Length: 52_
```json
{
    "SessionID":"d9563ef7-ec2b-4900-8e65-261222360ceb"
}
```

The Character with ID == Owner from the request joined the session automatically.

### Join Session

You need a new character before. Here the character has ID 2. 

For example:  
`curl -i -X POST http://localhost:8080/api/v1/fatecore/character/new -H "Content-Type: application/json" -d '{"Name": "Kalam", "Description": "type scripter"}'`

It will create Kalam with ID 2, who can now join:  
`curl -i -X POST http://localhost:8080/api/v1/fatecore/session/c68d83ff-f409-4313-8aa9-345af4e559f1/join -H "Content-Type: application/json" -d '{"CharacterID": 2 }'`

No return but status code

### Lookup session party as owner

UUID needs to be replaced by your session ID

`curl -i -X GET http://localhost:8080/api/v1/fatecore/session/c68d83ff-f409-4313-8aa9-345af4e559f1/party -H "Content-Type: application/json" -d '{"OwnerID": 1}'`

Would return: 

HTTP/1.1 200 OK  
Content-Type: application/json  
Date: Tue, 23 Jul 2024 10:25:32 GMT  
Content-Length: 139_

```json
{
  "SessionParty": [
    {
      "ID": 1,
      "Name": "Penomatikus",
      "Description": "goofy gopher",
      "SessionID": "c68d83ff-f409-4313-8aa9-345af4e559f1",
      "Points": 0
    },
    {
      "ID": 2,
      "Name": "Kalam",
      "Description": "type scripter",
      "SessionID": "c68d83ff-f409-4313-8aa9-345af4e559f1",
      "Points": 0
    }
  ]
}
```

### Add Points to a Party member as owner

The session ID needs to be yours.

`curl -i -X POST http://localhost:8080/api/v1/fatecore/character/do/points -H "Content-Type: application/json" -d '{
  "ActionName": "adjust points",
  "CharacterID": 2,
  "OwnerAction": true,
  "Costs": 100,
  "SessionID": "c68d83ff-f409-4313-8aa9-345af4e559f1"
}'`

Would return: 

_HTTP/1.1 200 OK  
Content-Type: application/json  
Date: Tue, 23 Jul 2024 11:18:01 GMT  
Content-Length: 36_

```json
{
    "dice_roll":0,
    "current_points":100
}
```

### Do an action as a member

`curl -i -X POST http://localhost:8080/api/v1/fatecore/character/do/action -H "Content-Type: application/json" -d '{
  "ActionName": "adjust points",
  "CharacterID": 2,
  "OwnerAction": false,
  "Costs": -10,
  "SessionID": "c68d83ff-f409-4313-8aa9-345af4e559f1"
}'`

Would return: 

_HTTP/1.1 200 OK  
Content-Type: application/json  
Date: Tue, 23 Jul 2024 11:18:01 GMT  
Content-Length: 36_

```json
{
    "dice_roll":2,
    "current_points":92
}
```

You spent 10 points but the dice gifted you with a 2. So you have 92 Points.

### Pull notification 

`curl -i -X GET "http://localhost:8080/api/v1/fatecore/session/c68d83ff-f409-4313-8aa9-345af4e559f1/notification?charID=1&offset=0"`

Would return: 

_HTTP/1.1 200 OK  
Content-Type: application/json  
Date: Tue, 23 Jul 2024 20:08:33 GMT  
Content-Length: 149_

```json
[
    {
        "Body":"MTAwIHBvaW50cyBzcGVudA==",
        "CreatedAt":"2024-07-23T22:07:39.610330771+02:00",
        "FromId":1,
        "SessionId":"c68d83ff-f409-4313-8aa9-345af4e559f1"
    }
]
```

The body is in base64.
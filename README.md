##### Go JWT Auth

### Routes
```
/ [GET]
/login [POST] 
/sign-up [POST]
/validate [GET]
```

1. / [GET]
Response
```
{
    "author": "Nidhey Indurkar",
    "data": [],
    "error": "",
    "message": "GO JWT Authentication - POSTGRES API",
    "status": true
}
```

2. /login [POST] 
Body:
```
{
    "Email": "nidhey29@gmail.com",
    "Password": "12345677"
}
```

Cookie: 
```
Authorization - xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```
Response
```
{
    "data": [],
    "error": "",
    "message": "Login Successful!",
    "status": true
}
```

3. /sign-up [POST]
Body
```
{
    "Email": "nidhey40@gmail.com",
    "Password": "12345677"
}
```

Response
```
{
    "data": {
        "ID": 5,
        "CreatedAt": "2022-11-22T05:16:42.606661644Z",
        "UpdatedAt": "2022-11-22T05:16:42.606661644Z",
        "DeletedAt": null,
        "Email": "nidhey40@gmail.com",
        "Password": "$2a$10$qjXGH0AAs4dMwKwhLlZB.e8TNJwq.LRLxfiu3h7W5z91CseLSeFYu"
    },
    "error": "",
    "message": "User Created!",
    "status": true
}
```

4. /validate [GET]
Cookie: 
```
Authorization - xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```
Response
```
{
    "data": [],
    "error": "",
    "message": "Authorized",
    "status": true
}
```
or
```
{
    "data": [],
    "error": "",
    "message": "UnAuthorized",
    "status": true
}
```
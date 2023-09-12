
# Go-Blog-Backend-API

Personal Project Making Backend Blog Website with golang


## Run Locally

Clone the project

```bash
  git clone https://github.com/Wazven/golang-blog.git
```

Go to the project directory

```bash
  cd backend
```

Install dependencies

```bash
  go mod download
```

Start the server

```bash
  go run main.go
```


## API Documentation
- [Register](#register)
- [Login](#login)
- [blog](#blog)

# API Reference

# Register
- Endpoint :
    - /blog/register
- Method :
    - POST
- Body :
```json
{
    "first_name":"string, required",
    "last_name":"string, required",
    "email":"string, email, required",
    "password":"string, min:6, required",
    "phone":"number, min:11, required"
}
```
- Response :
```json
{
    "message": "Akun Berhasil Dibuat",
    "user": {
        "id": 3,
        "first_name": "Testing",
        "last_name": "Test",
        "email": "test1234@gmail.com",
        "phone": "08289018215678"
    }
}
```

# Login
- Endpoint :
    - /blog/Login
- Method :
    - POST
- Body :
```json
{
    "email":"string, email, required",
    "password":"string, min:6, required",
}
```
- Response :
```json
    {
    "message": "Berhasil Login",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ2MjE2NjYsImlzcyI6IjMifQ.xBw8YcVLMvUyMAXB4PR29DSX6ggQ_b_Rnvf6oBUMahQ",
    "user": {
        "id": 3,
        "first_name": "Testing",
        "last_name": "Test",
        "email": "test1234@gmail.com",
        "phone": "08289018215678"
    }
```
# Blog

## POST Blog
- Endpoint :
    - /blog/post
- Method :
    - POST
- Headers :
    - Authorization: Bearer
- Body : 
```json
{
    "title":"Test Post Blog",
    "description":"Lorem Ipsum Dolor Sit Amet",
    "image":"https://picsum.photos/200/300",
    "userid":"1"
}
```
- Response :
```json
{
    "message": "Berhasil Posting"
}
```
## GET All Blog
- Endpoint :
    - /blog/post
- Method :
    - GET
- Headers :
    - Authorization: Bearer
- Response :
```json
{
    "data": [
        {
            "id": 4,
            "title": "Postingan Ngga Tau",
            "description": "Lorem Ipsum Dolor Sit Amet",
            "image": "https://picsum.photos/200/300",
            "userid": "2",
            "user": {
                "id": 2,
                "first_name": "Jimmy",
                "last_name": "Brully",
                "email": "somethingidk@gmail.com",
                "phone": "082791074091"
            }
        },
        {
            "id": 5,
            "title": "Postingan yang ada apa-apanya",
            "description": "Lorem Ipsum Dolor Sit Amet kokwokwokwko",
            "image": "https://picsum.photos/200/300/400/500",
            "userid": "1",
            "user": {
                "id": 1,
                "first_name": "Aldi",
                "last_name": "Rhiyadi",
                "email": "fuckingfake@gmail.com",
                "phone": "082791074092"
            }
        }
}
```
## Get Detail Blog
- Endpoint :
    - /blog/post
- Method :
    - GET {id}
- Headers :
    - Authorization: Bearer
- Response :
```json
{
      "data": {
        "id": 7,
        "title": "Test Post Blog",
        "description": "Lorem Ipsum Dolor Sit Amet",
        "image": "https://picsum.photos/200/300",
        "userid": "3",
        "user": {
            "id": 3,
            "first_name": "Testing",
            "last_name": "Test",
            "email": "Testing123@gmail.com",
            "phone": "082791074092"
        }
    }
}
```

## UPDATE Blog
- Endpoint :
    - /blog/post
- Method :
    - PUT {id}
- Headers :
    - Authorization: Bearer
- Body :
```json
{
        "title": "Postingan yang ada apa-apanya",
        "description": "Lorem Ipsum Dolor Sit Amet kokwokwokwko",
        "image": "https://picsum.photos/200/300/400/500",
        "userid": "3"
}
```
- Response :
```json
{
    "message": "Edit Posting Berhasil"
}
```
## DELETE Blog
- Endpoint :
    - /blog/post
- Method :
    - DELETE {id}
- Headers :
    - Authorization: Bearer
- Response :
```json
{
    "message": "Postingan Berhasil Dihapus"
}
```
# Image
## POST Image
- Endpoint :
    - /blog/imagepost
- Method :
    - POST
- Headers :
    - Content-Type: multipart/form-data
    - Authorization: Bearer
- Body :
    - Photo as File (Image must be jpg, jpeg, or png format)
- Response :
```json
{
    "url": "http://localhost:8000/blog/uploads/ajkqop-konz.png"
}
```


# Spaceships
### Creating a galactic database (using MySQL) that stores all the spacecraft’s details


## Getting Started

### Prerequistes 

To start using this application these prerequistes must be followed sequentially
- Install **[Go](https://go.dev/)**: (https://go.dev/doc/devel/release) (now version **1.16+** is required).


## Usage
### .env File
```sh 
DB_HOST="127.0.0.1"
DB_USER = "root"
DB_PASSWORD = "*****"
DB_NAME = "galactic"
DB_PORT = "3306"
```
```sh
$ go mod tidy
```

- Compile the main application

```sh
$ make run
```


## GET Get All Spaceships
```sh
localhost:3300/spaceships
```

### Body 
```sh

[
  {
    "id": 1,
    "name": "Devastator",
    "class": "Star Destroyer",
    "crew": 35000,
    "image": "https:\\url.to.image",
    "value": 1999.99,
    "status": "operational",
    "armaments": [
        {
            "title": "Turbo Laser",
            "qty": "60"
        },
        {
          "title": "Ion Cannons",
          "qty": "60"
        },
        {
          "title": "Tractor Beam",
          "qty": "10"
        }
    ]
  },
  {
    "id": 2,
    "name": "Red Five",
    "class": "Boeing 3745",
    "crew": 45000,
    "image": "https:\\url.to.image",
    "value": 3999.99,
    "status": "damaged",
    "armaments": [
        {
            "title": "Turbo Laser",
            "qty": "60"
        },
        {
          "title": "Ion Cannons",
          "qty": "60"
        },
        {
          "title": "Tractor Beam",
          "qty": "10"
        }
    ]
  },
  {
    "id": 3,
    "name": "High Five",
    "class": "Boeing 3745",
    "crew": 45000,
    "image": "https:\\url.to.image",
    "value": 3999.99,
    "status": "damaged",
    "armaments": [
        {
            "title": "Turbo Laser",
            "qty": "60"
        },
        {
          "title": "Ion Cannons",
          "qty": "60"
        },
        {
          "title": "Tractor Beam",
          "qty": "10"
        }
    ]
  },
]
```
## GET Get Spaceship
```sh
localhost:3300/spaceship/id
```

### Body 
```sh

[
  {
    "id": 1,
    "name": "Devastator",
    "class": "Star Destroyer",
    "crew": 35000,
    "image": "https:\\url.to.image",
    "value": 1999.99,
    "status": "operational",
    "armaments": [
        {
            "title": "Turbo Laser",
            "qty": "60"
        },
        {
          "title": "Ion Cannons",
          "qty": "60"
        },
  }
]
```

# POST Create Spaceship
```sh
localhost:3300/createspaceship
```

# DELETE Delete Spaceship
```sh
localhost:3300/spaceship/delete/id
```

# UPDATE Update Spaceship
```sh
localhost:3300/spaceship/update/id
```
# FILTER Filter By Name, Class, Status
```sh
localhost:3300/spaceships?name=ActualName
```
```sh
localhost:3300/spaceships?class=ActualClass
```
```sh
localhost:3300/spaceships?status=ActualStatus
```



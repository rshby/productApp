# Product App
## how to run
clone project ini kemudian buka menggunakan terminal/cmd. Arahkan ke direktori root foler. <br>
Selanjutnya ketikkan `docker compose up -d` <br>
Maka aplikasi akan secara otomatis menjalankan beberapa container yang diperlukan. <br>
Jika aplikasi tidak berjalan, bisa dicoba untuk dijalankan manual. hal ini terjadi karena aplikasi membutuhkan service lain, tetapi service lain terebut belum siap dijalankan.

## Database
Database yang digunakan adalah MySQL yang otomatis jalan menggunakan container docker. <br>
Jika ingin membuka koneksi dengan database, dapat menggunakan configurasi berikut. <br>
```
user = root
password = root
host = localhost
port = 3306
```
Database container yang dijalankan juga sudah berisi beberapa data yang siap digunakan untuk menjalankan apliasi.

## Endpoint API
### List Of Endpoint
| Endpoint  | Method |
| ------------- |:------:|
| http://localhost:5005/api/v1/product|  POST  |
| http://localhost:5005/api/v1/products?order=desc&sort=price|  GET   |<br>

## Detail Endpoint
### 1. Add Product
Method `POST`, enpoint `http://localhost:5005/api/v1/product` <br>
Request :
```
{
    "name" : "Mac Mini M2 Pro",
    "price" : 22000000,
    "description" : "Mac Mini Terbaru",
    "quantity": 100
}
``` 
Response :
```azure
{
    "status_code": 200,
    "status": "ok",
    "message": "success insert",
    "data": {
        "id": 5,
        "name": "Mac Mini M2 Pro",
        "price": 22000000,
        "description": "Mac Mini Terbaru",
        "quantity": 100,
        "created_at": "2024-02-07 03:59:00",
        "updated_at": "2024-02-07 03:59:00"
    }
}
```
### 2. Get List Products And Sorting
Terdapat beberapa pilihan sorting : `name, price, date` <br>
Secara default apabila parameter `sort` dan `order` tidak diisi, makan akan mengurutkan berdasarkan date DESC (product yang paling baru). <br>
Pilihan option untuk parameter `sort` adalah : `date`, `price`, `name` <br>
Pilihan option utuk parameter `order` adalah : `asc` atau `desc` <br>
Method `GET` endpoint `http://localhost:5005/api/v1/products?order=desc&sort=price`<br>
Response :
```
{
    "status_code": 200,
    "status": "ok",
    "message": "success get products",
    "data": [
        {
            "id": 4,
            "name": "Macbook Pro M3 Pro 18/1TB",
            "price": 44999000,
            "description": "Macbook Pro Terbaru",
            "quantity": 100,
            "created_at": "2023-10-30 10:00:00",
            "updated_at": "2023-10-30 10:00:00"
        },
        {
            "id": 3,
            "name": "Macbook Pro M2 Pro 16/1TB",
            "price": 34999000,
            "description": "Macbook tahun lalu",
            "quantity": 100,
            "created_at": "2022-10-10 10:00:00",
            "updated_at": "2022-10-10 10:00:00"
        },
        {
            "id": 1,
            "name": "iPhone 15 Pro",
            "price": 22999000,
            "description": "iPhone terbaru",
            "quantity": 100,
            "created_at": "2024-02-07 10:00:00",
            "updated_at": "2024-02-07 10:00:00"
        },
        {
            "id": 2,
            "name": "iPhone 14 Pro",
            "price": 16999000,
            "description": "iPhone tahun lalu",
            "quantity": 100,
            "created_at": "2022-09-22 10:00:00",
            "updated_at": "2022-09-22 10:00:00"
        }
    ]
}
```

## Arsitektur
Arsitektur code (_design pattern_) yang digunakan dalam source code aplikasi ini adalah clean code yang menggunakan beberapa layer seperti repository, service, dan handler.<br>
Sedangkan untuk arsitektur sistem yang digunakan adalah _microservices_ dan dapat dijalankan dari beberapa node dalam load balancer. Sehingga dapat menangani traffic request yang tinggi dan menjamin _high availability_. <br>
Gambar Arsitektur yang digunakan dapat dilihat pada gambar di bawah ini : <br>
[<img src="https://drive.google.com/uc?export=view&id=16E__3ly5yAM4a54WSviNY9Q3FPftks2t" width="300"/>](https://drive.google.com/uc?export=view&id=16E__3ly5yAM4a54WSviNY9Q3FPftks2t) <br>
[<img src="https://drive.google.com/uc?export=view&id=1xQk1LO8CnqnXjGKf1bbOidXwud9iqFPS" width="400"/>](https://drive.google.com/uc?export=view&id=1xQk1LO8CnqnXjGKf1bbOidXwud9iqFPS)

## Unit Test
Aplikasi ini juga mendukung Unit Test. <br>
Untuk menjalankan semua unit test dapat menggunakan syntax `go test -cover -v ./test`

## Tracing
Aplikasi ini mendukung tracing menggunakan jaeger. <br>
Anda dapat membua dashboar jaeger dengan url `http://localhost:16686` <br>
Dengan menggunakan tracing ini kita akan lebih mudah mengetahui masing-masing proses yang jalan di apliasi ini. <br>
Harapannya penambahan tracing ini akan mempermudah kita dalam menemukan root cause ketika apliasi ini dirasa memiliki response time yang lama.<br>
[<img src="https://drive.google.com/uc?export=view&id=1HntRZ0ShdjtDxGK-MFXWbS3rpIKbv-PN" width="400"/>](https://drive.google.com/uc?export=view&id=1HntRZ0ShdjtDxGK-MFXWbS3rpIKbv-PN) <br>
[<img src="https://drive.google.com/uc?export=view&id=1xyfRtxGjN1nIXq9e87c3qPQN5f_dpVPI" width="400"/>](https://drive.google.com/uc?export=view&id=1xyfRtxGjN1nIXq9e87c3qPQN5f_dpVPI)

## Log With ELK (Filebeat - Logstash - Elastic - Kibana)
Aplikasi ini juga mendukung logging yang ditampilkan dari dahshboard Kibana.<br>
Anda dapat membuka dashboard Kibana dengan url berikut `http://localhost:5601` <br>
Dengan menggunakan filebeat-logstash-elasticsearch-kibana harapannya kita akan lebih mudah dalam melihat log untuk setiap request yang masuk ke aplikasi ini<br>
[<img src="https://drive.google.com/uc?export=view&id=1Glz89WQPzp7fYhiRDhbTcK6xWkS-eS1H" width="400"/>](https://drive.google.com/uc?export=view&id=1Glz89WQPzp7fYhiRDhbTcK6xWkS-eS1H)
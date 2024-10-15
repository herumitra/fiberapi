## Tentang Fiber API
`(Golang Fiber Restful API)`
Dalam repositori ini kita menerapkan `Golang` sebagai platform dasar bahasa pemrograman yang digunakan dalam pembuatan `API`.
Di dalam repositori ini juga kami terapkan framework `Fiber` serta dependensi `GORM` dan `JWT` untuk mempermudah dalam pengerjaan di ranah sekuritas maupun pengelolaan databasenya, sehingga detail komponen yang kami gunakan bisa dijabarkan seperti berikut ini :
| NO. | KOMPONEN       |
|-----:|---------------|
|     1| Fiber         |
|     2| GORM          |
|     3| MySQL         |
|     4| JWT           |


>Tambahkan file `.env` dalam direktori paling luar dari project di repositori ini, dan masukkan teks berikut di dalamnya.

```bash
JWT_SECRET=123456qwerty
DB_USER=root
DB_PASSWORD=root
DB_NAME=fiber
DB_PORT=3306
DB_SERVER=localhost
SERVER_PORT=6969
```
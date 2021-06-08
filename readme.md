# POST COLLECTION
```bash
    #you can import post collection with this link
    https://www.getpostman.com/collections/2083b73af33a0130892e

```
# Cara Menggunakan unit testing
```bash
    - Jalankan Terlebih dahulu main.go
    - Masuk kedalam folder test 
    - kemudian Ketikan perintah berikut :
        - go test -v -run=TestTambahProduk
        - go test -v -run=TestTampilProduk
        - go test -v -run=TestHapusProduk

```


# Api Spesification

## Tambah Produk

Request :
- Method : POST
- Endpoint : `/tambah_produk`
- Body :
    Form-Data
```json 
{
    params[kode_produk]:Pisang Hijau
    params[kuantitas]:2
}
```

Response :

```json 
{
    {
    "response": {
        "KodeProduk": "Pisang Hijau",
        "Kuantitas": 2,
        "TotalProduk": 2
    },
    "status": 200
    }
}
```

## Tampil Produk

Request :
- Method : POST
- Endpoint : `/tampil_produk`
- Body :


Response :

```json 
{
    {
    "response": [
        {
            "Id": 3,
            "KodeProduk": "Aman Roti",
            "Kuantitas": 4
        },
        {
            "Id": 4,
            "KodeProduk": "Melon",
            "Kuantitas": 1
        }
    ],
    "status": 200
    }
}
```
## Hapus Produk

Request :
- Method : POST
- Endpoint : `/update_loan`
- Body :

```json 
{
    params[kode_produk]:Pisang Hijau
}
```

Response :

```json 
{
    {
    "response": {
        "KodeProduk": "Aman Marbel",
        "Message": "Berhasil di hapus"
    },
    "status": 200
    }
}
```

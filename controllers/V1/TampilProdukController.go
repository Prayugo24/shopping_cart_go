package V1

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
	
)

type V1TampilProdukController struct {
	Status int
}

type ProdukData struct {
	ID   int `json:"id"`
    Kode_produk string `json:"kode_produk"` 
	Kuantitas int `json:"kuantitas"` 
}

type ResponseTampilProduk struct {
	Id   int 
	KodeProduk string 
	Kuantitas int 
}

func (status *V1TampilProdukController) TampilProduk (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	ProdukDb, _ := db.Table("tb_produks").Rows()
	defer ProdukDb.Close()

	resultTampilProduk := ResponseTampilProduk{}
    result := []ResponseTampilProduk{}
    for ProdukDb.Next() {
		var id, kuantitas int
        var kode_produk string
        ProdukDb.Scan(&id, &kode_produk, &kuantitas)
        resultTampilProduk.Id = id
        resultTampilProduk.KodeProduk = kode_produk
        resultTampilProduk.Kuantitas = kuantitas
        result = append(result, resultTampilProduk)
    }
	fmt.Println(result)

	c.JSON(200, gin.H{"status": 200,"response":result})
	return
	
	

	
}



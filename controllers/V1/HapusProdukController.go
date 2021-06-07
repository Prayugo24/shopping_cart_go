package V1

import(
	"github.com/gin-gonic/gin"
	"Tes/sirclo_shopping_cart/models"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"strings"
	
)

type V1HapusProdukController struct {
	Status int
}

type ResponseHapusProduk struct {
	KodeProduk string 
	Message string
}

func (status *V1HapusProdukController) HapusProduk (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	kodeProduk := c.DefaultPostForm("params[kode_produk]", "")
	if kodeProduk == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Kode Produk Harus Di Isi",
		})
		return
	}
	kodeProdukConvert := strings.Title(kodeProduk)
	ProdukDb, _ := db.Table("tb_produks").Where("kode_produk = ?", kodeProdukConvert).Rows()
	defer ProdukDb.Close()
    // // Values to load into
    result := make([]ProdukData, 0)
    for ProdukDb.Next() {
        var row ProdukData
        result = append(result, row)
		
    }
	if len(result) <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data Tidak di Temukan, atau data sudah tidak ada",
		})
		return
	}else{
		kodeProdukConvert := strings.Title(kodeProduk)
		var produkModels models.Tb_Produk
		db.Where("kode_produk = ?", kodeProdukConvert).Delete(&produkModels)
		result := ResponseHapusProduk{
			KodeProduk: kodeProdukConvert,
			Message : "Berhasil di hapus",
		}
		fmt.Println(result)
		c.JSON(200, gin.H{"status": 200,"response":result})
		return
	}
	
	
}



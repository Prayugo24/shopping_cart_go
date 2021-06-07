package V1

import(
	"github.com/gin-gonic/gin"
    "Tes/sirclo_shopping_cart/models"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"strings"
	"strconv"
	
)

type V1TambahProdukController struct {
	Status int
}

type ResponseTambahProduk struct {
	KodeProduk string 
	Kuantitas int 
	TotalProduk int
}
type ResponseProduk struct {
	Id   int 
	KodeProduk string 
	Kuantitas int 
}


func (status *V1TambahProdukController) TambahProduk (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	kodeProduk := c.DefaultPostForm("params[kode_produk]", "")
	kuantitas, err := strconv.Atoi(c.PostForm("params[kuantitas]"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Kuantitas Harus di isi angka",
		})
		return
	}
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
	var resultTampilProduk ResponseProduk
    for ProdukDb.Next() {
		db.ScanRows(ProdukDb, &resultTampilProduk)
        var row ProdukData
        result = append(result, row)
		
    }

	if len(result) > 0 {
		var produkModels models.Tb_Produk
		totalProduk := resultTampilProduk.Kuantitas + kuantitas
		fmt.Println(totalProduk)

		db.Model(&produkModels).Where("id = ?", resultTampilProduk.Id).Update("kuantitas", totalProduk)
		
		response := ResponseTambahProduk {
			KodeProduk:resultTampilProduk.KodeProduk,
			Kuantitas: kuantitas,
			TotalProduk : totalProduk,
		}
		c.JSON(200, gin.H{"status": 200,"response":response})
		return
	}else{
		inputProduk := models.Tb_Produk{
			Kode_produk : kodeProdukConvert,
			Kuantitas : kuantitas,
		}
		db.Create(&inputProduk)

		response := ResponseTambahProduk {
			KodeProduk:kodeProdukConvert,
			Kuantitas: kuantitas,
			TotalProduk : kuantitas,
		}
		c.JSON(200, gin.H{"status": 200,"response":response})
		return
	}
	
	

	
}



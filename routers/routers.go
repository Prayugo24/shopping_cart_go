package routers

import (
	"github.com/gin-gonic/gin"
	"Tes/sirclo_shopping_cart/config"
	"net/http"
	"Tes/sirclo_shopping_cart/controllers/V1"
)

func RouterMain() http.Handler  {
	router := gin.New()
	db := config.SetupModels()

	router.Use(func(c * gin.Context){
		c.Set("db",db)
		c.Next()
	})
	v1TambahProduk := &V1.V1TambahProdukController{Status: 200}
	v1TampilProduk := &V1.V1TampilProdukController{Status: 200}
	v1HapusProduk := &V1.V1HapusProdukController{Status: 200}
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"data":"Welcome To Api Golang"})
	})

	router.POST("/tambah_produk", v1TambahProduk.TambahProduk)
	router.POST("/tampil_produk", v1TampilProduk.TampilProduk)
	router.POST("/hapus_produk", v1HapusProduk.HapusProduk)
	

	

	return router
}

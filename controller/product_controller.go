package controller

import (
	"go-api/model"
	_ "go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
		
}

func (p *productController) GetProducts(ctx *gin.Context) {
	// TODO: implement logic to get products
	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, products)

	
}

func (p *productController) CreateProducts(ctx *gin.Context) {
	
	var product model.Product
		err := ctx.BindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err) 
			return
		}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response {
			Message: "Id do produto precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}


	product, err := p.ProductUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response {
			Message: "Id do produto não encontrado na base de dados.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}
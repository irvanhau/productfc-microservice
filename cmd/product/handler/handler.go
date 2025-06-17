package handler

import (
	"fmt"
	"net/http"
	"productfc/cmd/product/usecase"
	"productfc/infrastructure/log"
	"productfc/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (h *ProductHandler) ProductCategoryManagement(c *gin.Context) {
	var param models.ProductCategoryManagementParameter
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Error(err.Error()) // Untuk Debugging
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Input",
		})
		return
	}

	if param.Action == "" {
		log.Logger.Error("missing parameter action")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Missing required parameter",
		})
		return
	}

	switch param.Action {
	case "add":
		if param.ID != 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product category id is not empty")
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid Request",
			})
			return
		}

		productCategoryID, err := h.ProductUsecase.CreateNewProductCategory(c.Request.Context(), &param.ProductCategory)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.CreateNewProductCategory got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("Successfully create new product category: %d", productCategoryID),
		})
		return
	case "edit":
		if param.ID == 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product id is empty")

			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid Request",
			})
			return
		}

		productCategory, err := h.ProductUsecase.EditProductCategory(c.Request.Context(), &param.ProductCategory)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.EditProductCategory got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success Edit Product",
			"product": productCategory,
		})
		return
	case "delete":
		if param.ID == 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product id is empty")
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid Request",
			})
			return
		}

		err := h.ProductUsecase.DeleteProductCategory(c.Request.Context(), param.ID)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param, // notes: kalau ada data pribadi --> prevent print log data pribadi
			}).Errorf("h.ProductUsecase.DeleteProductCategory got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Product Category ID %d Successfully Deleted!", param.ID),
		})
		return
	default:
		log.Logger.Error("Invalid Action")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Action",
		})
		return
	}
}

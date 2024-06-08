package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-sqlx/domain"
	"go-gin-sqlx/usecase"
	"strconv"
)

type pegawaiController struct {
	uc usecase.PegawaiUsecase
}

func NewPegawaiController(pegawaiUsecase usecase.PegawaiUsecase) *pegawaiController {
	return &pegawaiController{pegawaiUsecase}
}

func (c pegawaiController) GetAllPegawai(ctx *gin.Context) {
	pegawai, err := c.uc.FindAllPegawai()
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "Failed to get all pegawai",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "OK",
		"message": "Success to get all pegawai",
		"data":    pegawai,
	})
}

func (c pegawaiController) FindPegawaiByid(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "Params not valid",
		})
		return
	}
	findPegawai, err := c.uc.FindPegawaiById(id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "Pegawai not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "OK",
		"message": "Success to get all pegawai",
		"data":    findPegawai,
	})
}

func (c pegawaiController) CreatePegawai(ctx *gin.Context) {
	var pegawaiReq domain.PegawaiRequest
	err := ctx.ShouldBindJSON(&pegawaiReq)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newPegawai, err := c.uc.CreatePegawai(&pegawaiReq)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "Created",
		"message": newPegawai,
	})
}

func (c pegawaiController) UpdatePegawai(ctx *gin.Context) {
	var pegawaiReq domain.PegawaiRequest

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "Param not valid",
		})
		return
	}
	err = ctx.ShouldBindJSON(&pegawaiReq)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	pegawai, err := c.uc.UpdatePegawai(id, &pegawaiReq)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "OK",
		"message": "Success to update pegawai",
		"data":    pegawai,
	})
}

func (c pegawaiController) DeletePegawai(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "Params not valid",
		})
		return
	}
	err = c.uc.DeletePegawai(id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "OK",
		"message": "Success to delete pegawai",
	})
}

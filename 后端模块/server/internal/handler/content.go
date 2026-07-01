package handler

import (
	"net/http"
	"strconv"

	"ovo-gaming/server/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MiniHome(c *gin.Context) {
	var banners []model.Banner
	var notices []model.Notice
	var packages []model.Package

	if err := h.DB.Where("status = ?", "on").Order("sort DESC, id DESC").Find(&banners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Where("status = ?", "on").Order("sort DESC, id DESC").Limit(1).Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Where("status = ? AND is_hot = ?", "on", true).Order("sort DESC, id DESC").Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"banners":  banners,
		"notice":   firstNotice(notices),
		"packages": packages,
	})
}

func (h *Handler) AdminListBanners(c *gin.Context) {
	var banners []model.Banner
	if err := h.DB.Order("sort DESC, id DESC").Find(&banners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": banners})
}

func (h *Handler) AdminCreateBanner(c *gin.Context) {
	var input model.Banner
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"item": input})
}

func (h *Handler) AdminUpdateBanner(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input model.Banner
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Model(&model.Banner{}).Where("id = ?", id).Updates(map[string]any{
		"title":       input.Title,
		"subtitle":    input.Subtitle,
		"image_url":   input.ImageURL,
		"jump_type":   input.JumpType,
		"jump_target": input.JumpTarget,
		"status":      input.Status,
		"sort":        input.Sort,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) AdminListNotices(c *gin.Context) {
	var notices []model.Notice
	if err := h.DB.Order("sort DESC, id DESC").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": notices})
}

func (h *Handler) AdminCreateNotice(c *gin.Context) {
	var input model.Notice
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"item": input})
}

func (h *Handler) AdminListPackages(c *gin.Context) {
	var packages []model.Package
	if err := h.DB.Order("sort DESC, id DESC").Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": packages})
}

func (h *Handler) AdminCreatePackage(c *gin.Context) {
	var input model.Package
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"item": input})
}

func (h *Handler) AdminUpdatePackage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input model.Package
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.DB.Model(&model.Package{}).Where("id = ?", id).Updates(map[string]any{
		"game_code":        input.GameCode,
		"title":            input.Title,
		"price":            input.Price,
		"original_price":   input.OriginalPrice,
		"duration_minutes": input.DurationMinutes,
		"description":      input.Description,
		"cover_image":      input.CoverImage,
		"is_hot":           input.IsHot,
		"status":           input.Status,
		"sort":             input.Sort,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func firstNotice(items []model.Notice) gin.H {
	if len(items) == 0 {
		return gin.H{}
	}
	return gin.H{
		"id":      items[0].ID,
		"title":   items[0].Title,
		"content": items[0].Content,
		"type":    items[0].NoticeType,
		"sort":    items[0].Sort,
	}
}

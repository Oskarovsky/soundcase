package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"net/http"
	"soundcase/models"
)

func ReadArtist(c *gin.Context) {
	var artist models.Artist
	id := c.Param("id")
	res := initializers.DB.Find(&artist, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "artist not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"artist": artist,
	})
	return
}

func ReadArtists(c *gin.Context) {
	var artists []models.Artist
	res := initializers.DB.Find(&artists)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("artists not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"artists": artists,
	})
	return
}

func CreateArtist(c *gin.Context) {
	var artist *models.Artist
	err := c.ShouldBind(&artist)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := initializers.DB.Create(artist)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating an artist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"artist": artist,
	})
	return
}

func UpdateArtist(c *gin.Context) {
	var artist models.Artist
	id := c.Param("id")
	err := c.ShouldBind(&artist)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateArtist models.Artist
	res := initializers.DB.Model(&updateArtist).Where("id = ?", id).Updates(artist)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "artist not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"artist": artist,
	})
	return
}

func DeleteArtist(c *gin.Context) {
	var artist models.Artist
	id := c.Param("id")
	res := initializers.DB.Find(&artist, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "artist not found",
		})
		return
	}
	initializers.DB.Delete(&artist)
	c.JSON(http.StatusOK, gin.H{
		"message": "artist deleted successfully",
	})
	return
}

package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"net/http"
	"soundcase/models"
)

func ReadTrack(c *gin.Context) {
	var track models.Track
	id := c.Param("id")
	res := initializers.DB.Find(&track, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "track not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"track": track,
	})
	return
}

func ReadTracks(c *gin.Context) {
	var tracks []models.Track
	res := initializers.DB.Find(&tracks)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("tracks not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tracks": tracks,
	})
	return
}

func CreateTrack(c *gin.Context) {
	var track *models.Track
	err := c.ShouldBind(&track)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := initializers.DB.Create(track)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a track",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"track": track,
	})
	return
}

func UpdateTrack(c *gin.Context) {
	var track models.Track
	id := c.Param("id")
	err := c.ShouldBind(&track)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateTrack models.Track
	res := initializers.DB.Model(&updateTrack).Where("id = ?", id).Updates(track)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "track not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"track": track,
	})
	return
}

func DeleteTrack(c *gin.Context) {
	var track models.Track
	id := c.Param("id")
	res := initializers.DB.Find(&track, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "track not found",
		})
		return
	}
	initializers.DB.Delete(&track)
	c.JSON(http.StatusOK, gin.H{
		"message": "track deleted successfully",
	})
	return
}

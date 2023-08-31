package mahasiswacontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"GO-RESTAPI-GIN/databases"
	"GO-RESTAPI-GIN/models"

	ws "GO-RESTAPI-GIN/websocket"


)

//	@Summary		Show all students
//	@Description	get all students data
//	@ID				get-all-students
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Mahasiswa
//	@Router			/mahasiswas [get]
func Index(c *gin.Context) {

	var mahasiswas []models.Mahasiswa

	databases.DB.Find(&mahasiswas)
	c.JSON(http.StatusOK, gin.H{"mahasiswas": mahasiswas})
	ws.SendWebSocketUpdate("get all student!")

}

//	@Summary		Show a student
//	@Description	get a students data by id
//	@ID				get-a-student-by-id
//	@Tags			students
//	@Accept			json
//	@Produce		json
//  @Param id path int true "ID of the student"
//	@Success		200	{array}	models.Mahasiswa
//	@Failure		400	{string} http.StatusNotFound
//	@Failure		404	{string} http.StatusNotFound
//	@Router			/mahasiswas/{id} [get]
func Show(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := databases.DB.First(&mahasiswa, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Mahasiswa": mahasiswa})
	ws.SendWebSocketUpdate("get student by id!")
}

//	@Summary		create a student
//	@Description	create a student
//	@ID				create-a-student
//	@Tags			students
//	@Accept			json
//	@Produce		json
//  @Param	student	  {object}   body models.RequestMahasiswa true "data of the student"
//	@Success		200	{array}	models.Mahasiswa
//	@Failure		400	{string} http.StatusBadRequest
//	@Router			/mahasiswas [post]
func Create(c *gin.Context) {

	var mahasiswa models.Mahasiswa

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	databases.DB.Create(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"mahasiswas": mahasiswa})
	ws.SendWebSocketUpdate("create student has successfuly!")
}

//	@Summary		update a student
//	@Description	update a student data by id
//	@ID				update-a-student
//	@Tags			students
//	@Accept			json
//	@Produce		json
//  @Param id path int true "ID of the student"
//  @Param	student	  {object}   body models.RequestMahasiswa true "data of the student"
//	@Success		200	{string}  string "Successfully update student by id"
//	@Failure		400	{string} http.StatusBadRequest
//	@Router			/mahasiswas/{id} [put]
func Update(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if databases.DB.Model(&mahasiswa).Where("id = ?", id).Updates(&mahasiswa).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update student by id"})
	ws.SendWebSocketUpdate("successfully update student!")

}

//	@Summary		delete a student
//	@Description	delete a student data by id
//	@ID				delete-a-student
//	@Tags			students
//	@Accept			json
//	@Produce		json
//  @Param id path int true "ID of the student"
//	@Success		200	{string}  string "Successfully update student by id"
//	@Failure		400	{string} string "student not found"
//	@Failure		500	{string} string "Failed to delete Mahasiswa"
//	@Router			/mahasiswas/{id} [delete]
func Delete(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	id := c.Param("id")

	if err := databases.DB.First(&mahasiswa, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "student not found"})
		return
	}

	if err := databases.DB.Delete(&mahasiswa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete Mahasiswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student successfully deleted"})
	ws.SendWebSocketUpdate("successfully deleted!")
}

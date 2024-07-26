package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kundu-ramit/zocket/service"
)

type KeyValueController interface {
	CreateAuth(c *gin.Context)
	CreateKeyValue(c *gin.Context)
	GetKeyValue(c *gin.Context)
	UpdateKeyValue(c *gin.Context)
	DeleteKeyValue(c *gin.Context)
}

type keyValueController struct {
	service service.KeyValueService
}

func NewKeyValueController() KeyValueController {
	return &keyValueController{
		service: service.NewKeyValueService(),
	}
}

func (kv *keyValueController) CreateAuth(c *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := kv.service.CreateAuth(requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth)
}

func (kv *keyValueController) CreateKeyValue(c *gin.Context) {
	var requestBody struct {
		Key     string    `json:"key"`
		Value   string    `json:"value"`
		OwnedBy uuid.UUID `json:"ownedBy"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kvPair, err := kv.service.CreateKeyValue(requestBody.Key, requestBody.Value, requestBody.OwnedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kvPair)
}

func (kv *keyValueController) GetKeyValue(c *gin.Context) {
	key := c.Param("key")
	ownedBy := c.Query("ownedBy")

	ownerUUID, err := uuid.Parse(ownedBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	kvPair, err := kv.service.GetKeyValue(key, ownerUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, kvPair)
}

func (kv *keyValueController) UpdateKeyValue(c *gin.Context) {
	key := c.Param("key")
	var requestBody struct {
		Value   string    `json:"value"`
		OwnedBy uuid.UUID `json:"ownedBy"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kvPair, err := kv.service.UpdateKeyValue(key, requestBody.Value, requestBody.OwnedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, kvPair)
}

func (kv *keyValueController) DeleteKeyValue(c *gin.Context) {
	key := c.Param("key")
	ownedBy := c.Query("ownedBy")

	ownerUUID, err := uuid.Parse(ownedBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := kv.service.DeleteKeyValue(key, ownerUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key deleted"})
}

package handler

import (
	"github.com/gin-gonic/gin"
)

type EncodeRequest struct {
	Url string `json:"url"`
}

type EncodeResponse struct {
	Url string `json:"url"`
}

func (h Handler) encodeHandler(ctx *gin.Context) {
	// Bind the JSON payload to a User struct.
	request := &EncodeRequest{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	url, err := h.repo.Save(ctx, request.Url)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Return a success message.
	ctx.JSON(200, &EncodeResponse{Url: url})

}

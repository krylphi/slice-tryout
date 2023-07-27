package handler

import (
	"github.com/gin-gonic/gin"
)

type DecodeRequest struct {
	Url string `json:"url"`
}

type DecodeResponse struct {
	Url string `json:"url"`
}

func (h Handler) decodeHandler(ctx *gin.Context) {
	// Bind the JSON payload to a User struct.
	request := &DecodeRequest{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	url, err := h.repo.Load(ctx, request.Url)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Return a success message.
	ctx.JSON(200, &DecodeResponse{Url: url})

}

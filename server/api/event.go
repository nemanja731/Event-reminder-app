package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nemanja731/Event-reminder-web/server/db/sqlc"
	"github.com/nemanja731/Event-reminder-web/server/token"
)

type AddEventRequest struct {
	Title     string    `json:"title" binding: "required"`
	EventTime time.Time `json:"event_time" binding: "required"`
}

func (server *Server) addEvent(ctx *gin.Context) {
	var req AddEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req.EventTime)

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	arg := db.CreateEventParams{
		Username:  authPayload.Username,
		Title:     req.Title,
		EventTime: req.EventTime,
	}

	fmt.Println(arg.EventTime)

	res, err := server.store.CreateEvent(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

type getEventRequest struct {
	Limit  int32 `form:"limit" binding:"required,max=10"`
	Offset int32 `form:"offset" binding:"required,min=1"`
}

func (server *Server) getEvents(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	arg := db.GetAllEventsFromUserParams{
		Username: authPayload.Username,
		Offset:   req.Offset,
		Limit:    req.Limit,
	}
	events, err := server.store.GetAllEventsFromUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, events)
}

type deleteEventRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (server *Server) deleteEvent(ctx *gin.Context) {
	var req deleteEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	user, err := server.store.GetUser(ctx, authPayload.Username)
	fmt.Println(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Print(req.ID, user.ID)

	arg := db.GetSpecificEventFromUserParams{
		ID:     req.ID,
		IDUser: user.ID,
	}

	event, err := server.store.GetSpecificEventFromUser(ctx, arg)
	fmt.Println(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteEvent(ctx, event.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"status": true})

}

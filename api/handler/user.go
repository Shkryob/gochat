package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/model"
	"github.com/shkryob/gochat/utils"
)

func (handler *Handler) SignUp(context echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(context, &u); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := handler.userStore.Create(&u); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return utils.ResponseByContentType(context, http.StatusCreated, newUserResponse(&u))
}

func (handler *Handler) Login(context echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(context); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := handler.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return utils.ResponseByContentType(context, http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return utils.ResponseByContentType(context, http.StatusForbidden, utils.AccessForbidden())
	}
	return utils.ResponseByContentType(context, http.StatusOK, newUserResponse(u))
}

func userIDFromToken(context echo.Context) uint {
	id, ok := context.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}

func (handler *Handler) CurrentUser(context echo.Context) error {
	u, err := handler.userStore.GetByID(userIDFromToken(context))
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	return utils.ResponseByContentType(context, http.StatusOK, newUserResponse(u))
}

func (handler *Handler) GetUsers(context echo.Context) error {
	var (
		users []model.User
		count int
	)

	offset, err := strconv.Atoi(context.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(context.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	search := context.QueryParam("search")

	users, count, err = handler.userStore.List(offset, limit, search)

	return utils.ResponseByContentType(context, http.StatusOK, newUserListResponse(users, count))
}

func (handler *Handler) GetUser(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("user_id"), 10, 32)
	id := uint(id64)
	user, err := handler.userStore.GetByID(id)

	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if user == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	return utils.ResponseByContentType(context, http.StatusOK, newSimplifiedUserResponse(user))
}

func (handler *Handler) UploadAvatar(context echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := context.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	userID := userIDFromToken(context)
	// Destination
	dst, err := os.Create("uploads/avatars/" + fmt.Sprint(userID))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return utils.ResponseByContentType(context,
		http.StatusOK,
		fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func (handler *Handler) GetAvatar(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("user_id"), 10, 32)
	if err != nil {
		return err
	}
	id := uint(id64)

	return context.Inline("uploads/avatars/" + fmt.Sprint(id), "avatar_" + fmt.Sprint(id))
}

package handler

import (
	"CareerBoost/src/entity"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) addNewCourse(ctx *gin.Context) {
	var courseBody entity.CourseAdd
	if err := h.BindBody(ctx, &courseBody); err != nil {
		fmt.Println(err)
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	var courseDB entity.Course
	courseDB.Judul = courseBody.Judul
	courseDB.Deskripsi = courseBody.Deskripsi
	courseDB.Intro = courseBody.Intro
	courseDB.Rate = courseBody.Rate
	courseDB.Price = courseBody.Price

	if err := h.db.Create(&courseDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "failed to create course", nil)
		return
	}

	for _, playlist := range courseBody.Playlist {
		var playl entity.Playlist

		playl.Nama = playlist.Nama
		playl.CourseID = courseDB.ID

		var videos []entity.Video
		count := 0

		for _, video := range playlist.Video {

			durasi, err := time.ParseDuration(video.Durasi)
			if err != nil {
				h.ErrorResponse(ctx, http.StatusBadRequest, "invalid video duration", nil)
				return
			}

			count += int(durasi)

			videos = append(videos, entity.Video{
				Link:       video.Link,
				Judul:      video.Judul,
				Durasi:     video.Durasi,
				PlaylistID: playl.ID,
			})
		}

		playl.Durasi = time.Duration(count)
		playl.Video = videos

		if err := h.db.Create(&playl).Error; err != nil {
			h.ErrorResponse(ctx, http.StatusInternalServerError, "failed to add playlist", nil)
			return
		}
	}

	h.SuccessResponse(ctx, http.StatusOK, "Course berhasil ditambahkan", nil, nil)
}

func (h *handler) getAllCourse(ctx *gin.Context) {
	var courseParam entity.CourseParam
	if err := h.BindParam(ctx, &courseParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	courseParam.FormatPagination()

	var courseBody []entity.Course

	if err := h.db.
		Model(entity.Course{}).
		Limit(int(courseParam.Limit)).
		Offset(int(courseParam.Offset)).
		Find(&courseBody).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var totalElements int64

	if err := h.db.
		Model(entity.Course{}).
		Limit(int(courseParam.Limit)).
		Offset(int(courseParam.Offset)).
		Count(&totalElements).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	courseParam.ProcessPagination(totalElements)

	type resp struct {
		Foto      string  `json:"foto"`
		Judul     string  `json:"judul"`
		Deskripsi string  `json:"deskripsi"`
		Rate      float32 `json:"rate"`
		Price     float32 `json:"price"`
	}

	var courses []resp
	for _, course := range courseBody {

		var resps resp
		resps.Foto = course.Foto
		resps.Judul = course.Judul
		resps.Deskripsi = course.Deskripsi
		resps.Rate = course.Rate
		resps.Price = course.Price

		courses = append(courses, resps)
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", courses, &courseParam.PaginationParam)
}

func (h *handler) getCourseRecomendation(ctx *gin.Context) {
	var courseBody []entity.Course

	if err := h.db.Model(entity.Course{}).Order("rate desc").Limit(8).Find(&courseBody).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	type resp struct {
		Foto      string  `json:"foto"`
		Judul     string  `json:"judul"`
		Deskripsi string  `json:"deskripsi"`
		Rate      float32 `json:"rate"`
		Price     float32 `json:"price"`
	}

	var courses []resp
	for _, course := range courseBody {

		var resps resp
		resps.Foto = course.Foto
		resps.Judul = course.Judul
		resps.Deskripsi = course.Deskripsi
		resps.Rate = course.Rate
		resps.Price = course.Price

		courses = append(courses, resps)
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", courses, nil)
}

func (h *handler) getACourseHome(ctx *gin.Context) {
	var courseBody []entity.Course

	if err := h.db.
		Model(entity.Course{}).Limit(8).Find(&courseBody).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	type resp struct {
		Foto      string  `json:"foto"`
		Judul     string  `json:"judul"`
		Deskripsi string  `json:"deskripsi"`
		Rate      float32 `json:"rate"`
		Price     float32 `json:"price"`
	}

	var courses []resp
	for _, course := range courseBody {

		var resps resp
		resps.Foto = course.Foto
		resps.Judul = course.Judul
		resps.Deskripsi = course.Deskripsi
		resps.Rate = course.Rate
		resps.Price = course.Price

		courses = append(courses, resps)
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", courses, nil)
}

func (h *handler) getAllCourseRecomendation(ctx *gin.Context) {
	var courseParam entity.CourseParam
	if err := h.BindParam(ctx, &courseParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	courseParam.FormatPagination()

	var courseBody []entity.Course

	if err := h.db.
		Model(entity.Course{}).
		Order("rate desc").
		Limit(int(courseParam.Limit)).
		Offset(int(courseParam.Offset)).
		Find(&courseBody).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var totalElements int64

	if err := h.db.
		Model(entity.Course{}).
		Limit(int(courseParam.Limit)).
		Offset(int(courseParam.Offset)).
		Count(&totalElements).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	courseParam.ProcessPagination(totalElements)

	type resp struct {
		Foto      string  `json:"foto"`
		Judul     string  `json:"judul"`
		Deskripsi string  `json:"deskripsi"`
		Rate      float32 `json:"rate"`
		Price     float32 `json:"price"`
	}

	var courses []resp
	for _, course := range courseBody {

		var resps resp
		resps.Foto = course.Foto
		resps.Judul = course.Judul
		resps.Deskripsi = course.Deskripsi
		resps.Rate = course.Rate
		resps.Price = course.Price

		courses = append(courses, resps)
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", courses, &courseParam.PaginationParam)
}

func (h *handler) getCourseData(ctx *gin.Context) {
	var courseBody entity.CourseReqByID
	if err := h.BindBody(ctx, &courseBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to bind body", nil)
		return
	}

	var courseDB entity.Course

	err := h.db.Where("id = ?", courseBody.ID).Take(&courseDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var resp entity.CourseRespData

	resp.Judul = courseDB.Judul
	resp.Deskripsi = courseDB.Deskripsi
	resp.Intro = courseDB.Intro
	resp.Price = courseDB.Price
	resp.Rate = courseDB.Rate

	var playlists []entity.Playlist
	if err := h.db.Where("course_id = ?", courseDB.ID).Find(&playlists).Error; err != nil {
		fmt.Println(err)
		h.ErrorResponse(ctx, http.StatusInternalServerError, "error occurred", nil)
		return
	}

	for _, playlist := range playlists {
		var videos []entity.Video
		if err := h.db.Where("playlist_id = ?", playlist.ID).Find(&videos).Error; err != nil {
			fmt.Println(err)
			h.ErrorResponse(ctx, http.StatusInternalServerError, "error occurred", nil)
			return
		}

		var respVideos []entity.RespVideo
		for _, v := range videos {
			respVideos = append(respVideos, entity.RespVideo{
				Link:       v.Link,
				Judul:      v.Judul,
				Durasi:     v.Durasi,
				PlaylistID: v.PlaylistID,
			})
		}

		resp.Playlist = append(resp.Playlist, entity.RespPlaylist{
			Nama:     playlist.Nama,
			Durasi:   playlist.Durasi,
			CourseID: playlist.CourseID,
			Video:    respVideos,
		})
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", resp, nil)
}

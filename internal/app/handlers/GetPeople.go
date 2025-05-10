package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"taskEffectiveMobile/internal/app/models"
	"taskEffectiveMobile/internal/app/utils"
)

// GetPeople godoc
// @Summary Получить список пользователей
// @Description Получить пользователей с фильтрацией и пагинацией
// @Tags people
// @Accept json
// @Produce json
// @Param gender query string false "Пол"
// @Param nationality query string false "Национальность"
// @Param age query int false "Возраст"
// @Param page query int false "Страница"
// @Param limit query int false "Лимит"
// @Success 200 {array} models.Person
// @Failure 500 {object} utils.ErrorResponse
// @Router /person [get]
func (h *Handler) HandleGetPeople(w http.ResponseWriter, r *http.Request) error {
	utils.Logger.Info("GetPeople: запрос на получение списка людей")

	query := r.URL.Query()

	filter := models.PeopleFilter{}

	if gender := query.Get("gender"); gender != "" {
		filter.Gender = &gender
	}

	if nationality := query.Get("nationality"); nationality != "" {
		filter.Nationality = &nationality
	}

	if ageStr := query.Get("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err == nil {
			filter.Age = &age
		}
	}

	if pageStr := query.Get("page"); pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err == nil && page > 0 {
			filter.Page = page
		}
	} else {
		filter.Page = 1
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 {
			filter.Limit = limit
		}
	} else {
		filter.Limit = 10
	}
	fmt.Println(filter)

	utils.Logger.Debugf("применяемые фильтры: %+v", filter)

	people, err := h.DI.PersonUseCase.GetPeople(r.Context(), filter)
	if err != nil {
		utils.Logger.Errorf(" ошибка при получении списка: %v", err)
		return err
	}
	utils.Logger.Info("GetPeople: получено %d записей", len(people))
	return utils.WriteSuccessfulJSON(w, people)
}

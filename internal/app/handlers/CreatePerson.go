package handlers

import (
	"encoding/json"
	"net/http"
	"taskEffectiveMobile/internal/app/utils"
)

type CreatePersonRequest struct {
	FullName string `json:"full_name"`
}

// HandleCreatePerson godoc
// @Summary Добавить человека
// @Description Обогащает человека по имени и сохраняет в БД
// @Tags people
// @Accept json
// @Produce json
// @Param person body handlers.CreatePersonRequest true "Данные человека"
// @Success 200 {object} models.Person
// @Failure 400 {object} utils.ErrorResponse
// @Router /person [post]
func (h *Handler) HandleCreatePerson(w http.ResponseWriter, r *http.Request) error {
	utils.Logger.Info("CreatePerson: запрос на создание человека получен")

	req := CreatePersonRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Logger.Warnf("Ошибка декодирования тела запроса: %v", err)
		return utils.BadRequest("invalid JSON")
	}
	if req.FullName == "" {
		return utils.BadRequest("full name is required")
	}

	utils.Logger.Debugf("Добавление человека с данными: %+v", req)

	err := h.DI.PersonUseCase.Create(r.Context(), req.FullName)
	if err != nil {
		utils.Logger.Errorf("Ошибка при добавлении пользователя: %v", err)
		return err
	}
	utils.Logger.Infof("Пользователь добавлен: %s %s", req.FullName)
	w.WriteHeader(http.StatusCreated)
	return nil
}

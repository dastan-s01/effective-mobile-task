package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"taskEffectiveMobile/internal/app/models"
	"taskEffectiveMobile/internal/app/utils"
)

// UpdatePerson godoc
// @Summary Обновить данные пользователя
// @Description Обновляет поля пользователя по ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param person body models.Person true "Обновлённые данные пользователя"
// @Success 200 {object} models.Person
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /person/{id} [put]
func (h *Handler) HandleUpdatePerson(w http.ResponseWriter, r *http.Request) error {
	utils.Logger.Info("UpdatePerson: запрос на обновление человека")

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.BadRequest("invalid UUID")
	}
	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		utils.Logger.Warnf("Ошибка декодирования тела запроса: %v", err)
		return utils.BadRequest("invalid JSON")
	}
	person.ID = id
	utils.Logger.Debugf("обновляемые данные: %+v", person)

	err = h.DI.PersonUseCase.Update(r.Context(), &person)
	if err != nil {
		utils.Logger.Errorf("ошибка обновления данных пользователя: %v", err)
		return err
	}
	utils.Logger.Info(" обновление завершено успешно")
	return utils.WriteSuccessfulJSON(w, map[string]string{"status": "updated"})
}

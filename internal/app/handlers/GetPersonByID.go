package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"taskEffectiveMobile/internal/app/utils"
)

// GetPersonByID godoc
// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по его UUID
// @Tags people
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.Person
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /person/{id} [get]
func (h *Handler) HandleGetPersonByID(w http.ResponseWriter, r *http.Request) error {
	utils.Logger.Info("GetPersonByID: запрос на получение по ID")

	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		return utils.BadRequest("invalid UUID")
	}

	person, err := h.DI.PersonUseCase.GetByID(r.Context(), id)
	if err != nil {
		return utils.NotFound("person not found")
	}

	utils.Logger.Debugf("GetPersonByID: полученные данные: %+v", person)
	return utils.WriteSuccessfulJSON(w, person)
}

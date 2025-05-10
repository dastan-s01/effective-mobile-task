package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"taskEffectiveMobile/internal/app/utils"
)

// DeletePerson godoc
// @Summary Удалить пользователя
// @Description Удаляет пользователя по ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 204 {string} string "no content"
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /person/{id} [delete]
func (h *Handler) HandleDeletePerson(w http.ResponseWriter, r *http.Request) error {
	utils.Logger.Info("DeletePerson: запрос на удаление человека")

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.BadRequest("invalid UUID")
	}
	utils.Logger.Debugf("Удаление человека с id: %s", id)

	err = h.DI.PersonUseCase.Delete(r.Context(), id)
	if err != nil {
		utils.Logger.Errorf("Ошибка при удалении: %v", err)

		return utils.NotFound("person not found")
	}
	utils.Logger.Infof("Пользователь удален: %s ", id)
	w.WriteHeader(http.StatusNoContent)
	return nil
}

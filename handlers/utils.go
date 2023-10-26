package handlers

import (
	"context"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (*string, error) {
	token, err := middlewares.GetTokenInGin(c)
	if err != nil {
		return nil, err
	}
	return &token.RegisteredClaims.Subject, nil
}

func (h *Handler) GetUserById(
	ctx context.Context, userId string,
) (*ent.User, error) {
	userNode, err := h.DB.User.Query().
		Where(user.ID(userId)).
		WithQuestionnaireResponses(func(qrq *ent.QuestionnaireResponseQuery) {
			qrq.WithQuestionnaire(func(qq *ent.QuestionnaireQuery) {
				qq.WithQuestions()
			}).WithAnswers(func(aq *ent.AnswerQuery) {
				aq.WithChosen()
			})
		}).
		WithMycards().
		WithPrices().
		WithDeegoo().
		WithHealthkit(func(hkq *ent.HealthKitQuery) {
			hkq.WithData()
		}).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return userNode, nil
}
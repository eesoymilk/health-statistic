// Code generated by ent, DO NOT EDIT.

package userquestionnaire

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestionnaire applies the HasEdge predicate on the "questionnaire" edge.
func HasQuestionnaire() predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionnaireTable, QuestionnaireColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionnaireWith applies the HasEdge predicate on the "questionnaire" edge with a given conditions (other predicates).
func HasQuestionnaireWith(preds ...predicate.Questionnaire) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		step := newQuestionnaireStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnswers applies the HasEdge predicate on the "answers" edge.
func HasAnswers() predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnswersTable, AnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnswersWith applies the HasEdge predicate on the "answers" edge with a given conditions (other predicates).
func HasAnswersWith(preds ...predicate.Answer) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		step := newAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserQuestionnaire) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserQuestionnaire) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserQuestionnaire) predicate.UserQuestionnaire {
	return predicate.UserQuestionnaire(func(s *sql.Selector) {
		p(s.Not())
	})
}
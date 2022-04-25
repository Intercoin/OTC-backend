package postgres

import (
	"fmt"
	"strings"
	"time"
)

type Condition interface {
	Key() string
	Value() interface{}
	Query() string
}

func PrepareConditions(conditions []Condition) (string, map[string]interface{}) {
	queries := make([]string, 0, len(conditions))
	mapper := make(map[string]interface{})
	for _, condition := range conditions {
		mapper[condition.Key()] = condition.Value()
		queries = append(queries, condition.Query())
	}

	return strings.Join(queries, " AND "), mapper
}

type strictCondition struct {
	key   string
	value interface{}
}

func NewStrictCondition(key string, value interface{}) Condition {
	return &strictCondition{
		key:   key,
		value: value,
	}
}

func (s *strictCondition) Key() string {
	return s.key
}

func (s *strictCondition) Value() interface{} {
	return s.value
}

func (s *strictCondition) Query() string {
	return fmt.Sprintf("%s = :%s", s.key, s.key)
}

type strictNotCondition struct {
	key   string
	value interface{}
}

func NewStrictNotCondition(key string, value interface{}) Condition {
	return &strictNotCondition{
		key:   key,
		value: value,
	}
}

func (s *strictNotCondition) Key() string {
	return s.key
}

func (s *strictNotCondition) Value() interface{} {
	return s.value
}

func (s *strictNotCondition) Query() string {
	return fmt.Sprintf("%s != :%s", s.key, s.key)
}

type strictMoreCondition struct {
	key   string
	value interface{}
	r     int64
}

func NewStrictMoreCondition(key string, value interface{}) Condition {
	return &strictMoreCondition{
		key:   key,
		value: value,
		r:     time.Now().UnixNano(),
	}
}

func (s *strictMoreCondition) Key() string {
	return fmt.Sprintf("%s%d", s.key, s.r)
}

func (s *strictMoreCondition) Value() interface{} {
	return s.value
}

func (s *strictMoreCondition) Query() string {
	return fmt.Sprintf("%s >= :%s", s.key, s.Key())
}

type strictLessCondition struct {
	key   string
	value interface{}
	r     int64
}

func NewStrictLessCondition(key string, value interface{}) Condition {
	return &strictLessCondition{
		key:   key,
		value: value,
		r:     time.Now().UnixNano(),
	}
}

func (s *strictLessCondition) Key() string {
	return fmt.Sprintf("%s%d", s.key, s.r)
}

func (s *strictLessCondition) Value() interface{} {
	return s.value
}

func (s *strictLessCondition) Query() string {
	return fmt.Sprintf("%s <= :%s", s.key, s.Key())
}

type arrayContainsCondition struct {
	key   string
	value interface{}
}

func NewArrayContainsCondition(key string, value interface{}) Condition {
	return &arrayContainsCondition{
		key:   key,
		value: value,
	}
}

func (a *arrayContainsCondition) Key() string {
	return a.key
}

func (a *arrayContainsCondition) Value() interface{} {
	return a.value
}

func (a *arrayContainsCondition) Query() string {
	return fmt.Sprintf("%s @> :%s", a.key, a.key)
}

type jsonbContainsCondition struct {
	key   string
	value interface{}
}

func NewJSONBContainsCondition(key string, value interface{}) Condition {
	return &jsonbContainsCondition{
		key:   key,
		value: value,
	}
}

func (a *jsonbContainsCondition) Key() string {
	return a.key
}

func (a *jsonbContainsCondition) Value() interface{} {
	return a.value
}

func (a *jsonbContainsCondition) Query() string {
	return fmt.Sprintf("%s @> :%s", a.key, a.key)
}

type inCondition struct {
	key   string
	value interface{}
}

func NewInCondition(key string, values interface{}) Condition {
	return &inCondition{
		key:   key,
		value: values,
	}
}

func (i *inCondition) Key() string {
	return i.key
}

func (i *inCondition) Value() interface{} {
	return i.value
}

func (i *inCondition) Query() string {
	return fmt.Sprintf("%s IN (:%s)", i.key, i.key)
}

type inJSONBCondition struct {
	key   string
	value interface{}
}

func NewInJSONBCondition(key string, values interface{}) Condition {
	return &inJSONBCondition{
		key:   key,
		value: values,
	}
}

func (i *inJSONBCondition) Key() string {
	return i.key
}

func (i *inJSONBCondition) Value() interface{} {
	return i.value
}

func (i *inJSONBCondition) Query() string {
	return fmt.Sprintf("index IN (Select index from huntress, jsonb_to_recordset(attributes) as x(value varchar) where value = :%s)", i.key)
}

type isNotNullCondition struct {
	key   string
	value []interface{}
}

func NewIsNotNullCondition(key string, values []interface{}) Condition {
	return &isNotNullCondition{
		key:   key,
		value: values,
	}
}

func (i *isNotNullCondition) Key() string {
	return i.key
}

func (i *isNotNullCondition) Value() interface{} {
	return nil
}

func (i *isNotNullCondition) Query() string {
	return fmt.Sprintf("%s IS NOT NULL", i.key)
}

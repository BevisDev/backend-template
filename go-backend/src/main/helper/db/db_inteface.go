package db

type IDb interface {
	GetList(dest interface{}, state, query string, args map[string]interface{}) bool
	Get(dest interface{}, state, query string, args map[string]interface{}) bool
	Insert(state, query string, args interface{}) bool
	Update(state, query string, args interface{}) bool
	Delete(state, query string, args interface{}) bool
}

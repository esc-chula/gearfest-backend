package interfaces

type SqlHandler interface {
	Create(interface{}) error
	UpdateColumn(interface{}, string, interface{}) error
	GetByPrimaryKey(interface{}) error
	GetWithAssociations(interface{}) error
}
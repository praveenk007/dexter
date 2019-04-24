package interfaces

type IConfigService interface {
	GetConfig(id string, ctype string) *[]byte
}

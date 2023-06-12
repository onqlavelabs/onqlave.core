package tools

type ServiceEncryptor struct{}

func NewServiceEncryptor() *ServiceEncryptor {
	return &ServiceEncryptor{}
}

func (cv *ServiceEncryptor) MaskUserName(fullName string) (maskedName string, err error) {
	maskedName = fullName[:len(fullName)/2]
	for i := len(fullName) / 2; i < len(fullName); i++ {
		maskedName += "*"
	}
	return maskedName, nil
}

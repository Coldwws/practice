package mocks



type AuthServiceMock struct{
	LoginFunc func(username,password string)(string,error)
	RegisterFunc func(username,password string)error
}

func (m *AuthServiceMock)Login(username,password string)(string,error){
	return m.LoginFunc(username,password)
}

func(m *AuthServiceMock)Register(username,password string) error{

	if m.RegisterFunc != nil{
		return m.RegisterFunc(username,password)
	}
	return nil
}

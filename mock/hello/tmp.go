package hello


type Talker interface{
	SayHello(s string)string
}

type PerSon struct {
	Name string
}
func (p *PerSon)SayHello(s string)string{

	return "Hello"+s
}

func NewPerson()*PerSon{
	return &PerSon{"person"}
}

type ComP struct {
	 User Talker
}

func NreComP(talker Talker)*ComP{
	return &ComP{User:talker}
}

func (c *ComP)Metting (s string)string{
	return c.User.SayHello(s)
}


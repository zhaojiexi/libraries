package goconver


type Person struct{
	Name string
}


func (p *Person)NewPerson(name string)*Person{
	return &Person{Name: name}
}

func (p *Person)GetName()string{
	return p.Name
}


package onetomany

type OneToMany struct {
	builder map[uint]interface{}
	order   []uint
}

func NewOneToMany() *OneToMany {
	return &OneToMany{
		order:   []uint{},
		builder: make(map[uint]interface{}),
	}
}

func (o *OneToMany) Create(key uint, builder interface{}) {
	o.builder[key] = builder
	o.order = append(o.order, key)
}

func (o *OneToMany) ForEach(fn func(interface{})) {
	for i := 0; i < len(o.order); i++ {
		fn(o.builder[o.order[i]])
	}
}

func (o *OneToMany) Get(key uint) interface{} {
	builder := o.builder[key]
	return builder
}

func (o *OneToMany) Exist(key uint) bool {
	_, ok := o.builder[key]
	return ok
}

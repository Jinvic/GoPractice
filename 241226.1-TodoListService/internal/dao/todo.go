package dao

type todoDao struct{}

var TodoDao = &todoDao{}

func (d *todoDao) GetList() {}
func (d *todoDao) Create()  {}
func (d *todoDao) Get()     {}
func (d *todoDao) Update()  {}
func (d *todoDao) Delete()  {}

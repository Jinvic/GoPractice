package service

type todoService struct{}

var TodoService = &todoService{}

func (s *todoService) GetList() {}
func (s *todoService) Create()  {}
func (s *todoService) Get()     {}
func (s *todoService) Update()  {}
func (s *todoService) Delete()  {}

package article

import "blog-service/pkg/define"

func Create(req *define.ArticleCreateReq) (*define.ArticleInfo, error) {

}

func View(id uint) (*define.ArticleInfo, error) {

}

func Edit(id uint, req *define.ArticleEditReq) (*define.ArticleInfo, error) {

}

func Delete(id uint) error {

}

func List() (*[]define.ArticleInfo, error) {

}

func ListAll() (*[]define.ArticleInfo, error) {

}

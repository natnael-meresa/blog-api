package article

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"twof/blog-api/internal/constant/model"
)

func (s *service) GetArticles(article *[]model.Article) (err error) {
	err = s.articlePersist.GetAllArticles(article)
	if err != nil {
		return fmt.Errorf("failed to save new user %s", err.Error())
	}

	return nil
}

func (s *service) CreateArticle(article *model.Article) error {

	if err := s.articleRepo.ValidateArticle(article); err != nil {
		return err
	}

	err := s.articlePersist.CreateArticle(article)
	if err != nil {
		return fmt.Errorf("failed to save new Article %s", err.Error())
	}

	return nil
}

func (s *service) GetAllArticles(article *[]model.Article) (err error) {
	err = s.articlePersist.GetAllArticles(article)
	if err != nil {
		return fmt.Errorf("failed to Get Articles %s", err.Error())
	}

	return nil

}

func (s *service) GetArticleById(articleId uint, article *model.Article) (err error) {
	err = s.articlePersist.GetArticleById(articleId, article)

	if err != nil {
		return fmt.Errorf("failed to Get user %s", err.Error())
	}

	return nil

}

func (s *service) UploadImages(articleId int, files []*multipart.FileHeader) error {
	var images []string

	for _, fileHeader := range files {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}

		name := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

		images = append(images, name)
		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {

			return err
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		err = os.MkdirAll("./public", os.ModePerm)
		if err != nil {

			return err
		}

		f, err := os.Create("./public/" + name)
		if err != nil {

			return err
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
	}

	err := s.articlePersist.UpdateImage(articleId, images)

	if err != nil {
		return err
	}

	return nil
}

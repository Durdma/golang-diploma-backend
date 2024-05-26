package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/cache"
	"sas/pkg/logger"
	"time"
)

// UniversitiesService - Структура для работы с сервисом университетов
type UniversitiesService struct {
	repo  repository.Universities
	cache cache.Cache
}

// NewUniversitiesService - Создание структуры сервиса университетов
func NewUniversitiesService(repo repository.Universities, cache cache.Cache) *UniversitiesService {
	return &UniversitiesService{
		repo:  repo,
		cache: cache,
	}
}

func (s *UniversitiesService) AddUniversity(ctx context.Context, domainId primitive.ObjectID, domainName string, shortName string) (primitive.ObjectID, error) {
	id, err := s.repo.Create(ctx, models.University{
		DomainId:  domainId,
		Name:      domainName,
		ShortName: shortName,
		Settings: models.Settings{
			MainColor:                "red", //TODO refactor for valid opts
			MainColorHover:           "green",
			MainFooterFontColor:      "blue",
			MainFooterFontColorHover: "white",
			MainFooterBgColor:        "blue",
		},
	})

	err = s.createCss(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, err
}

func (s *UniversitiesService) createCss(universityId primitive.ObjectID) error {
	path := "../../static/css/"
	f, err := os.OpenFile(path+universityId.Hex()+"_"+"css.css", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("body {\n    margin: 0;\n    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',\n    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',\n    sans-serif, 'Josefin Sans';\n    -webkit-font-smoothing: antialiased;\n    -moz-osx-font-smoothing: grayscale;\n}\n\n.navigate {\n    font-family: 'Josefin Sans', sans-serif;\n    font-weight: 700;\n    font-size: 20px;\n    color: #1a890e;\n}\n\n.navigate-items {\n    background-color: #1a890e;\n}\n\n.news-link {\n    font-family: 'Josefin Sans', sans-serif;\n    font-weight: 700;\n    color: #1a890e;\n}\n\n.nav-link-custom:hover {\n    color: #0c2807;\n}\n\n.main {\n    min-height: calc(100vh - 200px);\n}\n\n.footer {\n    width: 100%;\n    background-color: #1a890e;\n    color: #faebd7;\n    padding: 2rem;\n    margin-top: 2rem;\n}\n\n.footer-navbar {\n    color: #faebd7;\n}\n\n.footer-navbar:hover {\n    color: #0c2807;\n}\n\ncode {\n    font-family: source-code-pro, Menlo, Monaco, Consolas, 'Courier New',\n    monospace;\n}\n\n.form-signin {\n    max-width: 330px;\n    padding: 1rem;\n}\n\n.form-signin .form-floating:focus-within {\n    z-index: 2;\n}\n\n.form-signin input[type=\"email\"] {\n    margin-bottom: -1px;\n    border-bottom-right-radius: 0;\n    border-bottom-left-radius: 0;\n}\n\n.form-signin input[type=\"password\"] {\n    margin-bottom: 10px;\n    border-top-left-radius: 0;\n    border-top-right-radius: 0;\n}")

	return err
}

// GetByDomain - Получение из БД записи об университете по полученному домену
func (s *UniversitiesService) GetByDomain(ctx context.Context, domainName string) (models.University, error) {
	if value, err := s.cache.Get(domainName); err == nil {
		return value.(models.University), nil
	}

	logger.Info(domainName)

	univ, err := s.repo.GetByDomain(ctx, domainName)
	if err != nil {
		return models.University{}, err
	}

	s.cache.Set(domainName, univ)

	return univ, nil
}

func (s *UniversitiesService) GetByUniversityId(ctx context.Context, universityId primitive.ObjectID) (models.University, error) {
	return s.repo.GetByUniversityId(ctx, universityId)
}

func (s *UniversitiesService) GetUniversityColors(ctx context.Context, universityId primitive.ObjectID) (map[string]string, error) {
	university, err := s.repo.GetByUniversityId(ctx, universityId)
	if err != nil {
		return nil, err
	}

	fmt.Println(")()()()()()()()()()()()()()())))))))))))))))))))")
	fmt.Println(university.Settings)

	colors := make(map[string]string)

	colors["main_color"] = university.Settings.MainColor
	colors["main_color_hover"] = university.Settings.MainColorHover
	colors["main_footer_font_color"] = university.Settings.MainFooterFontColor
	colors["main_footer_font_color_hover"] = university.Settings.MainFooterFontColorHover
	colors["main_footer_bg_color"] = university.Settings.MainFooterBgColor

	fmt.Println(colors)
	fmt.Println(")()()()()()()()()()()()()()())))))))))))))))))))")
	return colors, err
}

func (s *UniversitiesService) PatchUniversityCSS(ctx context.Context, universityId primitive.ObjectID, colors map[string]string) error {
	err := s.patchCSS(universityId, colors)
	if err != nil {
		return err
	}

	return s.repo.ChangeCSS(ctx, universityId, colors)
}

func (s *UniversitiesService) patchCSS(universityId primitive.ObjectID, colors map[string]string) error {
	path := "../../static/css/"

	f, err := os.OpenFile(path+universityId.Hex()+"_"+"css.css", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("body {\n    margin: 0;\n    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',\n    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',\n    sans-serif, 'Josefin Sans';\n    -webkit-font-smoothing: antialiased;\n    -moz-osx-font-smoothing: grayscale;\n}\n\n.navigate {\n    font-family: 'Josefin Sans', sans-serif;\n    font-weight: 700;\n    font-size: 20px;\n    color: %s;\n}\n\n.navigate-items {\n    background-color: %s;\n}\n\n.news-link {\n    font-family: 'Josefin Sans', sans-serif;\n    font-weight: 700;\n    color: %s;\n}\n\n.nav-link-custom:hover {\n    color: %s;\n}\n\n.main {\n    min-height: calc(100vh - 200px);\n}\n\n.footer {\n    width: 100%%;\n    background-color: %s;\n    color: %s;\n    padding: 2rem;\n    margin-top: 2rem;\n}\n\n.footer-navbar {\n    color: %s;\n}\n\n.footer-navbar:hover {\n    color: %s;\n}\n\ncode {\n    font-family: source-code-pro, Menlo, Monaco, Consolas, 'Courier New',\n    monospace;\n}\n\n.form-signin {\n    max-width: 330px;\n    padding: 1rem;\n}\n\n.form-signin .form-floating:focus-within {\n    z-index: 2;\n}\n\n.form-signin input[type=\"email\"] {\n    margin-bottom: -1px;\n    border-bottom-right-radius: 0;\n    border-bottom-left-radius: 0;\n}\n\n.form-signin input[type=\"password\"] {\n    margin-bottom: 10px;\n    border-top-left-radius: 0;\n    border-top-right-radius: 0;\n}", colors["main_color"], colors["main_color"], colors["main_color"], colors["main_color_hover"], colors["main_footer_bg_color"], colors["main_footer_font_color"], colors["main_footer_font_color"], colors["main_footer_font_color_hover"]))

	return err
}

func (s *UniversitiesService) SetUniversityHistory(ctx context.Context, universityId primitive.ObjectID, history models.History) error {
	history.CreatedAt = time.Now()
	history.UpdatedAt = time.Now()
	history.UpdatedBy = history.CreatedBy

	return s.repo.SetUniversityHistory(ctx, universityId, history)
}

package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/lfcifuentes/go-repository-pattern/app/http/handler/responses"
	"github.com/lfcifuentes/go-repository-pattern/app/model"
	"github.com/lfcifuentes/go-repository-pattern/app/repository"

	"io"
	"net/http"
	"strconv"
)

var uni *ut.UniversalTranslator
var validate *validator.Validate
var trans ut.Translator

func init() {
	esTranslator := es.New()
	uni = ut.New(esTranslator, esTranslator)
	trans, _ = uni.GetTranslator("es")
	validate = validator.New()
	err := es_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

type UserController struct {
	UserRepository repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: repo,
	}
}

func (uc *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserRepository.GetAll()
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	response := responses.ResponseOk(users, "")
	render.JSON(w, r, response)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {

	var newUser = model.User{}
	reqBody, _ := io.ReadAll(r.Body)
	newUser, _ = model.UnmarshalUser(reqBody)

	err := validate.Struct(newUser)
	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)
		responses.ResponseError(w, r, errs.Translate(trans), "Upps, no hemos podido procesar tu solicitud")
		return
	}

	err = uc.UserRepository.Create(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := responses.ResponseOk(nil, "usuario creado correctamente!")
	render.JSON(w, r, response)
}

func (uc *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := uc.UserRepository.GetID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := responses.ResponseOk(user, "pong")

	render.JSON(w, r, response)
}

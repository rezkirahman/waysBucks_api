package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	productdto "waysbucks/dto/product"
	dto "waysbucks/dto/result"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/go-playground/validator/v10"
	//"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)


type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range products {
		products[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(product)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data token
	//userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	//userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	request := productdto.CreateProductRequest{
		Name:  r.FormValue("name"),
		Price: price,
		Qty:   qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product := models.Product{
		Name:   request.Name,
		Price:  request.Price,
		Image:  filename,
		Qty:    request.Qty,
		//UserID: userId,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(productdto.UpdateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// dataContex := r.Context().Value("dataFile") // add this code
	// filename := dataContex.(string) // add this code

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepository.GetProduct(int(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	// if request.Image !="" {
	// 	product.Image=filename
	// }

	if request.Qty != 0 {
		product.Qty = request.Qty
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProduct(u models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
		Image: u.Image,
		Qty:   u.Qty,
	}
}

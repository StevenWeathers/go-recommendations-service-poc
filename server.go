package main

import (
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
)

func main() {
  router := httprouter.New()
  router.GET("/product/:productId/related", related)
  router.GET("/product/:productId/required", required)
  router.GET("/product/:productId/customeralsoviewed", customerAlsoViewed)

  log.Fatal(http.ListenAndServe(":8080", router))
}

func related(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  productId := ps.ByName("productId")

  recommended, err := getProductRecommendation("related", productId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  recommendation := Related{Related: recommended}

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(recommendation)
}

func required(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  productId := ps.ByName("productId")

  recommended, err := getProductRecommendation("required", productId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  recommendation := Required{Required: recommended}

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(recommendation)
}

func customerAlsoViewed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  productId := ps.ByName("productId")

  recommended, err := getProductRecommendation("customeralsoviewed", productId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  recommendation := CustomerAlsoViewed{CustomerAlsoViewed: recommended}

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(recommendation)
}

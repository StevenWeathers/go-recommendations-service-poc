package main

import (
  "fmt"
  "strings"
  "github.com/garyburd/redigo/redis"
)

func getProductRecommendation(recommendationType string, productId string) ([]string, error) {
  c, err := redis.Dial("tcp", "db:6379")
  if err != nil {
    return nil, err
  }

  recommended, err := redis.String(c.Do("GET", fmt.Sprintf("product_%s_%s", recommendationType, productId)))
  splitRecommendation := strings.Split(recommended, ",")

  defer c.Close()

  return splitRecommendation, nil
}
package go_expert_rate_limiter

import "fmt"
import (
    "github.com/rs/zerolog/log"
)

func handleErr(err error) {
	if err != nil {
		log.Error().Err(err).Msg("")
		panic(err)
	}
}


func Bootstap() {
	//TODO: Your code here! :)
	
fmt.Println("GO-ALUSOFT!")
fmt.Println("You may want to create an entity 👇")
fmt.Println("go-alusoft create entity mySampleEntity")
fmt.Println("Or maybe an usecase 👇")
fmt.Println("go-alusoft create usecase mySampleUsecase")
//	workdir, err := os.Getwd()
//	handleErr(err)
//	appConfig, err := configs.LoadConfig(workdir)
//    if err != nil {
//        panic(err)
//    }
//    fmt.Printf("%v", appConfig.SampleConfigVar)

}
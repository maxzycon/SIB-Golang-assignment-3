package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/maxzycon/SIB-Golang-Assigment-3/pkg/constant"
	"github.com/maxzycon/SIB-Golang-Assigment-3/pkg/model"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

type LogAutoReload struct {
	Water uint `json:"water"`
	Wind  uint `json:"wind"`
}

func (s Service) AutoReload() (err error) {
	// ---- auto run 15 detik skali pakai goroutine
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			dataWind := model.AutoReload{}
			if err = s.db.First(&dataWind, "name = ?", "wind").Error; err != nil {
				return
			}

			dataWater := model.AutoReload{}
			if err = s.db.First(&dataWater, "name = ?", "water").Error; err != nil {
				return
			}

			// --- 15 detik update
			// ----- Random value 1 - 20
			min := 1
			max := 20
			randValWater := rand.Intn(max-min) + min
			randValWind := rand.Intn(max-min) + min

			log := LogAutoReload{
				Water: uint(randValWater),
				Wind:  uint(randValWind),
			}

			logString, err := json.Marshal(&log)
			if err != nil {
				return
			}

			fmt.Println(string(logString))

			newStatusWater := constant.AMAN
			newStatusWind := constant.AMAN

			// ---- update settings

			// jika water antara 6 - 8 maka status siaga
			if randValWater >= 6 && randValWater <= 8 {
				newStatusWater = constant.SIAGA
			}

			// jika water diatas 8 maka status bahaya
			if randValWater > 8 {
				newStatusWater = constant.BAHAYA
			}

			// jika wind antara 7 - 15 maka status siaga
			if randValWind >= 7 && randValWind <= 15 {
				newStatusWind = constant.SIAGA
			}

			// jika wind diatas 15 maka status bahaya
			if randValWind > 15 {
				newStatusWind = constant.BAHAYA
			}

			// --- Update status
			if newStatusWater == constant.AMAN {
				fmt.Println("status water : aman")
			}

			if newStatusWater == constant.SIAGA {
				fmt.Println("status water : siaga")
			}

			if newStatusWater == constant.BAHAYA {
				fmt.Println("status water : bahaya")
			}

			if newStatusWind == constant.AMAN {
				fmt.Println("status wind : aman")
			}

			if newStatusWind == constant.SIAGA {
				fmt.Println("status wind : siaga")
			}

			if newStatusWind == constant.BAHAYA {
				fmt.Println("status wind : bahaya")
			}

			if err = s.db.Updates(&model.AutoReload{
				Model: gorm.Model{
					ID: dataWind.ID,
				},
				Value:  randValWind,
				Status: newStatusWind,
			}).Error; err != nil {
				return
			}
			if err = s.db.Updates(&model.AutoReload{
				Model: gorm.Model{
					ID: dataWater.ID,
				},
				Value:  randValWater,
				Status: newStatusWater,
			}).Error; err != nil {
				return
			}

			fmt.Println("=========== Wait 15 second ============")
			time.Sleep(15 * time.Second) // Menunggu 15 detik
		}
	}()

	return
}

package controllers

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type HVList struct {
	Mutex sync.Mutex
	HVs   map[string]*HV
}

var Cloud *HVList

func InitCloud() *HVList {
	Cloud = new(HVList)
	if err := getHVs(Cloud); err != nil {
		log.Fatal().Err(err).Msg("Failed to get HVs")
	} else {
		count := len(Cloud.HVs)
		log.Info().Msgf("Found %d HVs", count)
	}
	return Cloud
}

package main

import (
    "time"
)

type SleepData struct {
    Nights []SleepNight      `json:"sleep"`
}

type SleepNight struct {
    Awake                       int         `json:"awake"`
    Bedtime_end                 time.Time   `json:"bedtime_end"`
    Bedtime_end_delta           int         `json:"bedtime_end_delta"`
    Bedtime_start               time.Time   `json:"bedtime_start"`
    Bedtime_start_delta         int         `json:"bedtime_start_delta"`
    Breath_average              int         `json:"breath_average"`
    Deep                        int         `json:"deep"`
    Duration                    int         `json:"duration"`
    Efficiency                  int         `json:"efficiency"`
    Hr_5min                     []int       `json:"hr_5min"`
    Hr_average                  float32     `json:"hr_average"`
    Hr_lowest                   int         `json:"hr_lowest"`
    Hypnogram_5min              string      `json:"hypnogram_5min"`
    Is_longest                  int         `json:"is_longest"`
    Light                       int         `json:"light"`
    Midpoint_at_delta           int         `json:"midpoint_at_delta"`
    Midpoint_time               int         `json:"midpoint_time"`
    Onset_latency               int         `json:"onset_latency"`
    Period_id                   int         `json:"period_id"`
    Rem                         int         `json:"rem"`
    Restless                    int         `json:"restless"`
    Rmssd                       int         `json:"rmssd"`
    Rmssd_5min                  []int       `json:"rmssd_5min"`
    Score                       int         `json:"score"`
    Score_alignment             int         `json:"score_alignment"`
    Score_deep                  int         `json:"score_deep"`
    Score_disturbances          int         `json:"score_disturbances"`
    Score_efficiency            int         `json:"score_efficiency"`
    Score_latency               int         `json:"score_latency"`
    Score_rem                   int         `json:"score_rem"`
    Score_total                 int         `json:"score_total"`
    Summary_date                string      `json:"summary_date"`
    Temperature_delta           float32     `json:"temperature_delta"`
    Temperature_deviation       float32     `json:"temperature_deviation"`
    Temperature_trend_deviation float32     `json:"temperature_trend_deviation"`
    Timezone                    int         `json:"timezone"`
    Total                       int         `json:"total"`
}

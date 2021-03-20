package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tianzuoan/cobra-demmo/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间工具",
	Long:  "用于对时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := time.Now()
		log.Printf("当前时间是：%s, 时间戳：%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime, duration string

var timeCalcCmd = &cobra.Command{
	Use:   "calc",
	Short: "时间计算",
	Long:  "根据传入的时间calculate（简写c）（不传则为当前时间），持续时间参数duration(简写d)",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"

		if calculateTime == "" {
			currentTime = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			currentTime, err = time.Parse(layout, calculateTime)
			if err != nil {
				calculateTimeStamp, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(calculateTimeStamp), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(currentTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime failed! err:%v", err)
		}
		log.Printf("计算出的结果是：%s, 时间戳：%d", calculateTime.Format("2006-01-02 15:04:05"), calculateTime.Unix())
	},
}

func init() {
	//添加子命令
	timeCmd.AddCommand(timeNowCmd)
	timeCmd.AddCommand(timeCalcCmd)
	timeCalcCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或格式化后的时间！")
	timeCalcCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位有"ns", "us", "ms", "s", "m", "h"`)
}

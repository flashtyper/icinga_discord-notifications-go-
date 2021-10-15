package main

import (
    "fmt"
    "os"
)


func main() {
    argsWithoutProg := os.Args[1:]
    icingaMap := make(map[string]string)
    for i := 0; i < len(argsWithoutProg); i +=2 {
        icingaMap[argsWithoutProg[i]] = argsWithoutProg[i+1]
    }

    if icingaMap["notification_type"] == "CUSTOM" {
        custom(icingaMap)
    } else {
        host_or_service(icingaMap)
    }
}

func custom (icingaMap map [string]string) {
        emote := get_emote(icingaMap["notification_type"])
        notification := fmt.Sprintf("%v %v: %v%v%v%v%v%v%v%v%v: %v%v%v%v", emote, icingaMap["notification_type"], "Service **", icingaMap["service_name"], "** on host **", icingaMap["host_name"], "** is in state **", icingaMap["service_state"], "**!\n", "```", icingaMap["notification_author"], icingaMap["notification_comment"], "\n\n", icingaMap["service_output"], "```")
        fmt.Println(notification)
}

func host_or_service (icingaMap map[string]string) {
        if icingaMap["service_name"] != "" {
                emote := get_emote(icingaMap["service_state"])
                notification := fmt.Sprintf("%v %v: %v%v%v%v%v%v%v%v%v%v", emote, icingaMap["notification_type"], "Service **", icingaMap["service_name"], "** on host **", icingaMap["host_name"], "** changed to **", icingaMap["service_state"], "**!\n", "```", icingaMap["service_output"], "```")
        fmt.Println(notification)
        }
}


func get_emote (service_state string) (emote string) {

        emojiMap := map[string]string {
                "OK": ":white_check_mark:",
                "RECOVERY": ":white_check_mark:",
                "WARNING": ":warning:",
                "DOWN": ":rotating_light:",
                "UP": ":white_check_mark:",
                "PROBLEM": ":rotating_light:",
                "CRITICAL": ":rotating_light:",
                "UNKNOWN": ":grey_question:",
                "DOWNTIMESTART": ":construction:",
                "DOWNTIMEEND": ":construction:",
                "CUSTOM": ":paintbrush:",
        }
        emote = emojiMap[service_state]
        return
}



package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"
    "bytes"
)


func main() {
    argsWithoutProg := os.Args[1:]
    icingaMap := make(map[string]string)
    for i := 0; i < len(argsWithoutProg); i +=2 {
        icingaMap[argsWithoutProg[i]] = argsWithoutProg[i+1]
    }

    if icingaMap["notification_type"] == "CUSTOM" {
            message := custom(icingaMap)
            send_msg(icingaMap["webhook_username"], icingaMap["webhook_url"], message)
    } else {
            message := host_or_service(icingaMap)
            send_msg(icingaMap["webhook_username"], icingaMap["webhook_url"], message)
    }
}


func custom (icingaMap map [string]string) (message string){
        emote := get_emote(icingaMap["notification_type"])
        message = fmt.Sprintf("%v %v: %v%v%v%v%v%v%v%v%v: %v%v%v%v", emote, icingaMap["notification_type"], "Service **", icingaMap["service_name"], "** on host **", icingaMap["host_name"], "** is in state **", icingaMap["service_state"], "**!\n", "```", icingaMap["notification_author"], icingaMap["notification_comment"], "\n\n", icingaMap["service_output"], "```")
        return
}


func host_or_service (icingaMap map[string]string) (message string) {
        if icingaMap["service_name"] != "" {
                emote := get_emote(icingaMap["service_state"])
                message = fmt.Sprintf("%v %v: %v%v%v%v%v%v%v%v%v%v", emote, icingaMap["notification_type"], "Service **", icingaMap["service_name"], "** on host **", icingaMap["host_name"], "** changed to **", icingaMap["service_state"], "**!\n", "```", icingaMap["service_output"], "```")
        }
        return
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


func send_msg (username string, url string, message string) () {
        payload := make(map[string]string)
        payload["username"] = username
        payload["content"] = message
        json, _ := json.Marshal(payload)
        http.Post(url, "application/json", bytes.NewBuffer(json))
}

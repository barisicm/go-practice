package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cast"
)

type Res struct {
	DeviceName     string
	SubscriberName string
	DomainName     string
	ReqTimestamp   string
}

type SubscriberUser struct {
	SubscriberId int
	Name         string
	AgentGroupId int
	Surname      string
	Email        string
	Icon         string
}

func main() {

	// var user SubscriberUser
	var results []Res

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	baseQuery := db.Table("dns_requests").
		Debug().
		Select("subscriber_users.name as subscriber_name, subscriber_devices.name as device_name, requested_domain_name as domain_name, TO_CHAR(req_timestamp, 'YYYY-MM-DD HH24:MI:SS') as req_timestamp").
		Joins("JOIN mac_ip_pairs ON dns_requests.requester_addr = mac_ip_pairs.device_ip").
		Joins("JOIN subscriber_devices ON subscriber_devices.mac = mac_ip_pairs.device_mac").
		Joins("JOIN subscriber_users ON subscriber_users.id = subscriber_devices.subscriber_user_id")

	subscriberSelected(baseQuery, results)
	// agentSelected(baseQuery, results)
	// devicesSelected(baseQuery, results)
	// subscriberUserSelected(baseQuery, results)
}

// Dns list db queries

// Get requests made based on the subscriber id
func subscriberSelected(baseQuery *gorm.DB, resultsObj []Res) {
	baseQuery.
		Where("subscriber_users.subscriber_id = ?", 2).
		Where("requested_domain_name LIKE ?", "%"+"hotmail"+"%").
		Scan(&resultsObj)

	for _, item := range resultsObj {
		fmt.Println(item)
	}
}

// Get requests made based on agent selected
func agentSelected(baseQuery *gorm.DB, resultsObj []Res) {
	baseQuery.
		Where("mac_ip_pairs.agent_id = ?", 142151166981867).
		Scan(&resultsObj)

	for _, item := range resultsObj {
		fmt.Println(item)
	}
}

// Get requests made based on devices selected
func devicesSelected(baseQuery *gorm.DB, resultsObj []Res) {
	// add select by multiple ID-s
	baseQuery.
		Where("subscriber_devices.id = ?", 6).
		Scan(&resultsObj)

	for _, item := range resultsObj {
		fmt.Println(item)
	}
}

// Get requests made based on the subscriber user id
// refactor this in the main project
func subscriberUserSelected(baseQuery *gorm.DB, resultsObj []Res) {
	// add multiple select
	baseQuery.
		Where("subscriber_devices.subscriber_user_id = ANY (?::INT[])", 2).
		Scan(&resultsObj)

	for _, item := range resultsObj {
		fmt.Println(item)
	}

}

func formatIDIntArrayForDb(idIntArray []int) string {
	formattedArray := "{"
	for _, item := range idIntArray {
		formattedArray += cast.ToString(item) + ","
	}
	formattedArray = strings.TrimSuffix(formattedArray, ",")
	formattedArray += "}"

	return formattedArray
}

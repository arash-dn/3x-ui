package sub

import "x-ui/xray"

func CalculateClientTraffic(clientTraffics []xray.ClientTraffic) xray.ClientTraffic {
	var traffic xray.ClientTraffic
	for _, clientTraffic := range clientTraffics {
		traffic.Up += clientTraffic.Up
		traffic.Down += clientTraffic.Down

		// Total traffic
		// The `Total` should not be summed. It is assumed to be the same for all clients with the same subId.
		// We take the first non-zero value. If one client has a limit, the subscription has that limit.
		if traffic.Total == 0 {
			traffic.Total = clientTraffic.Total
		}

		// Expiry time
		// We take the earliest expiry time among all clients.
		if clientTraffic.ExpiryTime > 0 {
			if traffic.ExpiryTime == 0 || clientTraffic.ExpiryTime < traffic.ExpiryTime {
				traffic.ExpiryTime = clientTraffic.ExpiryTime
			}
		}
	}
	return traffic
} 